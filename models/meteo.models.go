package models

type Meteo struct {
	Latitude             int          `json:"latitude"`
	Longitude            int          `json:"code"`
	TimeZone             string       `json:"timezone"`
	TimeZoneAbbreviation string       `json:"timezone_abbreviation"`
	Elevation            int          `json:"elevation"`
	DailyUnits           []DailyUnits `json:"daily_units"`
	Daily                []Daily      `json:"Daily"`
}

type DailyUnits struct {
	Time           string `json:"time"`
	TemperatureMax string `json:"temperature_2m_max"`
	TemperatureMin string `json:"temperature_2m_min"`
	WeatherCode    string `json:"wmo_code"`
}

type Daily struct {
	Time           []string `json:"time"`
	TemperatureMax []int    `json:"temperature_2m_max"`
	TemperatureMin []int    `json:"temperature_2m_min"`
	WeatherCode    []int    `json:"wmo_code"`
}
