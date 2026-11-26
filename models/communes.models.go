package models

type Communes struct {
	Code       string `json:"code"`
	Name       string `json:"nom"`
	Population int    `json:"population"`
	Centre     Centre `json:"centre"`
}

type Centre struct {
	Coordinates []float64 `json:"coordinates"`
}
