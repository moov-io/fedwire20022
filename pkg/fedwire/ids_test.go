package fedwire_test

import (
	"strings"
	"testing"

	"github.com/moov-io/fedwire20022/pkg/fedwire"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestBusinessMessageID(t *testing.T) {
	id := fedwire.BusinessMessageID("a_")
	require.Len(t, id, 34)
	require.True(t, strings.HasPrefix(id, "a_"))
	require.NoError(t, fedwire.ValidBusinessMessageID(id))

	id = fedwire.BusinessMessageID("ct_")
	require.Len(t, id, 35)
	require.True(t, strings.HasPrefix(id, "ct_"))
	require.NoError(t, fedwire.ValidBusinessMessageID(id))

	id = fedwire.BusinessMessageID("crdt_")
	require.Len(t, id, 35)
	require.True(t, strings.HasPrefix(id, "crd")) // prefix is chopped off
	require.NoError(t, fedwire.ValidBusinessMessageID(id))
}

func TestBusinessMessageIDHash(t *testing.T) {
	id := fedwire.BusinessMessageIDHash("ct_", "abc123", "xyz987")
	require.Len(t, id, 35)
	require.True(t, strings.HasPrefix(id, "ct_"))
	require.NoError(t, fedwire.ValidBusinessMessageID(id))
	require.Equal(t, "ct_d99fe29b70a65320ac08d857b2b7af62", id)
}

func TestValidBusinessMessageID(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		inputs := []string{
			"abc123 456\\8*9_ZEF~u@[qq]$<!!>",
			strings.ReplaceAll(uuid.NewString(), "-", ""), // UUIDs are 36 characters with dashes
		}
		for _, input := range inputs {
			t.Run(input, func(t *testing.T) {
				err := fedwire.ValidBusinessMessageID(input)
				require.NoError(t, err)
			})
		}
	})

	t.Run("invalid", func(t *testing.T) {
		cases := []struct {
			input    string
			expected string
		}{
			{
				input:    uuid.NewString(),
				expected: "business message id must be 1-35 characters long - was 36",
			},
			{
				input:    `©®`,
				expected: "unexpected © in business message id",
			},
			{
				input:    "L’Allier",
				expected: "unexpected ’ in business message id",
			},
		}
		for _, tc := range cases {
			t.Run(tc.input, func(t *testing.T) {
				err := fedwire.ValidBusinessMessageID(tc.input)
				require.Error(t, err)
				require.ErrorContains(t, err, tc.expected)
			})
		}
	})
}
