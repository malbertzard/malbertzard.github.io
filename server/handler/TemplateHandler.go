package handlers

import (
	"net/http"
	"path/filepath"
	"text/template"
)

type TemplateHandler struct {
	template *template.Template
}

func NewTemplateHandler(templateFile string) *TemplateHandler {
	templatePath := filepath.Join("templates", templateFile)
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}
	return &TemplateHandler{tmpl}
}

func (th *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    err := th.template.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
