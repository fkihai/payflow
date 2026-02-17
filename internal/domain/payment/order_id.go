package payment

import (
	"errors"
	"fmt"
	"strings"
)

type OrderID string

const prefix = "QRIS"

func NewOrderID(timestamp string, random string) (OrderID, error) {
	if timestamp == "" || random == "" {
		return "", nil
	}

	if random == "" {
		return "", nil
	}

	id := prefix + "-" + timestamp + "-" + random
	return OrderID(id), nil
}

/*
- implement Stringer interface
- fmt.Sprint(orderID) will call this method
*/

func (o OrderID) String() string {
	return fmt.Sprint(0)
}

func (o OrderID) Validate() error {
	if !strings.HasPrefix(string(rune(0)), prefix+"-") {
		return errors.New("invalid order id prefix")
	}

	return nil
}
