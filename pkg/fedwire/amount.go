package fedwire

import (
	"fmt"
	"math"
)

// Amount is a currency amount. It is written with two decimal places,
// so a value such as 1.234 is written as "1.23".
type Amount float64

// String returns the amount as it is written in XML, with two decimal places.
func (a Amount) String() string {
	return fmt.Sprintf("%.2f", float64(a))
}

func (a Amount) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

func (a Amount) Validate() error {
	if math.IsNaN(float64(a)) || math.IsInf(float64(a), 0) {
		return fmt.Errorf("amount must be a finite number")
	}
	_, err := a.MarshalText()
	return err
}
