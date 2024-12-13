// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.052.001.08
package Master_camt_052_001_08

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
	Id          BalanceReportFRS1   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Id"`
	CreDtTm     fedwire.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 CreDtTm"`
	Acct        CashAccount391      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Acct"`
	RltdAcct    CashAccount381      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 RltdAcct"`
	Bal         []CashBalance81     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Bal"`
	TxsSummry   TotalTransactions61 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 TxsSummry"`
	AddtlRptInf *Max500Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 AddtlRptInf,omitempty"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value ActiveOrHistoricCurrencyAndAmountSimpleType `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode                `xml:"Ccy,attr"`
}

type AmountAndDirection35 struct {
	Amt       NonNegativeDecimalNumber `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Amt"`
	CdtDbtInd CreditDebitCode          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 CdtDbtInd"`
}

type BalanceType10Choice1 struct {
	Prtry *BalanceTypeFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Prtry,omitempty"`
}

type BalanceType131 struct {
	CdOrPrtry BalanceType10Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 CdOrPrtry"`
}

type BankToCustomerAccountReportV08 struct {
	GrpHdr GroupHeader811   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 GrpHdr"`
	Rpt    AccountReport251 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Rpt"`
}

type BankTransactionCodeStructure41 struct {
	Prtry ProprietaryBankTransactionCodeStructure11 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Prtry"`
}

type CashAccount381 struct {
	Id AccountIdentification4Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Id"`
}

type CashAccount391 struct {
	Id AccountIdentification4Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Id"`
	Tp CashAccountType2Choice1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Tp"`
}

type CashAccountType2Choice1 struct {
	Prtry *AccountTypeFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Prtry,omitempty"`
}

type CashBalance81 struct {
	Tp        BalanceType131                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Tp"`
	CdtLine   []*CreditLine31                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 CdtLine,omitempty"`
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 CdtDbtInd"`
	Dt        DateAndDateTime2Choice1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Dt"`
}

type CreditLine31 struct {
	Incl TrueFalseIndicator                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Incl"`
	Tp   CreditLineType1Choice1            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Tp"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Amt"`
	Dt   DateAndDateTime2Choice1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Dt"`
}

type CreditLineType1Choice1 struct {
	Prtry *CreditLineTypeFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Prtry,omitempty"`
}

type DateAndDateTime2Choice1 struct {
	DtTm *fedwire.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 DtTm,omitempty"`
}

type GenericAccountIdentification11 struct {
	Id RoutingNumberFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Id"`
}

type GroupHeader811 struct {
	MsgId       AccountReportingFedwireFunds1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 MsgId"`
	CreDtTm     fedwire.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 CreDtTm"`
	MsgPgntn    Pagination1                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 MsgPgntn"`
	OrgnlBizQry *OriginalBusinessQuery11      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 OrgnlBizQry,omitempty"`
}

type NumberAndSumOfTransactions11 struct {
	NbOfNtries *Max15NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 NbOfNtries,omitempty"`
	Sum        DecimalNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Sum"`
}

type OriginalBusinessQuery11 struct {
	MsgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 MsgId"`
	MsgNmId MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 MsgNmId"`
	CreDtTm fedwire.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 CreDtTm"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 PgNb"`
	LastPgInd YesNoIndicator  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 LastPgInd"`
}

type ProprietaryBankTransactionCodeStructure11 struct {
	Cd TransactionsSummaryTypeFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Cd"`
}

type TotalTransactions61 struct {
	TtlNtriesPerBkTxCd []TotalsPerBankTransactionCode51 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 TtlNtriesPerBkTxCd"`
}

type TotalsPerBankTransactionCode51 struct {
	TtlNetNtry AmountAndDirection35           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 TtlNetNtry"`
	CdtNtries  NumberAndSumOfTransactions11   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 CdtNtries"`
	DbtNtries  NumberAndSumOfTransactions11   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 DbtNtries"`
	BkTxCd     BankTransactionCodeStructure41 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 BkTxCd"`
	Dt         DateAndDateTime2Choice1        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Dt"`
}

// XSD SimpleType declarations

type AccountReportingFedwireFunds1 string

const AccountReportingFedwireFunds1Abar AccountReportingFedwireFunds1 = "ABAR"

type AccountTypeFRS1 string

const AccountTypeFRS1M AccountTypeFRS1 = "M"

type ActiveOrHistoricCurrencyAndAmountSimpleType fedwire.Amount

type ActiveOrHistoricCurrencyCode string

type BalanceReportFRS1 string

const BalanceReportFRS1Prov BalanceReportFRS1 = "PROV"
const BalanceReportFRS1Abms BalanceReportFRS1 = "ABMS"
const BalanceReportFRS1Itrm BalanceReportFRS1 = "ITRM"
const BalanceReportFRS1Finl BalanceReportFRS1 = "FINL"

type BalanceTypeFRS1 string

const BalanceTypeFRS1Abal BalanceTypeFRS1 = "ABAL"
const BalanceTypeFRS1Aval BalanceTypeFRS1 = "AVAL"
const BalanceTypeFRS1Avld BalanceTypeFRS1 = "AVLD"
const BalanceTypeFRS1Dlod BalanceTypeFRS1 = "DLOD"
const BalanceTypeFRS1Obfl BalanceTypeFRS1 = "OBFL"
const BalanceTypeFRS1Obnl BalanceTypeFRS1 = "OBNL"
const BalanceTypeFRS1Obpl BalanceTypeFRS1 = "OBPL"

type CreditDebitCode string

const CreditDebitCodeCrdt CreditDebitCode = "CRDT"
const CreditDebitCodeDbit CreditDebitCode = "DBIT"

type CreditLineTypeFRS1 string

const CreditLineTypeFRS1Coll CreditLineTypeFRS1 = "COLL"
const CreditLineTypeFRS1Clod CreditLineTypeFRS1 = "CLOD"
const CreditLineTypeFRS1Ulod CreditLineTypeFRS1 = "ULOD"
const CreditLineTypeFRS1Ncap CreditLineTypeFRS1 = "NCAP"
const CreditLineTypeFRS1Ccap CreditLineTypeFRS1 = "CCAP"

type DecimalNumber float64

type Max15NumericText string

type Max35Text string

type Max500Text string

type Max5NumericText string

type MessageNameIdentificationFRS1 string

type NonNegativeDecimalNumber float64

type RoutingNumberFRS1 string

type TransactionsSummaryTypeFRS1 string

const TransactionsSummaryTypeFRS1Fdws TransactionsSummaryTypeFRS1 = "FDWS"
const TransactionsSummaryTypeFRS1Nsse TransactionsSummaryTypeFRS1 = "NSSE"
const TransactionsSummaryTypeFRS1Avot TransactionsSummaryTypeFRS1 = "AVOT"
const TransactionsSummaryTypeFRS1Uvot TransactionsSummaryTypeFRS1 = "UVOT"
const TransactionsSummaryTypeFRS1Memo TransactionsSummaryTypeFRS1 = "MEMO"
const TransactionsSummaryTypeFRS1Fdwf TransactionsSummaryTypeFRS1 = "FDWF"
const TransactionsSummaryTypeFRS1Fdap TransactionsSummaryTypeFRS1 = "FDAP"
const TransactionsSummaryTypeFRS1Fdnf TransactionsSummaryTypeFRS1 = "FDNF"

type TrueFalseIndicator bool

type YesNoIndicator bool
