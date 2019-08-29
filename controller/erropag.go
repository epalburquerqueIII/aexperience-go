package controller

import (
	"fmt"
	"net/http"
)

// Pantalla Error 404
func Errorpag(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "404", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}
