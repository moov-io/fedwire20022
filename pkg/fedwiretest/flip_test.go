package fedwiretest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlipMessageDirection(t *testing.T) {
	t.Run("PACS008_CustomerCreditTransfer.txt", func(t *testing.T) {
		// Test with sample incoming message
		bs, err := os.ReadFile(filepath.Join("..", "..", "testdata", "PACS008_CustomerCreditTransfer.txt"))
		require.NoError(t, err)

		// This is an incoming message, flipping should make it outgoing
		flipped, err := FlipMessageDirection(bs)
		require.NoError(t, err)
		require.NotEqual(t, string(bs), string(flipped)) // Should be different
		require.Contains(t, string(flipped), "FedwireFundsOutgoing")
		require.Contains(t, string(flipped), "FedwireFundsCustomerCreditTransfer")

		// Flip back
		flippedBack, err := FlipMessageDirection(flipped)
		require.NoError(t, err)
		require.Contains(t, string(flippedBack), "FedwireFundsIncoming")
		require.Contains(t, string(flippedBack), "FedwireFundsCustomerCreditTransfer")
	})
}
