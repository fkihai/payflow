package midtrans

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/fkihai/payflow/internal/domain/payment"
	conf "github.com/fkihai/payflow/internal/infrastructure/config"
	p "github.com/fkihai/payflow/internal/infrastructure/payment"
)

type ClientMidtrans struct {
	Cfg    conf.PaymentGatewayConfig
	Client *http.Client
}

func NewClientMidtrans(cfg conf.PaymentGatewayConfig, client *http.Client) p.Gateway {
	return &ClientMidtrans{
		Cfg:    cfg,
		Client: client,
	}
}

func (c *ClientMidtrans) CreateTransaction(r p.GatewayRequest) (*p.GatewayResult, error) {

	var url string

	if c.Cfg.Env == conf.PaymentEnvSandbox {
		url = c.Cfg.SandBoxUrl
	} else {
		url = c.Cfg.ProductionUrl
	}

	payload := buildPayload(r.OrderID, r.Amount)
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

	var midRes midtransResponse
	if err := json.Unmarshal(bodyRes, &midRes); err != nil {
		return nil, err
	}

	result, err := toPaymentResponse(&midRes)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func toPaymentResponse(m *midtransResponse) (*p.GatewayResult, error) {

	amountInt, err := parseAmountToInt64(m.GrossAmount)
	if err != nil {
		return nil, err
	}

	var QrUrl string
	if len(m.Actions) > 0 {
		QrUrl = m.Actions[0].Url
	}

	result := &p.GatewayResult{
		OrderID:   m.OrderID,
		Amount:    amountInt,
		Status:    m.TransactionStatus,
		QrCodeUrl: QrUrl,
	}

	return result, nil
}

func parseAmountToInt64(amountStr string) (int64, error) {
	amountStr = strings.TrimSpace(amountStr)
	parts := strings.Split(amountStr, ".")
	return strconv.ParseInt(parts[0], 10, 64)
}

func buildPayload(orderID payment.OrderID, amount int64) map[string]any {
	return map[string]any{
		"payment_type": "qris",
		"transaction_details": map[string]any{
			"order_id":     orderID,
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
