package models

type Communes struct {
	Name   string   `json:"nom"`
	Code   string   `json:"code"`
	Centre []Centre `json:"centre"`
}

type Centre struct {
	Type        string `json:"type"`
	Coordinates []int  `json:"coordinates"`
}
