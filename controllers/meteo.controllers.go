package controllers

import (
	"errors"
	"meteo/api"
	"meteo/helpers"
	"meteo/models"
	"meteo/services"
	"net/http"
	"strings"
)

func MeteoController(w http.ResponseWriter, r *http.Request) {
	search := strings.TrimSpace(r.URL.Query().Get("commune"))
	if search == "" {
		http.Error(w, "Veuillez saisir une commune", http.StatusBadRequest)
		return
	}
	meteo, err := services.SearchMeteo(search)
	if err != nil {
		RenderAPIError(w, err)
		return
	}
	if err := helpers.Render(w, http.StatusOK, "resultat", meteo); err != nil {
		http.Error(w, "Erreur lors de l'affichage de la météo", http.StatusInternalServerError)
		return
	}
}

func VillesController(w http.ResponseWriter, r *http.Request) {
	search := strings.TrimSpace(r.URL.Query().Get("commune"))

	if search == "" {
		http.Error(w, "Veuillez saisir une ville", http.StatusBadRequest)
		return
	}

	communes, err := api.SearchCommunesByName(search)
	if err != nil {
		http.Error(w, "Erreur lors de la recherche", http.StatusInternalServerError)
		return
	}

	if err := helpers.Render(w, http.StatusOK, "list-villes", communes); err != nil {
		http.Error(w, "Erreur lors de l'affichage", http.StatusInternalServerError)
	}
}

func RenderAPIError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		http.Error(w, "Commune introuvable", http.StatusNotFound)

	case errors.Is(err, models.ErrBadRequest):
		http.Error(w, "Requête invalide", http.StatusBadRequest)

	case errors.Is(err, models.ErrUnavailable):
		http.Error(w, "Service indisponible", http.StatusServiceUnavailable)

	case errors.Is(err, models.ErrTimeout):
		http.Error(w, "Délai d'attente dépassé", http.StatusGatewayTimeout)

	default:
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}
