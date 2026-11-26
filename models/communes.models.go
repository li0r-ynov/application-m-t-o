package models

type Communes struct {
	Code   string   `json:"code"`
	Name   string   `json:"nom"`
	CodesPostaux
	Centre []Centre `json:"centre"`
}

type Centre struct {
	Type        string `json:"type"`
	Coordinates []int  `json:"coordinates"`
}
