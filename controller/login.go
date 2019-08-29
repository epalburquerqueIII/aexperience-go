package controller

import (
	"fmt"
	"net/http"
)

// Pantalla Login
func Login(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "login", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}
