package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fkihai/payflow/internal/delivery/httpx/dto"
	"github.com/fkihai/payflow/internal/delivery/httpx/response"
	"github.com/fkihai/payflow/internal/domain/entity"
	"github.com/fkihai/payflow/internal/infrastructure/config"
	"github.com/fkihai/payflow/internal/infrastructure/payment_gateway/midtrans"
	"github.com/fkihai/payflow/internal/usecase/payment"
)

type PaymentHandler struct {
	pu  *payment.PaymentUsecase
	cfg config.PaymentGatewayConfig
}

func NewPaymentHandler(pu *payment.PaymentUsecase, cfg config.PaymentGatewayConfig) *PaymentHandler {
	return &PaymentHandler{
		pu:  pu,
		cfg: cfg,
	}
}
func (h *PaymentHandler) CreateTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var trx dto.CreatePaymentRequest

		if err := json.NewDecoder(r.Body).Decode(&trx); err != nil {
			response.FAILED(w, 201, err)
			return
		}

		// 1. Request payment to midtrans (client.go)
		mt := midtrans.Client{
			ServerKey: h.cfg.ServerKey,
			SandBox:   true,
			Debug:     false,
		}
		res, err := mt.CreateTransactions(dto.CreatePaymentRequest{
			Amount: trx.Amount,
		})

		if err != nil {
			response.FAILED(w, http.StatusBadRequest, err)
			return
		}

		// 2. Create Payment on db
		amountInt, err := ParseAmount(res.GrossAmount)

		if err != nil {
			response.FAILED(w, http.StatusBadRequest, err)
			return
		}

		if len(res.Actions) == 0 {
			response.FAILED(w, http.StatusAccepted, fmt.Errorf("qr url not found"))
			return
		}

		t, err := h.pu.CreatePayment(r.Context(), entity.Transaction{
			OrderId: res.OrderID,
			Amount:  amountInt,
			Status:  res.TransactionStatus,
			QrUrl:   res.Actions[0].Url,
		})

		if err != nil {
			response.FAILED(w, http.StatusBadRequest, err)
			return
		}

		// 3.Response to client
		dataRes := dto.CreatePaymentResponse{
			ID:      t.ID,
			OrderID: t.OrderId,
			Amount:  t.Amount,
			Status:  t.Status,
			QrUrl:   t.QrUrl,
		}
		response.SUCCESS(w, 201, dataRes, nil)
	}
}

func ParseAmount(amountStr string) (int64, error) {
	amountStr = strings.TrimSpace(amountStr)
	f, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, err
	}
	return int64(f), nil
}
