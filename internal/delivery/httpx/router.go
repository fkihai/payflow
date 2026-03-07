package httpx

import (
	"net/http"

	"github.com/fkihai/payflow/internal/delivery/httpx/handler"
	"github.com/go-chi/chi/v5"
)

func Router(ph *handler.PaymentHandler, wh *handler.WebhookHandler) http.Handler {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Post("/create-transaction", ph.CreateTransaction())
			r.Post("/confirm-transaction", wh.ConfirmCharge())
		})
	})

	return r
}
