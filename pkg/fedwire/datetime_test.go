package fedwire_test

import (
	"encoding/xml"
	"testing"
	"time"

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

func TestISOTime_Layouts(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "utc", input: "13:15:05Z", want: "0000-01-01T13:15:05Z"},
		{name: "offset", input: "08:15:05-05:00", want: "0000-01-01T08:15:05-05:00"},
		{name: "compact offset", input: "08:15:05-0500", want: "0000-01-01T08:15:05-05:00"},
		{name: "no timezone", input: "10:51:23", want: "0000-01-01T10:51:23Z"},
		{name: "fraction no timezone", input: "10:51:23.123", want: "0000-01-01T10:51:23Z"},
		{name: "fraction utc", input: "10:51:23.123Z", want: "0000-01-01T10:51:23Z"},
		{name: "fraction offset", input: "10:51:23.123-05:00", want: "0000-01-01T10:51:23-05:00"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var when fedwire.ISOTime
			err := xml.Unmarshal([]byte("<ISOTime>"+tc.input+"</ISOTime>"), &when)
			require.NoError(t, err)
			require.Equal(t, tc.want, when.String())
		})
	}
}

func TestISOTime_FractionWithoutOffset(t *testing.T) {
	// Fed TODI messages send xs:time as hh:mm:ss.sss with no timezone.
	var when fedwire.ISOTime
	err := xml.Unmarshal([]byte("<ISOTime>10:51:23.123</ISOTime>"), &when)
	require.NoError(t, err)

	got := time.Time(when)
	require.Equal(t, 10, got.Hour())
	require.Equal(t, 51, got.Minute())
	require.Equal(t, 23, got.Second())
	require.Equal(t, 123000000, got.Nanosecond())
}

func TestISOTime_Invalid(t *testing.T) {
	var when fedwire.ISOTime
	err := xml.Unmarshal([]byte("<ISOTime>not-a-time</ISOTime>"), &when)
	require.Error(t, err)
	require.Contains(t, err.Error(), `parsing ISOTime "not-a-time"`)
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
	require.Equal(t, "2026-02-11", string(bs))

	// Local date format (YYYY-MM-DD)
	input = `<ISODate>2026-02-11</ISODate>`

	err = xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "2026-02-11T00:00:00Z", when.String())

	bs, err = when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2026-02-11", string(bs))

	// Local date with UTC offset format (YYYY-MM-DD+/-hh:mm)
	input = `<ISODate>2026-02-11-05:00</ISODate>`

	err = xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "2026-02-11T00:00:00-05:00", when.String())

	bs, err = when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2026-02-11", string(bs))

	// Local date with UTC offset format (YYYY-MM-DD+/-hh:mm)
	input = `<ISODate>2026-02-11+03:00</ISODate>`

	err = xml.Unmarshal([]byte(input), &when)
	require.NoError(t, err)
	require.Equal(t, "2026-02-11T00:00:00+03:00", when.String())

	bs, err = when.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2026-02-10", string(bs))
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

func TestISODateTime_Layouts(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "utc", input: "2024-04-02T13:15:05Z", want: "2024-04-02T13:15:05Z"},
		{name: "offset", input: "2024-04-02T08:15:05-05:00", want: "2024-04-02T08:15:05-05:00"},
		{name: "compact offset", input: "2024-04-02T08:15:05-0500", want: "2024-04-02T08:15:05-05:00"},
		{name: "no timezone", input: "2024-04-02T13:15:05", want: "2024-04-02T13:15:05Z"},
		{name: "fraction no timezone", input: "2024-04-02T13:15:05.123", want: "2024-04-02T13:15:05Z"},
		{name: "fraction utc", input: "2024-04-02T13:15:05.123Z", want: "2024-04-02T13:15:05Z"},
		{name: "fraction offset", input: "2026-08-25T11:31:30.822-04:00", want: "2026-08-25T11:31:30-04:00"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var when fedwire.ISODateTime
			err := xml.Unmarshal([]byte("<ISODateTime>"+tc.input+"</ISODateTime>"), &when)
			require.NoError(t, err)
			require.Equal(t, tc.want, when.String())
		})
	}
}

func TestISODateTime_Invalid(t *testing.T) {
	var when fedwire.ISODateTime
	err := xml.Unmarshal([]byte("<ISODateTime>not-a-datetime</ISODateTime>"), &when)
	require.Error(t, err)
	require.Contains(t, err.Error(), `parsing ISODateTime "not-a-datetime"`)
}
