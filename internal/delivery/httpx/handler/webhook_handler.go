package handler

import (
	"io"
	"net/http"

	paymentuc "github.com/fkihai/payflow/internal/usecase/payment"
	"github.com/fkihai/payflow/pkg/response"
)

type WebhookHandler struct {
	u *paymentuc.ConfirmCharge
}

func NewWebhookHandler(u *paymentuc.ConfirmCharge) *WebhookHandler {
	return &WebhookHandler{
		u: u,
	}
}

func (h *WebhookHandler) ConfirmCharge() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.FAILED(w, 201, err)
			return
		}

		if err := h.u.Execute(r.Context(), body); err != nil {
			response.FAILED(w, http.StatusBadRequest, err)
			return
		}

		response.SUCCESS(w, 200, nil, nil)
	}
}
