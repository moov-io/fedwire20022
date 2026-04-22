package tests

import (
	"encoding/xml"
	"testing"

	"github.com/moov-io/fedwire20022/gen/fedwirefunds_incoming"
	"github.com/moov-io/fedwire20022/gen/head_001_001_03"
	"github.com/moov-io/fedwire20022/gen/pacs_008_001_08"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsIncoming_Marshal(t *testing.T) {
	message := fedwirefunds_incoming.FedwireFundsIncoming{
		FedwireFundsIncomingMessage: fedwirefunds_incoming.FedwireFundsIncomingMessage{
			FedwireFundsCustomerCreditTransfer: &fedwirefunds_incoming.FedwireFundsCustomerCreditTransfer{
				AppHdr:   head_001_001_03.BusinessApplicationHeaderV03{},
				Document: pacs_008_001_08.Document{},
			},
		},
	}

	bs, err := xml.Marshal(message)
	require.NoError(t, err)

	output := string(bs)
	require.Contains(t, output, `<FedwireFundsIncoming xmlns="urn:fedwirefunds:incoming:v001">`)
	require.Contains(t, output, `<FedwireFundsIncomingMessage><FedwireFundsCustomerCreditTransfer>`)
	require.Contains(t, output, `<AppHdr xmlns="urn:iso:std:iso:20022:tech:xsd:head.001.001.03">`)
	require.Contains(t, output, `<Fr></Fr><To></To><BizMsgIdr></BizMsgIdr><MsgDefIdr></MsgDefIdr>`)
	require.Contains(t, output, `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08"><FIToFICstmrCdtTrf><GrpHdr>`)
	require.NotContains(t, output, `xmlns=""`)
}
