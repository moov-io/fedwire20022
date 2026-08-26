package fedwire

import (
	"fmt"
	"time"
)

// ISO 20022 ISOTime is xs:time. Fractional seconds and timezone are optional.
// Go still accepts a fractional second when the layout omits one.
var isoTimeLayouts = []string{
	"15:04:05Z07:00", // 13:15:05Z, 08:15:05-05:00, 10:51:23.123Z
	"15:04:05Z0700",  // 08:15:05-0500
	"15:04:05",       // 10:51:23.123 (no offset)
}

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	tt, err := parseLayouts(text, isoTimeLayouts)
	if err != nil {
		return fmt.Errorf("parsing ISOTime %q: %w", text, err)
	}
	*t = ISOTime(tt)
	return nil
}

func (t ISOTime) MarshalText() ([]byte, error) {
	out := time.Time(t).UTC().Format("15:04:05Z")
	return []byte(out), nil
}

func (t ISOTime) Validate() error {
	_, err := t.MarshalText()
	return err
}

func (t ISOTime) String() string {
	return time.Time(t).Format(time.RFC3339)
}

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	tt, err := time.Parse("2006-01-02Z", string(text))
	if err == nil {
		*t = ISODate(tt)
		return nil
	}

	tt, err = time.Parse("2006-01-02", string(text))
	if err == nil {
		*t = ISODate(tt)
		return nil
	}

	tt, err = time.Parse("2006-01-02-07:00", string(text))
	if err == nil {
		*t = ISODate(tt)
		return nil
	}

	return err
}

func (t ISODate) MarshalText() ([]byte, error) {
	out := time.Time(t).UTC().Format("2006-01-02")
	return []byte(out), nil
}

func (t ISODate) Validate() error {
	_, err := t.MarshalText()
	return err
}

func (t ISODate) String() string {
	return time.Time(t).Format(time.RFC3339)
}

// ISO 20022 ISODateTime is xs:dateTime. Fractional seconds and timezone are optional.
var isoDateTimeLayouts = []string{
	"2006-01-02T15:04:05Z07:00",
	"2006-01-02T15:04:05Z0700",
	"2006-01-02T15:04:05",
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	tt, err := parseLayouts(text, isoDateTimeLayouts)
	if err != nil {
		return fmt.Errorf("parsing ISODateTime %q: %w", text, err)
	}
	*t = ISODateTime(tt)
	return nil
}

func (t ISODateTime) MarshalText() ([]byte, error) {
	return time.Time(t).MarshalText()
}

func (t ISODateTime) Validate() error {
	_, err := t.MarshalText()
	return err
}

func (t ISODateTime) String() string {
	return time.Time(t).Format(time.RFC3339)
}

func parseLayouts(text []byte, layouts []string) (time.Time, error) {
	s := string(text)
	var firstErr error
	for _, layout := range layouts {
		tt, err := time.Parse(layout, s)
		if err == nil {
			return tt, nil
		}
		if firstErr == nil {
			firstErr = err
		}
	}
	if firstErr == nil {
		return time.Time{}, fmt.Errorf("no layouts")
	}
	return time.Time{}, firstErr
}
