package fedwire_test

import (
	"encoding/xml"
	"testing"

	"github.com/moov-io/fedwire20022/pkg/fedwire"

	"github.com/stretchr/testify/require"
)

func TestISOTime(t *testing.T) {
	// UTC
	input := `<ISOTime>13:15:05Z</ISOTime>`

	var when fedwire.ISOTime
	err := xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "0000-01-01T13:15:05Z", when.String())

	bs, err := when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "13:15:05Z", string(bs))

	// Local with offset
	input = `<ISOTime>08:15:05-05:00</ISOTime>`

	err = xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "0000-01-01T08:15:05-05:00", when.String())

	bs, err = when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "13:15:05Z", string(bs))
}

func TestISODate(t *testing.T) {
	// UTC
	input := `<ISODate>2026-02-11Z</ISODate>`

	var when fedwire.ISODate
	err := xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "2026-02-11T00:00:00Z", when.String())

	bs, err := when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2026-02-11Z", string(bs))

	// Local date format (YYYY-MM-DD)
	input = `<ISODate>2026-02-11</ISODate>`

	err = xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "2026-02-11T00:00:00Z", when.String())

	bs, err = when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2026-02-11Z", string(bs))

	// Local date with UTC offset format (YYYY-MM-DD+/-hh:mm)
	input = `<ISODate>2026-02-11-05:00</ISODate>`

	err = xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "2026-02-11T00:00:00-05:00", when.String())

	bs, err = when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2026-02-11Z", string(bs))

	// Local date with UTC offset format (YYYY-MM-DD+/-hh:mm)
	input = `<ISODate>2026-02-11+03:00</ISODate>`

	err = xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "2026-02-11T00:00:00+03:00", when.String())

	bs, err = when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2026-02-10Z", string(bs))
}

func TestISODateTime(t *testing.T) {
	// UTC
	input := `<ISODateTime>2024-04-02T13:15:05Z</ISODateTime>`

	var when fedwire.ISODateTime
	err := xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "2024-04-02T13:15:05Z", when.String())

	bs, err := when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2024-04-02T13:15:05Z", string(bs))

	// Local with offset
	input = `<ISODateTime>2024-04-02T08:15:05-05:00</ISODateTime>`

	err = xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "2024-04-02T08:15:05-05:00", when.String())

	bs, err = when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2024-04-02T08:15:05-05:00", string(bs))
}
