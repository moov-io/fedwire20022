package fedwire_test

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/fedwire20022/pkg/fedwire"

	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/require"
)

func TestISODateFormat(t *testing.T) {
	when := civil.Date{
		Year:  2019,
		Month: time.March,
		Day:   21,
	}

	require.Equal(t, fedwire.ISODate(when), fedwire.UnmarshalISODate("2019-03-21"))

	out, err := fedwire.ISODate(when).MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2019-03-21", string(out))

	out, err = xml.Marshal(fedwire.ISODate(when))
	require.NoError(t, err)
	require.Equal(t, "<ISODate>2019-03-21</ISODate>", string(out))

	var read fedwire.ISODate
	err = xml.Unmarshal([]byte("<ISODate>2019-03-21</ISODate>"), &read)
	require.NoError(t, err)
	require.True(t, when == (civil.Date)(read))
}

func TestISODateTimeFormat(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	when := time.Date(2019, time.March, 21, 10, 36, 19, 0, loc)

	require.Equal(t, fedwire.ISODateTime(when), fedwire.UnmarshalISODateTime("2019-03-21T10:36:19"))

	out, err := fedwire.ISODateTime(when).MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2019-03-21T10:36:19", string(out))

	out, err = xml.Marshal(fedwire.ISODateTime(when))
	require.NoError(t, err)
	require.Equal(t, "<ISODateTime>2019-03-21T10:36:19</ISODateTime>", string(out))

	var read fedwire.ISODateTime
	err = xml.Unmarshal([]byte("<ISODateTime>2019-03-21T10:36:19</ISODateTime>"), &read)
	require.NoError(t, err)
	require.True(t, when.Equal(time.Time(read)))
}
