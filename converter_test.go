package currency

import (
	"testing"
	"fmt"
)

func TestConverter_Convert(t *testing.T) {
	cn := &Converter{
		Base: EUR,
		Rates: map[Currency]float64{
			RUB: 65,
			USD: 1.2,
			INR: 70,
		},
	}
	
	fmt.Println(cn.ConvertTo(Amount{70, INR}, USD))
	fmt.Println(cn.ConvertTo(Amount{70, INR}, EUR))
	fmt.Println(cn.ConvertTo(Amount{70, INR}, RUB))
	fmt.Println(cn.ConvertTo(Amount{2, EUR}, RUB))
	fmt.Println(cn.ConvertTo(Amount{1, EUR}, EUR))
}
