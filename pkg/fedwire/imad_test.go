package fedwire_test

import (
	"fmt"
	"testing"

	"github.com/moov-io/fedwire20022/pkg/fedwire"

	"github.com/stretchr/testify/require"
)

func TestInputMessageAccountabilityData(t *testing.T) {
	cases := []struct {
		cycleDate      string
		endpointID     string
		sequenceNumber int

		expected string
	}{
		{
			cycleDate:      "20060310",
			endpointID:     "87654321",
			sequenceNumber: 245,
			expected:       "2006031087654321000245",
		},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("%s %s %d", tc.cycleDate, tc.endpointID, tc.sequenceNumber)

		t.Run(name, func(t *testing.T) {
			got := fedwire.InputMessageAccountabilityData(tc.cycleDate, tc.endpointID, tc.sequenceNumber)
			require.Equal(t, tc.expected, got)
			require.NoError(t, fedwire.ValidIMAD(got))
		})
	}
}
