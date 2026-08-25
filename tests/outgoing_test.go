package tests

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/fedwire20022/gen/fedwirefunds_outgoing"
	"github.com/moov-io/fedwire20022/gen/head_001_001_03"
	"github.com/moov-io/fedwire20022/gen/pacs_008_001_08"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsOutgoing_Marshal(t *testing.T) {
	message := fedwirefunds_outgoing.FedwireFundsOutgoing{
		FedwireFundsOutgoingMessage: fedwirefunds_outgoing.FedwireFundsOutgoingMessage{
			FedwireFundsCustomerCreditTransfer: &fedwirefunds_outgoing.FedwireFundsCustomerCreditTransfer{
				AppHdr:   head_001_001_03.BusinessApplicationHeaderV03{},
				Document: pacs_008_001_08.Document{},
			},
		},
	}

	bs, err := xml.Marshal(message)
	require.NoError(t, err)

	output := string(bs)
	require.Contains(t, output, `<FedwireFundsOutgoing xmlns="urn:fedwirefunds:outgoing:v001">`)
	require.Contains(t, output, `<FedwireFundsOutgoingMessage><FedwireFundsCustomerCreditTransfer>`)
	require.Contains(t, output, `<AppHdr xmlns="urn:iso:std:iso:20022:tech:xsd:head.001.001.03">`)
	require.Contains(t, output, `<Fr></Fr><To></To><BizMsgIdr></BizMsgIdr><MsgDefIdr></MsgDefIdr><CreDt>0001-01-01T00:00:00Z</CreDt>`)
	require.Contains(t, output, `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08"><FIToFICstmrCdtTrf><GrpHdr>`)
	require.NotContains(t, output, `xmlns=""`)
}

func TestFedwireFundsOutgoing_UnmarshalAccountBalanceReportMaster(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("..", "testdata", "CAMT052_AccountBalanceReport_Master.xml"))
	require.NoError(t, err)

	var doc fedwirefunds_outgoing.FedwireFundsOutgoing
	require.NoError(t, xml.Unmarshal(bs, &doc))

	rpt := doc.FedwireFundsOutgoingMessage.FedwireFundsAccountBalanceReportMaster
	require.NotNil(t, rpt, "FedwireFundsAccountBalanceReport_Master must bind (XSD underscore)")
	require.Equal(t, "camt.052.001.08", string(rpt.AppHdr.MsgDefIdr))
	require.Equal(t, "frb.fedwire.abm.01", string(rpt.AppHdr.MktPrctc.Id))
	require.Equal(t, "ABAR", string(rpt.Document.BkToCstmrAcctRpt.GrpHdr.MsgId))
	require.NotEmpty(t, rpt.Document.BkToCstmrAcctRpt.Rpt)
	require.Equal(t, "ABMS", string(rpt.Document.BkToCstmrAcctRpt.Rpt[0].Id))
}
