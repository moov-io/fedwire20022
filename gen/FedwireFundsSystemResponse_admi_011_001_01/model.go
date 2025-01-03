// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:admi.011.001.01
package FedwireFundsSystemResponse_admi_011_001_01

import (
	"encoding/xml"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	SysEvtAck SystemEventAcknowledgementV01 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 SysEvtAck"`
}

// XSD ComplexType declarations

type Event11 struct {
	EvtCd    EventFedwireFunds1              `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 EvtCd"`
	EvtParam EndpointIdentifierFedwireFunds1 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 EvtParam"`
	EvtTm    fedwire.ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 EvtTm"`
}

type SystemEventAcknowledgementV01 struct {
	MsgId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 MsgId"`
	AckDtls Event11   `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 AckDtls"`
}

// XSD SimpleType declarations

type EndpointIdentifierFedwireFunds1 string

type EventFedwireFunds1 string

const EventFedwireFunds1Ping EventFedwireFunds1 = "PING"

type Max35Text string
