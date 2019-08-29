package controller

import (
	"fmt"
	"net/http"
)

// Pantalla Olvido contraseña
func Iva(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "iva", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}
