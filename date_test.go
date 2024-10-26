package pflagdate

import (
	"testing"
	"time"

	date "github.com/hardfinhq/go-date"
	"github.com/spf13/pflag"
)

func TestDate(t *testing.T) {
	dt := new(Date)

	if dt.Type() != "Date" {
		t.Errorf("unexpected type name %q", dt.Type())
	}

	flag := &pflag.Flag{
		Name:      "date",
		Shorthand: "d",
		Usage:     "YYYY-MM-DD date",
		Value:     dt,
	}

	pfs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	pfs.AddFlag(flag)

	err := pfs.Parse([]string{"--date", "2001-03-06"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !dt.Equal(date.NewDate(2001, time.March, 6)) {
		t.Errorf("unexpected date: %s", dt)
	}

	if dt.String() != "2001-03-06" {
		t.Errorf("unexpected date: %s", dt)
	}

	err2 := pfs.Parse([]string{"--date", "invalid"})
	if err2 == nil {
		t.Error("expected error, got nil")
	}
}
