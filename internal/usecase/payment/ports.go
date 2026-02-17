package paymentuc

import (
	c "context"
	"time"

	d "github.com/fkihai/payflow/internal/domain"
)

type ChargeRepo interface {
	Create(ctx c.Context, chg *d.Charge) (*d.Charge, error)
	FindByOID(ctx c.Context, OID d.OID) (*d.Charge, error)
	Update(ctx c.Context, OID d.OID, status d.ChargeStatus, paidAt *time.Time) error
}

type OIDGenerator interface {
	Generate() (d.OID, error)
}

type Gateway interface {
	CreateCharge(req *GatewayRequest) (*GatewayResult, error)
	ConfirmCharge(ctx c.Context, payload []byte) (*WebhookEvent, error)
}
