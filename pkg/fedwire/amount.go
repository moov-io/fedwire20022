package fedwire

import (
	"fmt"
	"math"
)

type Amount float64

func (a Amount) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", a)), nil
}

func (a Amount) Validate() error {
	if math.IsNaN(float64(a)) || math.IsInf(float64(a), 0) {
		return fmt.Errorf("amount must be a finite number")
	}
	_, err := a.MarshalText()
	return err
}
