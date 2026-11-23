package tests

import (
	"encoding/xml"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/fedwire20022/gen/fedwirefunds_incoming"
	"github.com/moov-io/fedwire20022/gen/pacs_008_001_08"
	"github.com/moov-io/fedwire20022/gen/pacs_009_001_08"
	"github.com/stretchr/testify/require"
)

func readCustomerCreditTransfer(t *testing.T) fedwirefunds_incoming.FedwireFundsIncoming {
	t.Helper()

	bs, err := os.ReadFile(filepath.Join("..", "testdata", "PACS008_CustomerCreditTransfer.txt"))
	require.NoError(t, err)

	var msg fedwirefunds_incoming.FedwireFundsIncoming
	require.NoError(t, xml.Unmarshal(bs, &msg))
	require.NotNil(t, msg.FedwireFundsIncomingMessage.FedwireFundsCustomerCreditTransfer)

	return msg
}

func TestValidate_SettlementAmount(t *testing.T) {
	msg := readCustomerCreditTransfer(t)
	require.NoError(t, msg.Validate())

	cases := map[string]float64{
		"negative":     -5,
		"NaN":          math.NaN(),
		"infinity":     math.Inf(1),
		"total digits": 1e20,
	}
	for name, amount := range cases {
		t.Run(name, func(t *testing.T) {
			msg := readCustomerCreditTransfer(t)
			cct := msg.FedwireFundsIncomingMessage.FedwireFundsCustomerCreditTransfer
			cct.Document.FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Value = pacs_008_001_08.ActiveCurrencyAndAmountSimpleType(amount)

			err := msg.Validate()
			require.ErrorContains(t, err, "ActiveCurrencyAndAmount.Value")
		})
	}

	t.Run("valid", func(t *testing.T) {
		msg := readCustomerCreditTransfer(t)
		cct := msg.FedwireFundsIncomingMessage.FedwireFundsCustomerCreditTransfer
		cct.Document.FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Value = 1e15 // 18 digits with two decimals
		require.NoError(t, msg.Validate())
	})
}

func TestValidate_AmountSimpleType(t *testing.T) {
	require.NoError(t, pacs_009_001_08.ActiveCurrencyAndAmountSimpleType(0).Validate())
	require.NoError(t, pacs_009_001_08.ActiveCurrencyAndAmountSimpleType(1234.56).Validate())

	require.Error(t, pacs_009_001_08.ActiveCurrencyAndAmountSimpleType(-5).Validate())
	require.Error(t, pacs_009_001_08.ActiveCurrencyAndAmountSimpleType(math.NaN()).Validate())
	require.Error(t, pacs_009_001_08.ActiveCurrencyAndAmountSimpleType(1e20).Validate())

	amt := pacs_009_001_08.ActiveCurrencyAndAmount{
		Value: pacs_009_001_08.ActiveCurrencyAndAmountSimpleType(-5),
		Ccy:   "USD",
	}
	require.ErrorContains(t, amt.Validate(), "minInclusive")
}
