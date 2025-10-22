package main

import (
	"fmt"
	"log"

	"coinmate_balance/coinmate"
	"coinmate_balance/exchange"
)

func main() {
	var client coinmate.Client

	err := client.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	balances, err := client.GetBalances()
	if err != nil {
		log.Fatal("Error fetching balances:", err)
	}

	btcBalance := balances["BTC"]

	priceCZK, err := client.GetBTCPrice("CZK")
	if err != nil {
		log.Fatal(err)
	}

	priceEUR, err := client.GetBTCPrice("EUR")
	if err != nil {
		log.Fatal(err)
	}

	eurToUSD, err := exchange.GetRate("EUR", "USD")
	if err != nil {
		log.Fatal(err)
	}

	priceUSD := priceEUR * eurToUSD
	valueCZK := btcBalance * priceCZK

	fmt.Printf("BTC Price:\t\t%.0f\t\tUSD\n", priceUSD)
	fmt.Printf("BTC Balance:\t\t%.8f\tBTC\n", btcBalance)
	fmt.Printf("BTC Balance in CZK:\t%.0f\t\tCZK\n", valueCZK)
}
