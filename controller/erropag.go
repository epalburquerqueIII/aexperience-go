package controller

import (
	"fmt"
	"net/http"

	"../util"
)

// Pantalla Error 404
func Errorpag(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)

	error := tmpl.ExecuteTemplate(w, "404", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}
