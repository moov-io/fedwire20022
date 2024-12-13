package fedwire_test

import (
	"encoding/xml"
	"testing"

	"github.com/moov-io/fedwire20022/pkg/fedwire"

	"github.com/stretchr/testify/require"
)

func TestAmountFormat(t *testing.T) {
	var amt = fedwire.Amount(634)
	var amtTag = "<Amount>634.00</Amount>"

	out, err := amt.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "634.00", string(out))

	out, err = xml.Marshal(amt)
	require.NoError(t, err)
	require.Equal(t, amtTag, string(out))

	var read fedwire.Amount
	err = xml.Unmarshal([]byte(amtTag), &read)
	require.NoError(t, err)
	require.Equal(t, amt, read)

	t.Run("large", func(t *testing.T) {
		amt = fedwire.Amount(1_252_363.25)
		bs, err := xml.Marshal(amt)
		require.NoError(t, err)
		require.Equal(t, "<Amount>1252363.25</Amount>", string(bs))
	})
}

func TestAmountValidate(t *testing.T) {
	var amt = fedwire.Amount(634)
	require.NoError(t, amt.Validate())
}
