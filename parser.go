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

	var c Converter
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
