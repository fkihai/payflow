package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fkihai/payflow/internal/delivery/httpx/dto"
	"github.com/fkihai/payflow/internal/delivery/httpx/response"
	"github.com/fkihai/payflow/internal/infrastructure/payment"
	pu "github.com/fkihai/payflow/internal/usecase/payment"
)

type PaymentHandler struct {
	u *pu.PaymentUsecase
}

func NewPaymentHandler(u *pu.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{
		u: u,
	}
}
func (h *PaymentHandler) CreateTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req dto.CreatePaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.FAILED(w, 201, err)
			return
		}

		in := payment.CreateTransactionInput{
			Amount: req.Amount,
		}

		trx, err := h.u.CreatePayment(r.Context(), in)
		if err != nil {
			response.FAILED(w, http.StatusBadRequest, err)
			return
		}

		dataRes := dto.CreatePaymentResponse{
			ID:      trx.ID,
			OrderID: trx.OrderId,
			Amount:  trx.Amount,
			Status:  trx.Status,
			QrUrl:   trx.QrUrl,
		}
		response.SUCCESS(w, 201, dataRes, nil)
	}
}
