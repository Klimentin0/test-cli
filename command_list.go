package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func listPrompt() error {
	url := "https://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,XRP,SOL&tsyms=USD,EUR,RUB"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error receiving status code %d", resp.StatusCode)
	}

	var cryptoData map[string]map[string]float64

	if err := json.NewDecoder(resp.Body).Decode(&cryptoData); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v", err)
	}

	for crypto, values := range cryptoData {
		fmt.Printf("%s Prices:\n", crypto)
		for currency, price := range values {
			fmt.Printf(" %s: %.2f\n", currency, price)
		}
	}

	return nil
}
