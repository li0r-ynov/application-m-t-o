package services

import (
	"fmt"
	"meteo/api"
	"meteo/models"
)

// SearchMeteo cherche d'abord la commune, puis récupère la météo
// correspondant à la position géographique de cette commune.
func SearchMeteo(query string) (models.Meteo, error) {
	// Demande à l'API des communes celles dont le nom correspond à la recherche.
	communes, err := api.SearchCommunesByName(query)
	if err != nil {
		// Si la recherche de commune échoue, on renvoie l'erreur au contrôleur.
		return models.Meteo{}, err
	}

	// Une recherche vide signifie qu'aucune commune n'a été trouvée.
	if len(communes) == 0 {
		return models.Meteo{}, models.ErrNotFound
	}

	// Le centre géographique doit contenir deux nombres : longitude puis latitude.
	// Sans ces coordonnées, l'API météo ne peut pas situer la commune.
	if len(communes[0].Centre.Coordinates) < 2 {
		return models.Meteo{}, models.ErrBadRequest
	}

	// L'API Géo renvoie les coordonnées dans l'ordre [longitude, latitude].
	longitude := communes[0].Centre.Coordinates[0]
	latitude := communes[0].Centre.Coordinates[1]

	// Utilise les coordonnées trouvées pour demander les prévisions météo.
	weather, err := api.SearchWeatherByCoordinates(latitude, longitude)

	if err != nil {
		// Transmet l'erreur si l'appel à l'API météo échoue.
		return models.Meteo{}, err
	}
	weather.CommuneName = communes[0].Name

	// Renvoie les prévisions météo obtenues.
	return weather, nil
}

func SearchMeteoByCode(code string) (models.Meteo, error) {
	communes, err := api.SearchCommuneByCode(code)
	if err != nil {
		return models.Meteo{}, err
	}

	if len(communes) == 0 {
		return models.Meteo{}, models.ErrNotFound
	}

	if len(communes[0].Centre.Coordinates) < 2 {
		return models.Meteo{}, models.ErrBadRequest
	}

	longitude := communes[0].Centre.Coordinates[0]
	latitude := communes[0].Centre.Coordinates[1]

	weather, err := api.SearchWeatherByCoordinates(latitude, longitude)
	if err != nil {
		return models.Meteo{}, err
	}

	weather.CommuneName = communes[0].Name
	weather.MapURL = GenerateMapURL(latitude, longitude)
	CalculateTemperatureBars(&weather)
	return weather, nil
}

func GenerateMapURL(latitude, longitude float64) string {
	d := 0.05

	south := latitude - d
	north := latitude + d
	west := longitude - d
	east := longitude + d

	return fmt.Sprintf(
		"https://www.openstreetmap.org/export/embed.html?bbox=%f,%f,%f,%f&layer=mapnik&marker=%f,%f",
		west, south, east, north, latitude, longitude,
	)
}

func CalculateTemperatureBars(weather *models.Meteo) {
	minWeek := weather.Daily.TemperatureMin[0]
	maxWeek := weather.Daily.TemperatureMax[0]

	for i := range weather.Daily.TemperatureMin {
		if weather.Daily.TemperatureMin[i] < minWeek {
			minWeek = weather.Daily.TemperatureMin[i]
		}

		if weather.Daily.TemperatureMax[i] > maxWeek {
			maxWeek = weather.Daily.TemperatureMax[i]
		}
	}

	weather.TemperatureMinWeek = minWeek
	weather.TemperatureMaxWeek = maxWeek

	for i := range weather.Daily.TemperatureMin {
		left := (weather.Daily.TemperatureMin[i] - minWeek) / (maxWeek - minWeek) * 100
		width := (weather.Daily.TemperatureMax[i] - weather.Daily.TemperatureMin[i]) / (maxWeek - minWeek) * 100

		weather.TemperatureBars = append(weather.TemperatureBars, models.TemperatureBar{
			Left:  left,
			Width: width,
		})
	}
}
