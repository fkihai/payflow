package payment

type Gateway interface {
	CreateTransaction(req GatewayRequest) (*GatewayResult, error)
}
