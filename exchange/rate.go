package exchange

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRate(from, to string) (float64, error) {
	url := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", from)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Rates map[string]float64 `json:"rates"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	rate, ok := result.Rates[to]
	if !ok {
		return 0, fmt.Errorf("currency %s not found", to)
	}
	return rate, nil
}
