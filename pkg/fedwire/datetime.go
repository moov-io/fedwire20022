package fedwire

import (
	"time"
)

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	tt, err := time.Parse("15:04:05Z", string(text))
	if err == nil {
		*t = ISOTime(tt)
		return nil
	}

	tt, err = time.Parse("15:04:05-07:00", string(text))
	if err == nil {
		*t = ISOTime(tt)
		return nil
	}

	return err
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
	out := time.Time(t).UTC().Format("2006-01-02Z")
	return []byte(out), nil
}

func (t ISODate) Validate() error {
	_, err := t.MarshalText()
	return err
}

func (t ISODate) String() string {
	return time.Time(t).Format(time.RFC3339)
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	tt, err := time.Parse("2006-01-02T15:04:05Z", string(text))
	if err == nil {
		*t = ISODateTime(tt)
		return nil
	}

	tt, err = time.Parse("2006-01-02T15:04:05-07:00", string(text))
	if err == nil {
		*t = ISODateTime(tt)
		return nil
	}

	return err
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
