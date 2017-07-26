package currency

import (
	"strings"
)

type Currency struct {
	Code string
	Num  string
}

//https://en.wikipedia.org/wiki/ISO_4217
var XXX Currency = Currency{}
var RUB Currency = Currency{"RUB", "643"}
var USD Currency = Currency{"USD", "840"}
var EUR Currency = Currency{"EUR", "978"}
var AUD Currency = Currency{"AUD", "036"}
var BGN Currency = Currency{"BGN", "975"}
var BRL Currency = Currency{"BRL", "986"}
var CAD Currency = Currency{"CAD", "124"}
var CHF Currency = Currency{"CHF", "756"}
var CNY Currency = Currency{"CNY", "156"}
var CZK Currency = Currency{"CZK", "203"}
var DKK Currency = Currency{"DKK", "208"}
var GBP Currency = Currency{"GBP", "826"}
var HKD Currency = Currency{"HKD", "344"}
var HRK Currency = Currency{"HRK", "191"}
var HUF Currency = Currency{"HUF", "348"}
var IDR Currency = Currency{"IDR", "360"}
var ILS Currency = Currency{"ILS", "376"}
var INR Currency = Currency{"INR", "356"}
var JPY Currency = Currency{"JPY", "392"}
var KRW Currency = Currency{"KRW", "410"}
var MXN Currency = Currency{"MXN", "484"}
var MYR Currency = Currency{"MYR", "458"}
var NOK Currency = Currency{"NOK", "578"}
var NZD Currency = Currency{"NZD", "554"}
var PHP Currency = Currency{"PHP", "608"}
var PLN Currency = Currency{"PLN", "985"}
var RON Currency = Currency{"RON", "946"}
var SEK Currency = Currency{"SEK", "752"}
var SGD Currency = Currency{"SGD", "702"}
var THB Currency = Currency{"THB", "764"}
var TRY Currency = Currency{"TRY", "949"}
var ZAR Currency = Currency{"ZAR", "710"}

var currencies = map[string]Currency{
	RUB.Code: RUB,
	RUB.Num:  RUB,
	
	USD.Code: USD,
	USD.Num:  USD,
	
	EUR.Code: EUR,
	EUR.Num:  EUR,
	
	AUD.Code: AUD,
	AUD.Num:  AUD,
	
	BGN.Code: BGN,
	BGN.Num:  BGN,
	
	BRL.Code: BRL,
	BRL.Num:  BRL,
	
	CAD.Code: CAD,
	CAD.Num:  CAD,
	
	CHF.Code: CHF,
	CHF.Num:  CHF,
	
	CNY.Code: CNY,
	CNY.Num:  CNY,
	
	CZK.Code: CZK,
	CZK.Num:  CZK,
	
	DKK.Code: DKK,
	DKK.Num:  DKK,
	
	GBP.Code: GBP,
	GBP.Num:  GBP,
	
	HKD.Code: HKD,
	HKD.Num:  HKD,
	
	HRK.Code: HRK,
	HRK.Num:  HRK,
	
	HUF.Code: HUF,
	HUF.Num:  HUF,
	
	IDR.Code: IDR,
	IDR.Num:  IDR,
	
	ILS.Code: ILS,
	ILS.Num:  ILS,
	
	INR.Code: INR,
	INR.Num:  INR,
	
	JPY.Code: JPY,
	JPY.Num:  JPY,
	
	KRW.Code: KRW,
	KRW.Num:  KRW,
	
	MXN.Code: MXN,
	MXN.Num:  MXN,
	
	MYR.Code: MYR,
	MYR.Num:  MYR,
	
	NOK.Code: NOK,
	NOK.Num:  NOK,
	
	NZD.Code: NZD,
	NZD.Num:  NZD,
	
	PHP.Code: PHP,
	PHP.Num:  PHP,
	
	PLN.Code: PLN,
	PLN.Num:  PLN,
	
	RON.Code: RON,
	RON.Num:  RON,
	
	SEK.Code: SEK,
	SEK.Num:  SEK,
	
	SGD.Code: SGD,
	SGD.Num:  SGD,
	
	THB.Code: THB,
	THB.Num:  THB,
	
	TRY.Code: TRY,
	TRY.Num:  TRY,
	
	ZAR.Code: ZAR,
	ZAR.Num:  ZAR,
}

func (cr Currency) String() string {
	return cr.Code
}

func Parse(s string) (cr Currency, ok bool) {
	cr, ok = currencies[strings.ToUpper(s)]
	return
}

func ParseDefault(s string, def Currency) Currency {
	cr, ok := Parse(s)
	if !ok {
		return def
	}
	
	return cr
}
