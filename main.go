package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string   `json:"name"`
	Main mainData `json:"main"`
}

type mainData struct {
	Kelvin     float64 `json:"temp"`
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return apiConfigData{}, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go!  \n"))
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?APPID=%s&q=%s", apiConfig.OpenWeatherMapApiKey, city)
	fmt.Println("Request URL:", url) // Debug print to verify the URL
	resp, err := http.Get(url)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weatherData{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return weatherData{}, fmt.Errorf("non-200 response: %d %s", resp.StatusCode, string(body))
	}

	fmt.Println("Response body:", string(body)) // Log the response body

	var d weatherData
	if err = json.Unmarshal(body, &d); err != nil {
		return weatherData{}, fmt.Errorf("error unmarshalling JSON: %v, response: %s", err, string(body))
	}

	// Calculate Celsius and Fahrenheit
	d.Main.Celsius = d.Main.Kelvin - 273.15
	d.Main.Fahrenheit = (d.Main.Kelvin-273.15)*9/5 + 32

	return d, nil
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}
