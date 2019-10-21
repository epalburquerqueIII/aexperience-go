package controller

import (
	"fmt"
	"net/http"
)

// Pantalla Pagina vacia
func EmailNewsletter(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "emailnewsletter", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}
