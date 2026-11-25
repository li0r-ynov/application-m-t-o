package models

type Meteo struct {
	Latitude  int           `json:"latitude"`
	Longitude int           `json:"code"`
	TimeZone  string        `json:"timezone"`
	Elevation int           `json:"elevation"`
	UnitHour  []HourlyUnits `json:"hourly_units"`
	Hourly    []Hourly      `json:"hourly"`
}

type HourlyUnits struct {
	Time        string `json:"time"`
	Temperature string `json:"temperature_2m"`
}

type Hourly struct {
	Time        []string `json:"time"`
	Temperature []int    `json:"temperature_2m"`
}
