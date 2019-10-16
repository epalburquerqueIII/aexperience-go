package controller

import (
	"fmt"
	"net/http"
)

// Recuperarcontrasena pantalla
func Recuperarcontrasena(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "recuperarcontrasena", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}
