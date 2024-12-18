package currency

import (
	"encoding/json"
	"io"
	"net/http"
)

type Currency struct {
	Code  string
	Name  string
	Rates map[string]float64
}

type MyCurrencyExchange struct {
	Currencies map[string]Currency
}

func (ce *MyCurrencyExchange) FetchAllCurrencies() error {
	resp, err := http.Get(
		"http://cdn.jsdelivr.net/npm/fawazahmed0/currency-api@latest/v1/currencies.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	cs, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	csMap := make(map[string]string)
	err = json.Unmarshal(cs, &csMap)
	if err != nil {
		return err
	}

	i := 0
	for code, name := range csMap {
		if i > 10 {
			break
		}
		c := Currency{
			Code:  code,
			Name:  name,
			Rates: make(map[string]float64),
		}
		ce.Currencies[code] = c
		i++
	}
}
