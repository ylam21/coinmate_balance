package coinmate

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TickerResponse struct {
	Error bool `json:"error"`
	Data  struct {
		Last float64 `json:"last"` // "Last" executed trade price
	} `json:"data"`
}

// CoinMate only supports certain trading pairs.
// They currently (today is 12.10.2025) only list BTC_EUR and BTC_CZK pairs (and a few altcoins).
func (c *Client) GetBTCPrice(currency string) (float64, error) {
	if currency != "EUR" && currency != "CZK" {
		return 0, fmt.Errorf("Unsupported currency: %s", currency)
	}
	url := fmt.Sprintf("https://coinmate.io/api/ticker?currencyPair=BTC_%s", currency)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var ticker TickerResponse
	if err := json.NewDecoder(resp.Body).Decode(&ticker); err != nil {
		return 0, err
	}
	if ticker.Error {
		return 0, fmt.Errorf("Error fetching BTC_%s price", currency)
	}
	return ticker.Data.Last, nil
}
