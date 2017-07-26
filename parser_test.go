package currency

import (
	"testing"
	"time"
)

func TestFixerIoParser_Parse(t *testing.T) {
	p := FixerIoParser{}

	dt := time.Date(2000, 1, 3, 0, 0, 0, 0, time.Local)
	c, err := p.Parse(dt, EUR)
	if err != nil {
		t.Fatal(err)
	}

	if c.Rates[KRW] != 1140 {
		t.Fatal("invalid currecy by date")
	}
}
