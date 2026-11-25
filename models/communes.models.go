package models

type Meteo struct {
	Name        string        `json:"nom"`
	Code        string        `json:"code"`
	Score       float64       `json:"_score"`
	Departement []Departement `json:"departement"`
}

type Departement struct {
	Code string `json:"code"`
	Nom  string `json:"nom"`
}
