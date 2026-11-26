package api

import (
	"fmt"
	"meteo/models"
)

var (
	baseCommunesURL = "https://geo.api.gouv.fr"
)

func SearchCommunesByName(query string) ([]models.Communes, error) {
	url := fmt.Sprintf("%s/communes?nom=%s&fields=code,nom,population,centre", baseCommunesURL, query)

	var resultCommunes []models.Communes
	err := get(url, &resultCommunes)
	return resultCommunes, err
}

func SearchCommuneByCode(code string) ([]models.Communes, error) {
	url := fmt.Sprintf("%s/communes?code=%s&fields=code,nom,population,centre", baseCommunesURL, code)

	var resultCommunes []models.Communes
	err := get(url, &resultCommunes)
	return resultCommunes, err
}
