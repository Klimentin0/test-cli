package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
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

	var sortCrypto []string
	for item := range cryptoData {
		sortCrypto = append(sortCrypto, item)
	}

	sort.Strings(sortCrypto)

	for _, crypto := range sortCrypto {
		fmt.Printf("%s Prices:\n", crypto)
		values := cryptoData[crypto]
		for currency, price := range values {
			fmt.Printf(" %s: %.2f\n", currency, price)
		}
		fmt.Println()
	}

	return nil
}
