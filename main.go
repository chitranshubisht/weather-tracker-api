package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name    string   `json:"name"`
	Main    mainData `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
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
	w.Write([]byte("hello from go!\n"))
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
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	// Enable CORS
	c := cors.Default()
	handler := c.Handler(mux)

	mux.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		if city == "" {
			http.Error(w, "City not specified", http.StatusBadRequest)
			return
		}

		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
