// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.060.001.04
package camt_060_001_04

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	AcctRptgReq AccountReportingRequestV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 AcctRptgReq"`
}

// XSD ComplexType declarations

type AccountIdentification4Choice struct {
	IBAN *IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Othr,omitempty"`
}

type AccountReportingRequestV04 struct {
	GrpHdr      GroupHeader76         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 GrpHdr"`
	RptgReq     []ReportingRequest4   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 RptgReq"`
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 SplmtryData,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    *ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Cd,omitempty"`
	Prtry *Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Prtry,omitempty"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value ActiveOrHistoricCurrencyAndAmountSimpleType `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode                `xml:"Ccy,attr"`
}

type BalanceSubType1Choice struct {
	Cd    *ExternalBalanceSubType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Cd,omitempty"`
	Prtry *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Prtry,omitempty"`
}

type BalanceType10Choice struct {
	Cd    *ExternalBalanceType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Cd,omitempty"`
	Prtry *Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Prtry,omitempty"`
}

type BalanceType13 struct {
	CdOrPrtry BalanceType10Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 CdOrPrtry"`
	SubTp     *BalanceSubType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 SubTp,omitempty"`
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 FinInstnId"`
	BrnchId    *BranchData2                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      *Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Id,omitempty"`
	Nm      *Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Nm,omitempty"`
	PstlAdr *PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 PstlAdr,omitempty"`
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Id"`
	Tp  *CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Tp,omitempty"`
	Ccy *ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Ccy,omitempty"`
	Nm  *Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    *ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Cd,omitempty"`
	Prtry *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Prtry,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    *ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Cd,omitempty"`
	Prtry *Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 ClrSysId,omitempty"`
	MmbId    Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   *NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 NmPrfx,omitempty"`
	Nm       *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Nm,omitempty"`
	PhneNb   *PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 PhneNb,omitempty"`
	MobNb    *PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 MobNb,omitempty"`
	FaxNb    *PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 FaxNb,omitempty"`
	EmailAdr *Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 EmailAdr,omitempty"`
	Othr     *Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Othr,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     fedwire.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 BirthDt"`
	PrvcOfBirth *Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 CityOfBirth"`
	CtryOfBirth CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 CtryOfBirth"`
}

type DatePeriodDetails1 struct {
	FrDt fedwire.ISODate  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 FrDt"`
	ToDt *fedwire.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 ToDt,omitempty"`
}

type EntryStatus1Choice struct {
	Cd    *ExternalEntryStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Cd,omitempty"`
	Prtry *Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Prtry,omitempty"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    *ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Cd,omitempty"`
	Prtry *Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Prtry,omitempty"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       *BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 BICFI,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 ClrSysMmbId,omitempty"`
	Nm          *Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Nm,omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 SchmeNm,omitempty"`
	Issr    *Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 SchmeNm,omitempty"`
	Issr    *Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 SchmeNm,omitempty"`
	Issr    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 SchmeNm,omitempty"`
	Issr    *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Issr,omitempty"`
}

type GroupHeader76 struct {
	MsgId   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 MsgId"`
	CreDtTm fedwire.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 CreDtTm"`
	MsgSndr *Party35Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 MsgSndr,omitempty"`
}

type Limit2 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Amt"`
	CdtDbtInd FloorLimitType1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 CdtDbtInd"`
}

type OrganisationIdentification8 struct {
	AnyBIC *AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 AnyBIC,omitempty"`
	Othr   []*GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Cd,omitempty"`
	Prtry *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Prtry,omitempty"`
}

type Party34Choice struct {
	OrgId  *OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 OrgId,omitempty"`
	PrvtId *PersonIdentification13      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 PrvtId,omitempty"`
}

type Party35Choice struct {
	Pty *PartyIdentification125                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Agt,omitempty"`
}

type PartyIdentification125 struct {
	Nm        *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Nm,omitempty"`
	PstlAdr   *PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 PstlAdr,omitempty"`
	Id        *Party34Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Id,omitempty"`
	CtryOfRes *CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 CtryOfRes,omitempty"`
	CtctDtls  *ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 CtctDtls,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 DtAndPlcOfBirth,omitempty"`
	Othr            []*GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Cd,omitempty"`
	Prtry *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Prtry,omitempty"`
}

type PostalAddress6 struct {
	AdrTp       *AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 AdrTp,omitempty"`
	Dept        *Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Dept,omitempty"`
	SubDept     *Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 SubDept,omitempty"`
	StrtNm      *Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 StrtNm,omitempty"`
	BldgNb      *Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 BldgNb,omitempty"`
	PstCd       *Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 PstCd,omitempty"`
	TwnNm       *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 TwnNm,omitempty"`
	CtrySubDvsn *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 CtrySubDvsn,omitempty"`
	Ctry        *CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Ctry,omitempty"`
	AdrLine     []*Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 AdrLine,omitempty"`
}

type ReportingPeriod2 struct {
	FrToDt DatePeriodDetails1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 FrToDt"`
	FrToTm *TimePeriodDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 FrToTm,omitempty"`
	Tp     QueryType3Code      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Tp"`
}

type ReportingRequest4 struct {
	Id          *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Id,omitempty"`
	ReqdMsgNmId Max35Text                                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 ReqdMsgNmId"`
	Acct        *CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Acct,omitempty"`
	AcctOwnr    Party35Choice                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 AcctOwnr"`
	AcctSvcr    *BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 AcctSvcr,omitempty"`
	RptgPrd     *ReportingPeriod2                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 RptgPrd,omitempty"`
	RptgSeq     *SequenceRange1Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 RptgSeq,omitempty"`
	ReqdTxTp    *TransactionType2                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 ReqdTxTp,omitempty"`
	ReqdBalTp   []*BalanceType13                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 ReqdBalTp,omitempty"`
}

type SequenceRange1 struct {
	FrSeq Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 FrSeq"`
	ToSeq Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 ToSeq"`
}

type SequenceRange1Choice struct {
	FrSeq   *Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 FrSeq,omitempty"`
	ToSeq   *Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 ToSeq,omitempty"`
	FrToSeq []SequenceRange1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 FrToSeq"`
	EQSeq   []Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 EQSeq"`
	NEQSeq  []Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 NEQSeq"`
}

type SupplementaryData1 struct {
	PlcAndNm *Max350Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TimePeriodDetails1 struct {
	FrTm ISOTime  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 FrTm"`
	ToTm *ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 ToTm,omitempty"`
}

type TransactionType2 struct {
	Sts       EntryStatus1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 Sts"`
	CdtDbtInd CreditDebitCode    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 CdtDbtInd"`
	FlrLmt    []*Limit2          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04 FlrLmt,omitempty"`
}

// XSD SimpleType declarations

type ActiveOrHistoricCurrencyAndAmountSimpleType fedwire.Amount

type ActiveOrHistoricCurrencyCode string

type AddressType2Code string

const AddressType2CodeAddr AddressType2Code = "ADDR"
const AddressType2CodePbox AddressType2Code = "PBOX"
const AddressType2CodeHome AddressType2Code = "HOME"
const AddressType2CodeBizz AddressType2Code = "BIZZ"
const AddressType2CodeMlto AddressType2Code = "MLTO"
const AddressType2CodeDlvy AddressType2Code = "DLVY"

type AnyBICIdentifier string

type BICFIIdentifier string

type CountryCode string

type CreditDebitCode string

const CreditDebitCodeCrdt CreditDebitCode = "CRDT"
const CreditDebitCodeDbit CreditDebitCode = "DBIT"

type ExternalAccountIdentification1Code string

type ExternalBalanceSubType1Code string

type ExternalBalanceType1Code string

type ExternalCashAccountType1Code string

type ExternalClearingSystemIdentification1Code string

type ExternalEntryStatus1Code string

type ExternalFinancialInstitutionIdentification1Code string

type ExternalOrganisationIdentification1Code string

type ExternalPersonIdentification1Code string

type FloorLimitType1Code string

const FloorLimitType1CodeCred FloorLimitType1Code = "CRED"
const FloorLimitType1CodeDebt FloorLimitType1Code = "DEBT"
const FloorLimitType1CodeBoth FloorLimitType1Code = "BOTH"

type IBAN2007Identifier string

type ISOTime time.Time

type Max140Text string

type Max16Text string

type Max2048Text string

type Max34Text string

type Max350Text string

type Max35Text string

type Max70Text string

type NamePrefix1Code string

const NamePrefix1CodeDoct NamePrefix1Code = "DOCT"
const NamePrefix1CodeMist NamePrefix1Code = "MIST"
const NamePrefix1CodeMiss NamePrefix1Code = "MISS"
const NamePrefix1CodeMadm NamePrefix1Code = "MADM"

type PhoneNumber string

type QueryType3Code string

const QueryType3CodeAlll QueryType3Code = "ALLL"
const QueryType3CodeChng QueryType3Code = "CHNG"
const QueryType3CodeModf QueryType3Code = "MODF"
