package currency

import (
	"errors"
	"time"
)

type Converter struct {
	Base  Currency
	Rates map[Currency]float64
	Time  time.Time

	ratescr []Currency
}

func (cn *Converter) ConvertTo(a Amount, tocr Currency) (Amount, error) {
	var ok bool

	arate := 1.0
	if a.Currency != cn.Base {
		arate, ok = cn.Rates[a.Currency]
		if !ok {
			return Amount{}, errors.New("undefined amount currency rate")
		}
	}

	abase := a.Sum / arate

	crate := 1.0
	if tocr != cn.Base {
		crate, ok = cn.Rates[tocr]
		if !ok {
			return Amount{}, errors.New("undefined convert currency rate")
		}
	}

	return Amount{abase * crate, tocr}, nil
}

func (cn *Converter) ConvertToSpec(a Amount, tocrs []Currency) (map[Currency]Amount, error) {
	res := make(map[Currency]Amount)
	for _, cr := range tocrs {
		ca, err := cn.ConvertTo(a, cr)
		if err != nil {
			return nil, err
		}

		res[cr] = ca
	}

	return res, nil
}

func (cn *Converter) ConvertToAll(a Amount) (map[Currency]Amount, error) {
	if cn.ratescr == nil {
		var rates []Currency
		for cr := range cn.Rates {
			rates = append(rates, cr)
		}
		cn.ratescr = rates
	}

	return cn.ConvertToSpec(a, cn.ratescr)
}

type Amount struct {
	Sum      float64
	Currency Currency
}

type MarginAmount struct {
	Amount

	Percent    float64
	MarginSum  float64
	MarginDiff float64
}

type MarginConverter struct {
	*Converter
	DefaultPercent float64
}

func (self *MarginConverter) calcMargin(a Amount, percent float64) MarginAmount {
	if percent == 0 {
		percent = self.DefaultPercent
	}

	ma := MarginAmount{
		Amount:  a,
		Percent: percent,
	}

	ma.MarginDiff = a.Sum * percent / 100
	ma.MarginSum = a.Sum - ma.MarginDiff

	return ma
}

func (self *MarginConverter) calcMargins(amounts map[Currency]Amount, percent float64) map[Currency]MarginAmount {
	res := make(map[Currency]MarginAmount)

	for cur, a := range amounts {
		res[cur] = self.calcMargin(a, percent)
	}

	return res
}

func (self *MarginConverter) ConvertTo(a Amount, tocr Currency, marginPercent float64) (MarginAmount, error) {
	amount, err := self.Converter.ConvertTo(a, tocr)
	if err != nil {
		return MarginAmount{}, err
	}

	return self.calcMargin(amount, marginPercent), nil
}

func (self *MarginConverter) ConvertToSpec(a Amount, tocrs []Currency, marginPercent float64) (map[Currency]MarginAmount, error) {
	amounts, err := self.Converter.ConvertToSpec(a, tocrs)
	if err != nil {
		return nil, err
	}

	return self.calcMargins(amounts, marginPercent), nil
}

func (self *MarginConverter) ConvertToAll(a Amount, marginPercent float64) (map[Currency]MarginAmount, error) {
	amounts, err := self.Converter.ConvertToAll(a)
	if err != nil {
		return nil, err
	}

	return self.calcMargins(amounts, marginPercent), nil
}
