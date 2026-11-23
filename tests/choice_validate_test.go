package tests

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/fedwire20022/gen/fedwirefunds_incoming"
	"github.com/moov-io/fedwire20022/gen/fedwirefunds_outgoing"
	"github.com/moov-io/fedwire20022/gen/head_001_001_03"
	"github.com/stretchr/testify/require"
)

func TestValidate_ChoiceSamples(t *testing.T) {
	t.Run("incoming", func(t *testing.T) {
		bs, err := os.ReadFile(filepath.Join("..", "testdata", "PACS008_CustomerCreditTransfer.txt"))
		require.NoError(t, err)

		var msg fedwirefunds_incoming.FedwireFundsIncoming
		require.NoError(t, xml.Unmarshal(bs, &msg))
		require.NoError(t, msg.Validate())

		// Two messages in one choice must fail
		msg.FedwireFundsIncomingMessage.FedwireFundsConnectionCheck = &fedwirefunds_incoming.FedwireFundsConnectionCheck{}
		err = msg.Validate()
		require.ErrorContains(t, err, "FedwireFundsIncomingMessage: requires exactly one of")
		require.ErrorContains(t, err, "found 2")
	})

	t.Run("outgoing", func(t *testing.T) {
		bs, err := os.ReadFile(filepath.Join("..", "testdata", "CAMT052_AccountBalanceReport_Master.xml"))
		require.NoError(t, err)

		var msg fedwirefunds_outgoing.FedwireFundsOutgoing
		require.NoError(t, xml.Unmarshal(bs, &msg))
		require.NoError(t, msg.Validate())
	})
}

func TestValidate_EmptyIncomingMessage(t *testing.T) {
	var msg fedwirefunds_incoming.FedwireFundsIncoming
	err := msg.Validate()
	require.ErrorContains(t, err, "FedwireFundsIncomingMessage: requires exactly one of (FedwireFundsCustomerCreditTransfer, ")
}

func TestValidate_EmptyPartyChoice(t *testing.T) {
	var party head_001_001_03.Party44Choice
	require.ErrorContains(t, party.Validate(), "Party44Choice: requires exactly one of (OrgId, FIId), found 0")

	party.FIId = &head_001_001_03.BranchAndFinancialInstitutionIdentification6{}
	err := party.Validate()
	if err != nil {
		require.NotContains(t, err.Error(), "requires exactly one of (OrgId, FIId)")
	}

	party.OrgId = &head_001_001_03.PartyIdentification135{}
	require.ErrorContains(t, party.Validate(), "requires exactly one of (OrgId, FIId), found 2")
}
