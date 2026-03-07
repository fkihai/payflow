package handler

import (
	"encoding/json"
	"net/http"

	pu "github.com/fkihai/payflow/internal/usecase/payment"
	"github.com/fkihai/payflow/pkg/response"
)

type PaymentHandler struct {
	uc *pu.CreateCharge
}

func NewPaymentHandler(u *pu.CreateCharge) *PaymentHandler {
	return &PaymentHandler{
		uc: u,
	}
}
func (h *PaymentHandler) CreateTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req pu.CreateChargeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.FAILED(w, 201, err)
			return
		}

		in := pu.CreateChargeRequest{
			Amount: req.Amount,
		}

		trx, err := h.uc.Execute(r.Context(), &in)
		if err != nil {
			response.FAILED(w, http.StatusBadRequest, err)
			return
		}

		dataRes := ChargerResponse{
			ID:     trx.ID,
			OID:    trx.OID,
			Amount: trx.Amount,
			Status: trx.Status,
			QrUrl:  trx.QrUrl,
		}
		response.SUCCESS(w, 201, dataRes, nil)
	}
}
