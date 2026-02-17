package payment

type OrderIDGenerator interface {
	Generate() (OrderID, error)
}
