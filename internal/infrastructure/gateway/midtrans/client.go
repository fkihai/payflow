package midtrans

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/fkihai/payflow/internal/domain"
	conf "github.com/fkihai/payflow/internal/infrastructure/config"
	pu "github.com/fkihai/payflow/internal/usecase/payment"
)

type ClientMidtrans struct {
	Cfg    conf.PaymentGatewayConfig
	Client *http.Client
}

func NewClientMidtrans(cfg conf.PaymentGatewayConfig, client *http.Client) pu.Gateway {
	return &ClientMidtrans{
		Cfg:    cfg,
		Client: client,
	}
}

func (c *ClientMidtrans) CreateCharge(req *pu.GatewayRequest) (*pu.GatewayResult, error) {

	var url string

	if c.Cfg.Env == conf.PaymentEnvSandbox {
		url = c.Cfg.SandBoxUrl
	} else {
		url = c.Cfg.ProductionUrl
	}

	payload := buildPayload(req.OID, req.Amount)
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	buildHeader(httpReq, c.Cfg.ServerKey)

	resp, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("midtrans request error: %w", err)
	}
	defer resp.Body.Close()

	bodyRes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("midtrans error (%d): %s", resp.StatusCode, string(bodyRes))
	}

	var res Response
	if err := json.Unmarshal(bodyRes, &res); err != nil {
		return nil, err
	}

	result, err := toGatewayResult(&res)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ClientMidtrans) ConfirmCharge(ctx context.Context, payload []byte) (*pu.WebhookEvent, error) {

	var req midtransWebhookEvent
	if err := json.Unmarshal(payload, &req); err != nil {
		return nil, err
	}

	data := string(req.OID) + req.StatusCode + req.GrossAmount + c.Cfg.ServerKey
	gen := generateSignature(data)
	if gen != req.Signature {
		return nil, fmt.Errorf("invalid signature")
	}

	amountInt, err := parseAmountToInt64(req.GrossAmount)
	if err != nil {
		return nil, err
	}

	return &pu.WebhookEvent{
		ChgID:       req.TrxID,
		ChgStatus:   req.TrxStatus,
		ChgPaid:     req.PaidAt,
		GrossAmount: amountInt,
	}, nil

}

func toGatewayResult(m *Response) (*pu.GatewayResult, error) {

	amountInt, err := parseAmountToInt64(m.GrossAmount)
	if err != nil {
		return nil, err
	}

	var QrUrl string
	if len(m.Actions) > 0 {
		QrUrl = m.Actions[0].Url
	}

	result := &pu.GatewayResult{
		OID:    m.OID,
		Amount: amountInt,
		Status: m.TrxStatus,
		QrUrl:  QrUrl,
	}

	return result, nil
}

func parseAmountToInt64(amountStr string) (int64, error) {
	amountStr = strings.TrimSpace(amountStr)
	parts := strings.Split(amountStr, ".")
	return strconv.ParseInt(parts[0], 10, 64)
}

func buildPayload(OID domain.OID, amount int64) map[string]any {
	return map[string]any{
		"payment_type": "qris",
		"transaction_details": map[string]any{
			"order_id":     OID,
			"gross_amount": amount,
		},
	}
}

func buildHeader(r *http.Request, serverKey string) {
	auth := base64.StdEncoding.EncodeToString([]byte(serverKey + ":"))
	r.Header.Set("Authorization", "Basic "+auth)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")
}
