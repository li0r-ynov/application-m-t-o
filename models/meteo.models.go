package models

type Meteo struct {
	Latitude             int           `json:"latitude"`
	Longitude            int           `json:"code"`
	TimeZone             string        `json:"timezone"`
	TimeZoneAbbreviation string        `json:"timezone_abbreviation"`
	Elevation            int           `json:"elevation"`
	HourlyUnits          []HourlyUnits `json:"hourly_units"`
	Hourly               []Hourly      `json:"hourly"`
}

type HourlyUnits struct {
	Time        string `json:"time"`
	Temperature string `json:"temperature_2m_min"`
}

type Hourly struct {
	Time        []string `json:"time"`
	Temperature []int    `json:"temperature_2m_min"`
}
