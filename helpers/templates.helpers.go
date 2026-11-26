package helpers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

var funcMap = template.FuncMap{
	"frDate":       frDate,
	"formatTemp":   formatTemp,
	"weatherLabel": weatherLabel,
	"weatherIcon":  weatherIcon,
}

func frDate(date string) string {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return date
	}

	jours := []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."}
	mois := []string{"jan.", "fév.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."}

	return fmt.Sprintf("%s %d %s", jours[parsedDate.Weekday()], parsedDate.Day(), mois[parsedDate.Month()-1])
}

func formatTemp(temp float64) string {
	return fmt.Sprintf("%.1f °C", temp)
}

func weatherLabel(code int) string {
	switch {
	case code == 0:
		return "Ciel dégagé"
	case code == 1:
		return "Principalement dégagé"
	case code == 2:
		return "Partiellement nuageux"
	case code == 3:
		return "Couvert"
	case code >= 45 && code <= 48:
		return "Brouillard"
	case code >= 51 && code <= 57:
		return "Bruine"
	case code >= 61 && code <= 67:
		return "Pluie"
	case code >= 71 && code <= 77:
		return "Neige"
	case code >= 80 && code <= 82:
		return "Averses"
	case code >= 85 && code <= 86:
		return "Averses de neige"
	case code >= 95 && code <= 99:
		return "Orage"
	default:
		return "Conditions inconnues"
	}
}

func weatherIcon(code int) string {
	switch {
	case code == 0:
		return "☀️"
	case code == 1 || code == 2:
		return "🌤️"
	case code == 3:
		return "☁️"
	case code >= 45 && code <= 48:
		return "🌫️"
	case code >= 51 && code <= 67:
		return "🌧️"
	case code >= 71 && code <= 77:
		return "❄️"
	case code >= 80 && code <= 82:
		return "🌦️"
	case code >= 85 && code <= 86:
		return "🌨️"
	case code >= 95 && code <= 99:
		return "⛈️"
	default:
		return "❔"
	}
}

func Render(w http.ResponseWriter, status int, templateName string, data any) error {
	path := filepath.Join("templates", templateName+".html")
	tmpl, err := template.New(templateName + ".html").Funcs(funcMap).ParseFiles(path)
	if err != nil {
		return err
	}
	var buffer bytes.Buffer

	if err := tmpl.Execute(&buffer, data); err != nil {
		return err
	}

	w.WriteHeader(status)
	_, err = buffer.WriteTo(w)
	return err

}
