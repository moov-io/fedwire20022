// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:admi.004.001.02
package ConnectionCheck_admi_004_001_02

import (
	"encoding/xml"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	SysEvtNtfctn SystemEventNotificationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 SysEvtNtfctn"`
}

// XSD ComplexType declarations

type Event21 struct {
	EvtCd    EventFedwireFunds1              `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtCd"`
	EvtParam EndpointIdentifierFedwireFunds1 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtParam"`
	EvtTm    fedwire.ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtTm"`
}

type SystemEventNotificationV02 struct {
	EvtInf Event21 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtInf"`
}

// XSD SimpleType declarations

type EndpointIdentifierFedwireFunds1 string

type EventFedwireFunds1 string

const EventFedwireFunds1Ping EventFedwireFunds1 = "PING"
