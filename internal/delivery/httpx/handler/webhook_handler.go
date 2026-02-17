package handler

import "net/http"

type WebhookHandler interface {
	Webhook() http.HandlerFunc
}
