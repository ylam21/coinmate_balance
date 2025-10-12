package coinmate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type BalancesResponse struct {
	Error        bool   `json:"error"`
	ErrorMessage string `json:"errorMessage"`
	Data         map[string]struct {
		Available float64 `json:"available"`
		Reserved  float64 `json:"reserved"`
	} `json:"data"`
}

func (c *Client) GetBalances() (map[string]float64, error) {
	nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)
	signature := c.computeSignature(nonce)

	data := url.Values{}
	data.Set("clientId", c.ClientID)
	data.Set("publicKey", c.PublicKey)
	data.Set("nonce", nonce)
	data.Set("signature", signature)

	resp, err := http.PostForm("https://coinmate.io/api/balances", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result BalancesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.Error {
		return nil, fmt.Errorf(result.ErrorMessage)
	}

	balances := make(map[string]float64)
	for symbol, v := range result.Data {
		balances[symbol] = v.Available
	}
	return balances, nil
}
