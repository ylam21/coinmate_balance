package main

import (
	"coinmate_balance/coinmate"
	"coinmate_balance/exchange"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

func printTable(priceUSD, btcBalance, valueCZK float64) {
	yellow := color.New(color.FgYellow).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()
	faint := color.New(color.Faint).SprintFunc()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmtPrice := yellow(fmt.Sprintf("%12.2f", priceUSD))
	fmtBal := bold(fmt.Sprintf("%12.8f", btcBalance))
	fmtValCZK := bold(fmt.Sprintf("%12.2f", valueCZK))

	fmt.Fprintf(w, "%s\t%s \t%s\n", faint("BTC Price:"), yellow("$"), fmtPrice)
	fmt.Fprintln(w, "\t\t")
	fmt.Fprintf(w, "%s\t\t\n", bold("Portfolio:"))
	fmt.Fprintf(w, "%s\t\t\n", bold("----------"))
	fmt.Fprintf(w, "%s\t%s \t%s\n", faint("Balance (BTC):"), bold("₿"), fmtBal)
	fmt.Fprintf(w, "%s\t%s \t%s\n", faint("Value (CZK):"), bold("Kč"), fmtValCZK)

	w.Flush()
}

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

	btcBalance := balances["BTC"]
	priceUSD := priceEUR * eurToUSD
	valueCZK := btcBalance * priceCZK
	printTable(priceUSD, btcBalance, valueCZK)
}
