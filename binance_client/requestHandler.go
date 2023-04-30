package binance_client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Ticker struct {
	Symbol string
	Price  float64 `json:",string"`
}

// A function to get the ticker for a request symbol
// @name name of the ticker to be retrieved
func GetTicker(name string) *Ticker {
	binanceURL := "https://www.binance.com/api/v3/ticker/price?symbol="
	resp, err := http.Get(binanceURL + name + "USDT")
	if err != nil {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	ticker := new(Ticker)
	json.Unmarshal(body, ticker)
	ticker.Symbol = name
	return ticker
}
