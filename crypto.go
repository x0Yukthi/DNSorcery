package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CryptoResponse map[string]struct {
	USD       float64 `json:"usd"`
	Change24h float64 `json:"usd_24h_change"`
}

var cryptoAliases = map[string]string{
	"btc":  "bitcoin",
	"eth":  "ethereum",
	"sol":  "solana",
	"doge": "dogecoin",
}

func getCrypto(coin string) string {
	if alias, ok := cryptoAliases[coin]; ok {
		coin = alias
	}

	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + coin + "&vs_currencies=usd&include_24hr_change=true"
	resp, err := http.Get(url)
	if err != nil {
		return "error: query request failed"
	}
	defer resp.Body.Close()

	var crypto CryptoResponse
	if err := json.NewDecoder(resp.Body).Decode(&crypto); err != nil {
		return "error: could not parse query"
	}
	data, ok := crypto[coin]
	if !ok {
		return "error: coin not found — try bitcoin, ethereum, solana, dogecoin"
	}

	return fmt.Sprintf("%s: $%.2f | 24h: %+.2f%%",
		coin,
		data.USD,
		data.Change24h,
	)
}
