package controller

import (
	"fmt"
	"net/http"
)

// Pantalla Registro
func Registro(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "registro", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}
