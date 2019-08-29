package controller

import (
	"fmt"
	"net/http"
)

// Pantalla Olvido contraseña
func Paginavacia(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "paginavacia", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}
