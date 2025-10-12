package main

import (
	"fmt"
	"log"

	"coinmate_balance/coinmate"
	"coinmate_balance/exchange"
)

func main() {
	client, err := coinmate.NewClientFromEnv()
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

	fmt.Printf("BTC Price: %.0f USD\n", priceUSD)
	fmt.Printf("BTC Balance: %.8f BTC\n", btcBalance)
	fmt.Printf("BTC Balance in CZK: %.0f CZK\n", valueCZK)
}
