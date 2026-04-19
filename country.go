package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type CountryResponse struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Capital    []string `json:"capital"`
	Region     string   `json:"region"`
	Subregion  string   `json:"subregion"`
	Population int      `json:"population"`
	Currencies map[string]struct {
		Name string `json:"name"`
	} `json:"currencies"`
}

func getCountry(name string) string {
	url := "https://restcountries.com/v3.1/name/" + name
	resp, err := http.Get(url)
	if err != nil {
		return "error: Country request failed"
	}
	defer resp.Body.Close()

	var countries []CountryResponse
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return "error: could not parse response"
	}

	if len(countries) == 0 {
		return "error: country not found"
	}
	c := countries[0]
	currencyName := "unknown"
	for _, v := range c.Currencies {
		currencyName = v.Name
		break
	}
	return fmt.Sprintf("%s | Capital: %s | Region: %s | Subregion: %s | Pop: %.1fM  | Currency: %s",
		c.Name.Common,
		c.Capital[0],
		c.Region,
		c.Subregion,
		float64(c.Population)/1_000_000,
		currencyName,
	)

}

type ConversionResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func getConversion(query string) string {
	cur := strings.Split(query, " ")
	url := "https://api.frankfurter.app/latest?from=" + cur[1] + "&to=" + cur[2]
	resp, err := http.Get(url)
	amount, err := strconv.ParseFloat(cur[0], 64)
	if err != nil {
		return "error: Currency conversion request failed"
	}
	if len(cur) != 3 {
		return "error: use convert.[amount].[from].[to]"
	}
	defer resp.Body.Close()

	var c ConversionResponse
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return "error: could not parse currency"
	}

	rate := c.Rates[cur[2]]
	result := amount * rate

	return fmt.Sprintf("%.2f %s = %.2f %s", amount, cur[1], result, cur[2])
}
