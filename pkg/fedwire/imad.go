package fedwire

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// InputMessageAccountabilityData is an Identifier composed of the FedWire Input Cycle Date, Input Source, and Input Sequence Number.
//
// - The Input Cycle Date must be the Fedwire funds-transfer business day (8 characters, CCYYMMDD).
// - The Input Source must be the unique Endpoint ID of the Fedwire Sender (8 characters, alphanumeric).
// - The Input Sequence Number should be incremental and should start with 000001 at the start of each Input Cycle Date per Endpoint ID (6 characters, numeric).
func InputMessageAccountabilityData(cycleDate string, endpointID string, sequenceNumber int) string {
	return fmt.Sprintf("%8.8s%8.8s%6.6d", cycleDate, endpointID, sequenceNumber)
}

var (
	imadRegex = regexp.MustCompile(`^[0-9]{8}[A-Z0-9]{8}[0-9]{6}$`)
)

func ValidIMAD(input string) error {
	n := utf8.RuneCountInString(input)
	if n != 22 {
		return fmt.Errorf("IMAD must be 22 characters - found %d", n)
	}

	if !imadRegex.MatchString(input) {
		return errors.New("IMAD does not match regex")
	}

	return nil
}
