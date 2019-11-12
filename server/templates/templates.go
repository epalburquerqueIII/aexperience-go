package templates

import (
	"html/template"
	"log"
	"net/http"

	"../../model"
)

type LoginPage struct {
	BAlertUser bool
	AlertMsg   string
}

type RegisterPage struct {
	BAlertUser bool
	AlertMsg   string
}

// RestrictedPage Estructura para página restringidas
type RestrictedPage struct {
	AuthWeb model.AuthWeb
	Menus   []model.Tmenuconfig
}

var templates = template.Must(template.ParseGlob("./views/*.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		log.Printf("Temlate error here: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
