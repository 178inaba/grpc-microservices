package template

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func Render(w http.ResponseWriter, name string, content interface{}) error {
	t := template.Must(template.ParseFiles(
		"template/layout.tmpl", "template/header.tmpl", filepath.Join("template", name)))
	if err := t.ExecuteTemplate(w, "layout", content); err != nil {
		return err
	}

	return nil
}
