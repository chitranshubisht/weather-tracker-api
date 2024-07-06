# Weather Tracker API 🌤️

Welcome to the Weather Tracker API project! This API fetches weather data from OpenWeatherMap and provides temperature information in Kelvin, Celsius, and Fahrenheit for a given city.

## Prerequisites 📋

Before running the project, ensure you have the following installed:

- Go (Golang)
- Git

## Getting Started 🚀

Follow these steps to set up and run the Weather Tracker API:

### 1. Clone the Repository 🌀

```bash
git clone git@github.com:chitranshubisht/weather-tracker-api.git
cd weather-tracker-api
```

## Setting Up the Weather Tracker API

### Set Up API Configuration 🛠️

Create a `.apiConfig` file in the root directory with your OpenWeatherMap API key:

```json
{
    "OpenWeatherMapApiKey": "your_api_key_here"
}
```

### 2. Build and Run the API 🏃‍♂️

Run the following command to build and start the API locally at [http://localhost:8080](http://localhost:8080):

```bash
go run main.go
```

### 3. Access Weather Data 🌍

To access weather data for a specific city, make a GET request to:

http://localhost:8080/weather/{city_name}


Replace `{city_name}` with the name of the city you want weather data for. For example:



