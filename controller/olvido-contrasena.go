package controller

import (
	"fmt"
	"net/http"
)

// Pantalla Olvido contraseña
func Olvidocontrasena(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "olvido-contrasena", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}
