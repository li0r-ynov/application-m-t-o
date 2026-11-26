package helpers

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
)

func Render(w http.ResponseWriter, status int, templateName string, data any) error {
	path := filepath.Join("templates", templateName+".html")
	tmpl, err := template.ParseFiles(path)
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
