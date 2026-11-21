package tests

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/fedwire20022/gen/fedwirefunds_incoming"
	"github.com/moov-io/fedwire20022/gen/fedwirefunds_outgoing"
	"github.com/moov-io/fedwire20022/gen/pacs_009_001_08"
	"github.com/stretchr/testify/require"
)

func TestValidate_PatternsMatchWholeValue(t *testing.T) {
	require.NoError(t, pacs_009_001_08.Max15NumericText("1").Validate())
	require.Error(t, pacs_009_001_08.Max15NumericText("abc1").Validate())
	require.Error(t, pacs_009_001_08.Max15NumericText("12345678901234567890").Validate())

	require.NoError(t, pacs_009_001_08.ActiveCurrencyCode("USD").Validate())
	require.Error(t, pacs_009_001_08.ActiveCurrencyCode("USDX").Validate())

	uetr := "8a562c67-ca16-48ba-b074-65581be6f001"
	require.NoError(t, pacs_009_001_08.UUIDv4Identifier(uetr).Validate())
	require.Error(t, pacs_009_001_08.UUIDv4Identifier("zzz"+uetr+"zzz").Validate())
}

func TestValidate_TestdataSamples(t *testing.T) {
	t.Run("incoming", func(t *testing.T) {
		bs, err := os.ReadFile(filepath.Join("..", "testdata", "PACS008_CustomerCreditTransfer.txt"))
		require.NoError(t, err)

		var msg fedwirefunds_incoming.FedwireFundsIncoming
		require.NoError(t, xml.Unmarshal(bs, &msg))
		require.NoError(t, msg.Validate())

		// Junk around the number of transactions must fail
		cct := msg.FedwireFundsIncomingMessage.FedwireFundsCustomerCreditTransfer
		cct.Document.FIToFICstmrCdtTrf.GrpHdr.NbOfTxs = "1x"
		require.ErrorContains(t, msg.Validate(), "NbOfTxs")
	})

	t.Run("outgoing", func(t *testing.T) {
		bs, err := os.ReadFile(filepath.Join("..", "testdata", "CAMT052_AccountBalanceReport_Master.xml"))
		require.NoError(t, err)

		var msg fedwirefunds_outgoing.FedwireFundsOutgoing
		require.NoError(t, xml.Unmarshal(bs, &msg))
		require.NoError(t, msg.Validate())
	})
}
