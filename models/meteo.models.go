package models

type Meteo struct {
	Latitude             float64    `json:"latitude"`
	Longitude            float64    `json:"longitude"`
	TimeZone             string     `json:"timezone"`
	TimeZoneAbbreviation string     `json:"timezone_abbreviation"`
	Elevation            float64    `json:"elevation"`
	DailyUnits           DailyUnits `json:"daily_units"`
	Daily                Daily      `json:"daily"`
	CommuneName          string     `json:"-"`
	MapURL               string
	TemperatureMinWeek   float64
	TemperatureMaxWeek   float64
	TemperatureBars      []TemperatureBar
}

type DailyUnits struct {
	Time           string `json:"time"`
	TemperatureMax string `json:"temperature_2m_max"`
	TemperatureMin string `json:"temperature_2m_min"`
	WeatherCode    string `json:"weather_code"`
}

type Daily struct {
	Time           []string  `json:"time"`
	TemperatureMax []float64 `json:"temperature_2m_max"`
	TemperatureMin []float64 `json:"temperature_2m_min"`
	WeatherCode    []int     `json:"weather_code"`
}

type TemperatureBar struct {
	Left  float64
	Width float64
}
