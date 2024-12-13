// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:admi.004.001.02
package FedwireFundsBroadcast_admi_004_001_02

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
	EvtCd    EventFedwireFunds1  `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtCd"`
	EvtParam fedwire.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtParam"`
	EvtDesc  *Max1000Text        `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtDesc,omitempty"`
	EvtTm    fedwire.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtTm"`
}

type SystemEventNotificationV02 struct {
	EvtInf Event21 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtInf"`
}

// XSD SimpleType declarations

type EventFedwireFunds1 string

const EventFedwireFunds1Open EventFedwireFunds1 = "OPEN"
const EventFedwireFunds1Clsd EventFedwireFunds1 = "CLSD"
const EventFedwireFunds1Extn EventFedwireFunds1 = "EXTN"
const EventFedwireFunds1Adhc EventFedwireFunds1 = "ADHC"

type Max1000Text string
