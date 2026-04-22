package fedwiretest

import (
	"encoding/xml"

	"github.com/moov-io/fedwire20022/gen/fedwirefunds_incoming"
	"github.com/moov-io/fedwire20022/gen/fedwirefunds_outgoing"
)

// FlipMessageDirection will attempt to switch XML messages from incoming to outgoing and vice-versa.
// This is needed so we can receive our own originated messages.
//
// Right now only a subset of possible FedWire messages are supported
func FlipMessageDirection(input []byte) ([]byte, error) {
	// Try to unmarshal as outgoing
	var outgoingDoc fedwirefunds_outgoing.FedwireFundsOutgoing
	err := xml.Unmarshal(input, &outgoingDoc)
	if err == nil {
		// Check which message type it is
		switch {
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsCustomerCreditTransfer != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsCustomerCreditTransfer
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsCustomerCreditTransfer: &fedwirefunds_incoming.FedwireFundsCustomerCreditTransfer{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsMessageReject != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsMessageReject
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsMessageReject: &fedwirefunds_incoming.FedwireFundsMessageReject{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsPaymentReturn != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsPaymentReturn
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsPaymentReturn: &fedwirefunds_incoming.FedwireFundsPaymentReturn{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsMessageReject != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsMessageReject
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsMessageReject: &fedwirefunds_incoming.FedwireFundsMessageReject{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsPaymentReturn != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsPaymentReturn
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsPaymentReturn: &fedwirefunds_incoming.FedwireFundsPaymentReturn{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsFinancialInstitutionCreditTransfer != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsFinancialInstitutionCreditTransfer
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsFinancialInstitutionCreditTransfer: &fedwirefunds_incoming.FedwireFundsFinancialInstitutionCreditTransfer{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsPaymentStatusRequest != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsPaymentStatusRequest
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsPaymentStatusRequest: &fedwirefunds_incoming.FedwireFundsPaymentStatusRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsDrawdownRequest != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsDrawdownRequest
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsDrawdownRequest: &fedwirefunds_incoming.FedwireFundsDrawdownRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsDrawdownResponse != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsDrawdownResponse
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsDrawdownResponse: &fedwirefunds_incoming.FedwireFundsDrawdownResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsReturnRequestResponse != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsReturnRequestResponse
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsReturnRequestResponse: &fedwirefunds_incoming.FedwireFundsReturnRequestResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsReturnRequest != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsReturnRequest
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsReturnRequest: &fedwirefunds_incoming.FedwireFundsReturnRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsInvestigationRequest != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsInvestigationRequest
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsInvestigationRequest: &fedwirefunds_incoming.FedwireFundsInvestigationRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsInvestigationResponse != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsInvestigationResponse
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsInvestigationResponse: &fedwirefunds_incoming.FedwireFundsInvestigationResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsBroadcast != nil:
			msg := outgoingDoc.FedwireFundsOutgoingMessage.FedwireFundsBroadcast
			return xml.Marshal(fedwirefunds_incoming.FedwireFundsIncoming{
				FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
					FedwireFundsConnectionCheck: &fedwirefunds_incoming.FedwireFundsConnectionCheck{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		}
	}

	// Try to unmarshal as incoming
	var incomingDoc fedwirefunds_incoming.FedwireFundsIncoming
	err = xml.Unmarshal(input, &incomingDoc)
	if err == nil {
		// Check which message type it is
		switch {
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsCustomerCreditTransfer != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsCustomerCreditTransfer
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsCustomerCreditTransfer: &fedwirefunds_outgoing.FedwireFundsCustomerCreditTransfer{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsMessageReject != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsMessageReject
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsMessageReject: &fedwirefunds_outgoing.FedwireFundsMessageReject{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsPaymentReturn != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsPaymentReturn
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsPaymentReturn: &fedwirefunds_outgoing.FedwireFundsPaymentReturn{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsFinancialInstitutionCreditTransfer != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsFinancialInstitutionCreditTransfer
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsFinancialInstitutionCreditTransfer: &fedwirefunds_outgoing.FedwireFundsFinancialInstitutionCreditTransfer{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsPaymentStatusRequest != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsPaymentStatusRequest
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsPaymentStatusRequest: &fedwirefunds_outgoing.FedwireFundsPaymentStatusRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsDrawdownRequest != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsDrawdownRequest
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsDrawdownRequest: &fedwirefunds_outgoing.FedwireFundsDrawdownRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsDrawdownResponse != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsDrawdownResponse
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsDrawdownResponse: &fedwirefunds_outgoing.FedwireFundsDrawdownResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsReturnRequestResponse != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsReturnRequestResponse
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsReturnRequestResponse: &fedwirefunds_outgoing.FedwireFundsReturnRequestResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsReturnRequest != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsReturnRequest
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsReturnRequest: &fedwirefunds_outgoing.FedwireFundsReturnRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsInvestigationRequest != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsInvestigationRequest
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsInvestigationRequest: &fedwirefunds_outgoing.FedwireFundsInvestigationRequest{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsInvestigationResponse != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsInvestigationResponse
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsInvestigationResponse: &fedwirefunds_outgoing.FedwireFundsInvestigationResponse{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		case incomingDoc.FedwireFundsIncomingMessage.FedwireFundsConnectionCheck != nil:
			msg := incomingDoc.FedwireFundsIncomingMessage.FedwireFundsConnectionCheck
			return xml.Marshal(fedwirefunds_outgoing.FedwireFundsOutgoing{
				FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
					FedwireFundsBroadcast: &fedwirefunds_outgoing.FedwireFundsBroadcast{
						AppHdr:   msg.AppHdr,
						Document: msg.Document,
					},
				},
			})
		}
	}

	return input, nil
}
