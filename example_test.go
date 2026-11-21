package fedwire20022_test

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/moov-io/fedwire20022/gen/fedwirefunds_incoming"
	"github.com/moov-io/fedwire20022/gen/head_001_001_03"
	"github.com/moov-io/fedwire20022/gen/pacs_009_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// Example_readMessage reads a Fedwire Funds customer credit transfer (pacs.008)
// from an XML file and validates it.
func Example_readMessage() {
	bs, err := os.ReadFile(filepath.Join("testdata", "PACS008_CustomerCreditTransfer.txt"))
	if err != nil {
		panic(err)
	}

	var msg fedwirefunds_incoming.FedwireFundsIncoming
	if err := xml.Unmarshal(bs, &msg); err != nil {
		panic(err)
	}
	if err := msg.Validate(); err != nil {
		panic(err)
	}

	cct := msg.FedwireFundsIncomingMessage.FedwireFundsCustomerCreditTransfer
	fmt.Println("MsgDefIdr:", cct.AppHdr.MsgDefIdr)
	fmt.Println("MsgId:", cct.Document.FIToFICstmrCdtTrf.GrpHdr.MsgId)

	txn := cct.Document.FIToFICstmrCdtTrf.CdtTrfTxInf[0]
	fmt.Printf("Amount: %.2f %s\n", float64(txn.IntrBkSttlmAmt.Value), txn.IntrBkSttlmAmt.Ccy)

	// Output:
	// MsgDefIdr: pacs.008.001.08
	// MsgId: 20240228B1QDRCQR000001
	// Amount: 500000.00 USD
}

// Example_writeMessage builds a Fedwire Funds financial institution credit transfer (pacs.009),
// validates it, and writes it as XML.
func Example_writeMessage() {
	now := time.Date(2026, time.November, 20, 9, 0, 0, 0, time.UTC)

	// Identify the sender and receiver with their ABA routing numbers.
	agent := func(routingNumber string) pacs_009_001_08.BranchAndFinancialInstitutionIdentification6 {
		cd := pacs_009_001_08.ExternalClearingSystemIdentification1Code("USABA")
		return pacs_009_001_08.BranchAndFinancialInstitutionIdentification6{
			FinInstnId: pacs_009_001_08.FinancialInstitutionIdentification18{
				ClrSysMmbId: &pacs_009_001_08.ClearingSystemMemberIdentification2{
					ClrSysId: &pacs_009_001_08.ClearingSystemIdentification2Choice{Cd: &cd},
					MmbId:    pacs_009_001_08.Max35Text(routingNumber),
				},
			},
		}
	}
	party := func(routingNumber string) head_001_001_03.Party44Choice {
		return head_001_001_03.Party44Choice{
			FIId: &head_001_001_03.BranchAndFinancialInstitutionIdentification6{
				FinInstnId: head_001_001_03.FinancialInstitutionIdentification18{
					ClrSysMmbId: &head_001_001_03.ClearingSystemMemberIdentification2{
						MmbId: head_001_001_03.Max35Text(routingNumber),
					},
				},
			},
		}
	}

	// The IMAD identifies the message: cycle date, endpoint ID, and sequence number.
	imad := fedwire.InputMessageAccountabilityData("2026-11-20", "B1QDRCQR", 1)
	clearingSystem := pacs_009_001_08.ExternalCashClearingSystem1Code("FDW")
	localInstrument := pacs_009_001_08.Max35Text("FDWC")
	instgAgt, instdAgt := agent("011000015"), agent("021151080")
	uetr := pacs_009_001_08.UUIDv4Identifier("8a562c67-ca16-48ba-b074-65581be6f011")
	settlementDate := fedwire.ISODate(now)

	msg := fedwirefunds_incoming.FedwireFundsIncoming{
		FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
			FedwireFundsFinancialInstitutionCreditTransfer: &fedwirefunds_incoming.FedwireFundsFinancialInstitutionCreditTransfer{
				AppHdr: head_001_001_03.BusinessApplicationHeaderV03{
					Fr:        party("011000015"),
					To:        party("021151080"),
					BizMsgIdr: head_001_001_03.Max35Text(imad),
					MsgDefIdr: "pacs.009.001.08",
					CreDt:     fedwire.ISODateTime(now),
				},
				Document: pacs_009_001_08.Document{
					FICdtTrf: pacs_009_001_08.FinancialInstitutionCreditTransferV08{
						GrpHdr: pacs_009_001_08.GroupHeader93{
							MsgId:   pacs_009_001_08.Max35Text(imad),
							CreDtTm: fedwire.ISODateTime(now),
							NbOfTxs: "1",
							SttlmInf: pacs_009_001_08.SettlementInstruction7{
								SttlmMtd: pacs_009_001_08.SettlementMethod1CodeClrg,
								ClrSys: &pacs_009_001_08.ClearingSystemIdentification3Choice{
									Cd: &clearingSystem,
								},
							},
						},
						CdtTrfTxInf: []pacs_009_001_08.CreditTransferTransaction36{
							{
								PmtId: pacs_009_001_08.PaymentIdentification7{
									EndToEndId: "E2E-0001",
									UETR:       &uetr,
								},
								PmtTpInf: &pacs_009_001_08.PaymentTypeInformation28{
									LclInstrm: &pacs_009_001_08.LocalInstrument2Choice{Prtry: &localInstrument},
								},
								IntrBkSttlmAmt: pacs_009_001_08.ActiveCurrencyAndAmount{
									Value: 1250.75,
									Ccy:   "USD",
								},
								IntrBkSttlmDt: &settlementDate,
								InstgAgt:      &instgAgt,
								InstdAgt:      &instdAgt,
								Dbtr:          agent("011000015"),
								Cdtr:          agent("021151080"),
							},
						},
					},
				},
			},
		},
	}
	if err := msg.Validate(); err != nil {
		panic(err)
	}

	bs, err := xml.Marshal(msg)
	if err != nil {
		panic(err)
	}
	out := string(bs)
	fmt.Println(strings.Contains(out, `<FedwireFundsIncoming xmlns="urn:fedwirefunds:incoming:v001">`))
	fmt.Println(strings.Contains(out, `<IntrBkSttlmAmt Ccy="USD">1250.75</IntrBkSttlmAmt>`))

	// Output:
	// true
	// true
}
