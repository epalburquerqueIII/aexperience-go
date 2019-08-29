package controller

import (
	"fmt"
	"net/http"
)

// Pantalla de estadisticas
func Estadisticas(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "estadisticas", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}
