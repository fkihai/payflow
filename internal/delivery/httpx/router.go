package httpx

import (
	"net/http"

	"github.com/fkihai/payflow/internal/delivery/httpx/handler"
	"github.com/go-chi/chi/v5"
)

func Router(ph *handler.PaymentHandler) http.Handler {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Post("/create-transaction", ph.CreateTransaction())
		})
	})

	return r
}
