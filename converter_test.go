package currency

import (
	"testing"
)

func TestConverter_ConvertTo(t *testing.T) {
	cn := &Converter{
		Base: EUR,
		Rates: map[Currency]float64{
			RUB: 65,
			USD: 1.2,
			INR: 70,
		},
	}

	usdTest, err := cn.ConvertTo(Amount{70, INR}, USD)
	if err != nil {
		t.Error(err)
	} else if usdTest.Sum != 1.2 {
		t.Error("invalid usd sum")
	}

	eurTest, err := cn.ConvertTo(Amount{70, INR}, EUR)
	if err != nil {
		t.Error(err)
	} else if eurTest.Sum != 1 {
		t.Error("invalid eur sum")
	}
}

func TestMarginConverter_ConvertTo(t *testing.T) {
	cn := &Converter{
		Base: EUR,
		Rates: map[Currency]float64{
			RUB: 65,
			USD: 1.2,
			INR: 70,
		},
	}

	mcn := &MarginConverter{cn, 10}

	usdTest, err := mcn.ConvertTo(Amount{70, INR}, USD, 20)

	if err != nil {
		t.Error(err)
	} else if usdTest.Sum != 1.2 {
		t.Error("invalid usd sum")
	} else if usdTest.MarginDiff != 0.24 || usdTest.MarginSum != 0.96 {
		t.Error("invalid usd margin sum")
	}

	eurTest, err := mcn.ConvertTo(Amount{70, INR}, EUR, 0)
	if err != nil {
		t.Error(err)
	} else if eurTest.Sum != 1 {
		t.Error("invalid eur sum")
	} else if eurTest.MarginDiff != 0.1 || eurTest.MarginSum != 0.9 {
		t.Error("invalid eur margin sum")
	}
}
