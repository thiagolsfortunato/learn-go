package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// Set HTML templates into templates var
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecTemplates(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
