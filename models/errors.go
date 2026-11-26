package models

import "errors"

var (
	ErrNotFound    = errors.New("Commune introuvable")
	ErrBadRequest  = errors.New("Paramètres de recherche invalides")
	ErrUnavailable = errors.New("Service météo temporairement indisponible")
	ErrTimeout     = errors.New("Délai d'attente dépassé")
)
