// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03
package PaymentStatusRequest_pacs_028_001_03

import (
	"encoding/xml"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	FIToFIPmtStsReq FIToFIPaymentStatusRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 FIToFIPmtStsReq"`
}

// XSD ComplexType declarations

type BranchAndFinancialInstitutionIdentification61 struct {
	FinInstnId FinancialInstitutionIdentification181 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 FinInstnId"`
}

type ClearingSystemIdentification2Choice1 struct {
	Cd *ExternalClearingSystemIdentification1CodeFixed `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 Cd,omitempty"`
}

type ClearingSystemMemberIdentification21 struct {
	ClrSysId ClearingSystemIdentification2Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 ClrSysId"`
	MmbId    RoutingNumberFRS1                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 MmbId"`
}

type FIToFIPaymentStatusRequestV03 struct {
	GrpHdr GroupHeader911         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 GrpHdr"`
	TxInf  PaymentTransaction1131 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 TxInf"`
}

type FinancialInstitutionIdentification181 struct {
	ClrSysMmbId ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 ClrSysMmbId"`
}

type GroupHeader911 struct {
	MsgId   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 MsgId"`
	CreDtTm fedwire.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 CreDtTm"`
}

type OriginalGroupInformation291 struct {
	OrgnlMsgId   IMADFedwireFunds1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlMsgId"`
	OrgnlMsgNmId MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlMsgNmId"`
	OrgnlCreDtTm fedwire.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlCreDtTm"`
}

type PaymentTransaction1131 struct {
	OrgnlGrpInf     OriginalGroupInformation291                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlGrpInf"`
	OrgnlInstrId    *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlEndToEndId,omitempty"`
	OrgnlTxId       *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlTxId,omitempty"`
	OrgnlUETR       UUIDv4Identifier                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlUETR"`
	InstgAgt        BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 InstgAgt"`
	InstdAgt        BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 InstdAgt"`
}

// XSD SimpleType declarations

type ExternalClearingSystemIdentification1CodeFixed string

const ExternalClearingSystemIdentification1CodeFixedUsaba ExternalClearingSystemIdentification1CodeFixed = "USABA"

type IMADFedwireFunds1 string

type Max35Text string

type MessageNameIdentificationFRS1 string

type RoutingNumberFRS1 string

type UUIDv4Identifier string
