package controllers

import (
	"errors"
	"meteo/helpers"
	"meteo/services"
	"meteo/models"
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
	helpers.Render(w, http.StatusOK, "commune")
}

func RenderAPIError(w http.ResponseWriter, err error){
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
