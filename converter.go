package currency

import (
	"errors"
)

type Converter struct {
	Base  Currency
	Rates map[Currency]float64
	
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
	if cr != cn.Base {
		crate, ok = cn.Rates[cr]
		if !ok {
			return Amount{}, errors.New("undefined convert currency rate")
		}
	}
	
	return Amount{abase * crate, cr}, nil
}

func (cn *Converter) ConvertToSpec(a Amount, tocrs []Currency) (map[Currency]Amount, error) {
	res := make(map[Currency]Amount)
	for _, cr := range tocrs {
		ca, err := cn.ConvertTo(a)
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
		for cr, _ := range cn.Rates {
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
