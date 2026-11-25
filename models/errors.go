package models

import "errors"

var (
	ErrNotFound    = errors.New("Ressource introuvable")
	ErrBadRequest  = errors.New("Demande invalide")
	ErrUnavailable = errors.New("API indisponible")
	ErrTimeout     = errors.New("Délai d'attente dépassé")
)