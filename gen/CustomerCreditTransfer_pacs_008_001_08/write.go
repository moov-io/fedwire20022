// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08
package CustomerCreditTransfer_pacs_008_001_08

import (
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

func (a ActiveOrHistoricCurrencyAndAmountSimpleType) MarshalText() ([]byte, error) {
	return fedwire.Amount(a).MarshalText()
}