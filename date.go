package pflagdate

import (
	date "github.com/hardfinhq/go-date"
)

type Date struct {
	date.Date
}

// Implement pflag.Value interface

// Type implements pflag.Value.Type().
// Returns "Date".
func (*Date) Type() string {
	return "Date"
}

// String implements pflag.Value.String().
// Returns date formatted as time.DateOnly.
func (d *Date) String() string {
	return d.Date.String()
}

// Set implements pflag.Value.Set().
func (d *Date) Set(value string) error {
	parsed, err := date.FromString(value)
	if err != nil {
		return err
	}
	d.Date = parsed
	return nil
}
