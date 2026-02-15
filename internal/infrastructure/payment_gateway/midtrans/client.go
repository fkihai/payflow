package midtrans

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/fkihai/payflow/internal/delivery/httpx/dto"
)

type Client struct {
	ServerKey string
	SandBox   bool
	Debug     bool
}

func (c *Client) CreateTransactions(req dto.CreatePaymentRequest) (*SnapResponse, error) {

	url := "https://api.midtrans.com/v2/charge"

	if c.SandBox {
		url = "https://api.sandbox.midtrans.com/v2/charge"
	}

	orderID, err := generateOrderID()
	if err != nil {
		return nil, err
	}

	payload := map[string]any{
		"payment_type": "qris",
		"transaction_details": map[string]any{
			"order_id":     orderID,
			"gross_amount": req.Amount,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(c.ServerKey + ":"))
	httpReq.Header.Set("Authorization", auth)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	// debug request

	if c.Debug {
		fmt.Println("========= REQUEST =========")
		dumpReq, _ := httputil.DumpRequestOut(httpReq, true)
		fmt.Println(string(dumpReq))
		fmt.Println("===========================")
	}

	client := http.Client{}

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// debug response
	if c.Debug {
		fmt.Println("========= RESPONSE =========")
		fmt.Println("Status Code:", resp.StatusCode)
		fmt.Println("Headers:", resp.Header)
		fmt.Print("Body: ")
		fmt.Println(string(bodyBytes))
		fmt.Println("============================")

	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("midtrans error (%d): %s", resp.StatusCode, string(bodyBytes))
	}

	var res SnapResponse
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func generateOrderID() (string, error) {
	// Timestamp format: YYYYMMDDHHMMSS
	timestamp := time.Now().Format("20060102150405")

	randomStr, err := randomBase64(8)
	if err != nil {
		return "", nil
	}

	orderID := fmt.Sprintf("QRIS-%s-%s", timestamp, randomStr)
	return orderID, nil
}

func randomBase64(length int) (string, error) {
	result := make([]byte, length)
	for i := range length {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(base62))))
		if err != nil {
			return "", err
		}
		result[i] = base62[num.Int64()]
	}
	return string(result), nil
}
