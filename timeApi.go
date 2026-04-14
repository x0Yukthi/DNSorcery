package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type GeoResponse struct {
	Results []struct {
		Timezone string `json:"timezone"`
	} `json:"results"`
}

var locationApi = "https://geocoding-api.open-meteo.com/v1/search?name="

func findLocation(location string) string {
	location = strings.ReplaceAll(location, " ", "%20")

	url := locationApi + location + "&count=1"

	resp, err := http.Get(url)
	if err != nil {
		return "error: request failed"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "error: could not read response"
	}

	var geo GeoResponse
	json.Unmarshal(body, &geo)
	if len(geo.Results) == 0 {
		return "error: location not found"
	}
	timezone := geo.Results[0].Timezone

	return timezone

}

func getTime(location string) string {
	timezone := findLocation(location)
	if timezone != "" {
		loc, err := time.LoadLocation(timezone)
		if err != nil {
			return "error: bad timezone"
		}
		return time.Now().In(loc).Format("15:04:05 MST, Mon 02 Jan 2006")
	}
	return "error: location not supported"

}
