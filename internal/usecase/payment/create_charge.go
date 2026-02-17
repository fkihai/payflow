package paymentuc

import (
	"context"

	"github.com/fkihai/payflow/internal/domain"
)

type CreateCharge struct {
	repo   ChargeRepo
	gw     Gateway
	oidGen OIDGenerator
}

func NewCreateCharge(repo ChargeRepo, gw Gateway, oidGen OIDGenerator) *CreateCharge {
	return &CreateCharge{
		repo:   repo,
		gw:     gw,
		oidGen: oidGen,
	}
}
func (uc *CreateCharge) Execute(ctx context.Context, req *CreateChargeRequest) (*CreateChargeResult, error) {

	// Generate OID
	oid, err := uc.oidGen.Generate()
	if err != nil {
		return nil, err
	}

	// request payment gateway
	in := GatewayRequest{
		OID:    oid,
		Amount: req.Amount,
	}

	result, err := uc.gw.CreateCharge(&in)
	if err != nil {
		return nil, err
	}

	// Store database
	store := domain.Charge{
		OID:       result.OID,
		Amount:    result.Amount,
		Status:    result.Status,
		ExpiresAt: 0,
		QrUrl:     result.QrUrl,
	}

	chg, err := uc.repo.Create(ctx, &store)
	if err != nil {
		return nil, err
	}

	out := CreateChargeResult{
		OID:    chg.OID,
		Amount: chg.Amount,
		Status: chg.Status,
		QrUrl:  chg.QrUrl,
	}

	return &out, nil

}
