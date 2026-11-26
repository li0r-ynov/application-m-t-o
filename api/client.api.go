package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"meteo/models"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}



func get(endpoint string, dest any) error {
	// 1. Préparer la requête : méthode, URL et corps (nil pour un GET).
	request, err := http.NewRequest(http.MethodGet, baseMeteoURL+baseCommunesURL+endpoint, nil)
	if err != nil {
		return fmt.Errorf("création de la requête : %w", err)
	}
	request.Header.Set("Accept", "application/json")
	// 2. Envoyer la requête avec le client.
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("%w : %v", models.ErrUnavailable, err)
	}
	defer response.Body.Close()

	// 3. Vérifier le code de statut avant de lire le corps.
	switch {
	case response.StatusCode == http.StatusOK:
		// 200 : succès, on peut décoder le corps.
	case response.StatusCode == http.StatusNotFound:
		return models.ErrNotFound
	case response.StatusCode == http.StatusBadRequest:
		return models.ErrBadRequest
	case response.StatusCode >= 500:
		return fmt.Errorf("%w : %s", models.ErrUnavailable, response.Status)
	default:
		return fmt.Errorf("réponse inattendue : %s", response.Status)
	}

	// 4. Décoder le flux JSON directement dans la variable de destination.
	if err := json.NewDecoder(response.Body).Decode(dest); err != nil {
		return fmt.Errorf("décodage de la réponse : %w", err)
	}
	return nil
}
