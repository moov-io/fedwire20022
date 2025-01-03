// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.052.001.08
package EndpointGapReport_camt_052_001_08

import (
	"encoding/xml"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	BkToCstmrAcctRpt BankToCustomerAccountReportV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 BkToCstmrAcctRpt"`
}

// XSD ComplexType declarations

type AccountIdentification4Choice1 struct {
	Othr *GenericAccountIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Othr,omitempty"`
}

type AccountReport251 struct {
	Id          GapTypeFedwireFunds1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Id"`
	CreDtTm     fedwire.ISODateTime  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 CreDtTm"`
	Acct        CashAccount391       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Acct"`
	AddtlRptInf Max500Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 AddtlRptInf"`
}

type BankToCustomerAccountReportV08 struct {
	GrpHdr GroupHeader811     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 GrpHdr"`
	Rpt    []AccountReport251 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Rpt"`
}

type CashAccount391 struct {
	Id AccountIdentification4Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Id"`
}

type GenericAccountIdentification11 struct {
	Id EndpointIdentifierFedwireFunds1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Id"`
}

type GroupHeader811 struct {
	MsgId    AccountReportingFedwireFunds1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 MsgId"`
	CreDtTm  fedwire.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 CreDtTm"`
	MsgPgntn Pagination1                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 MsgPgntn"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 PgNb"`
	LastPgInd YesNoIndicator  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 LastPgInd"`
}

// XSD SimpleType declarations

type AccountReportingFedwireFunds1 string

const AccountReportingFedwireFunds1Gapr AccountReportingFedwireFunds1 = "GAPR"

type EndpointIdentifierFedwireFunds1 string

type GapTypeFedwireFunds1 string

const GapTypeFedwireFunds1Omad GapTypeFedwireFunds1 = "OMAD"
const GapTypeFedwireFunds1Imad GapTypeFedwireFunds1 = "IMAD"

type Max500Text string

type Max5NumericText string

type YesNoIndicator bool
