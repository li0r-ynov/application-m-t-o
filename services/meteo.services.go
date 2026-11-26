package services

import (
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
	weather, err := api.SearchWeatherByCoordinates(longitude, latitude)

	if err != nil {
		// Transmet l'erreur si l'appel à l'API météo échoue.
		return models.Meteo{}, err
	}
	weather.CommuneName = communes[0].Name

	// Renvoie les prévisions météo obtenues.
	return weather, nil
}
