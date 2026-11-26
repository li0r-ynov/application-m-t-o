package api

import (
	"fmt"
	"meteo/models"
)

var (
	baseMeteoURL = "https://api.open-meteo.com"
)

func SearchWeatherByCoordinates(latitude float64, longitude float64) (models.Meteo, error) {
	url := fmt.Sprintf("%s/v1/forecast?latitude=%f&longitude=%f&daily=temperature_2m_max,temperature_2m_min,weather_code&timezone=auto&forecast_days=7", baseMeteoURL, latitude, longitude)

	var resultWeather models.Meteo
	err := get(url, &resultWeather)
	return resultWeather, err

}
