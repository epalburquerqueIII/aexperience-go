package controller

import (
	"fmt"
	"net/http"

	"../util"
)

// Pantalla Olvido contraseña
func Iva(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)
	error := tmpl.ExecuteTemplate(w, "iva", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}
