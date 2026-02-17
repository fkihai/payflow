package domain

import (
	"errors"
	"fmt"
	"strings"
)

type OID string

const prefix = "QRIS"

func NewOID(timestamp string, random string) (OID, error) {
	if timestamp == "" || random == "" {
		return "", errors.New("timestamp/random empty")
	}

	id := fmt.Sprintf("%s-%s-%s", prefix, timestamp, random)
	o := OID(id)

	if err := o.Validate(); err != nil {
		return "", err
	}
	return o, nil
}

func (o OID) String() string {
	return string(o)
}

func (o OID) Validate() error {
	s := string(o)

	if !strings.HasPrefix(s, prefix+"-") {
		return errors.New("invalid order id prefix")
	}

	parts := strings.Split(s, "-")
	if len(parts) != 3 {
		return errors.New("invalid order id format")
	}

	return nil
}
