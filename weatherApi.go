package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	Timezone string `json:"timezone"`
	Current  struct {
		Temperature float64 `json:"temperature_2m"`
		Humidity    int     `json:"relative_humidity_2m"`
		Windspeed   float64 `json:"wind_speed_10m"`
		Weathercode int     `json:"weathercode"`
	} `json:"current"`
}

func getWeather(lat float64, lon float64) string {

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m,relative_humidity_2m,weather_code,wind_speed_10m", lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		return "error: weather request failed"
	}
	defer resp.Body.Close()

	var w WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		return "error: could not parse weather"
	}

	return fmt.Sprintf("Temp:%.1f deg C, Humidity: %d%%, Wind: %.1fkm/h",
		w.Current.Temperature,
		w.Current.Humidity,
		w.Current.Windspeed)
}
