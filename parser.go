package currency

import (
	"net/http"
	"time"
	"errors"
	"encoding/json"
)

const fixerIoApiUrl = "://api.fixer.io/"

type FixerIoParser struct {
	Client *http.Client
	UseTls bool
}

func (self *FixerIoParser) getClient() *http.Client {
	if self.Client == nil {
		return http.DefaultClient
	}

	return self.Client
}

func (self *FixerIoParser) getApiUrl(t time.Time, base Currency) string {
	scheme := "http"
	if self.UseTls {
		scheme += "s"
	}

	return scheme + fixerIoApiUrl + "/" + t.Format("2006-01-02") + "?base=" + base.Code
}

func (self *FixerIoParser) Parse(t time.Time, base Currency) (*Converter, error) {
	resp, err := self.getClient().Get(self.getApiUrl(t, base))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid status code")
	}

	res := struct {
		Base  string `json:"base"`
		Rates map[string]float64 `json:"rates"`
		Date  string `json:"date"`
		Error string `json:"error"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	if res.Error != "" {
		return nil, errors.New("invalid response: " + res.Error)
	}

	var ok bool

	c := &Converter{
		Rates: make(map[Currency]float64),
	}
	if c.Base, ok = Parse(res.Base); !ok {
		return nil, errors.New("undefined base currency")
	}
	for scur, rate := range res.Rates {
		cur, ok := Parse(scur)
		if !ok {
			continue
		}

		c.Rates[cur] = rate
	}
	date, err := time.Parse("2006-01-02", res.Date)
	if err != nil {
		return nil, err
	}
	c.Time = date

	return c, nil
}

func (self *FixerIoParser) MarginParse(t time.Time, base Currency, marginPercent float64) (*MarginConverter, error) {
	c, err := self.Parse(t, base)
	if err != nil {
		return nil, err
	}

	return &MarginConverter{c, marginPercent}, nil
}
