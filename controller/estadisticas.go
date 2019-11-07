package controller

import (
	"fmt"
	"net/http"

	"../util"
)

// Pantalla de estadisticas
func Estadisticas(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)
	error := tmpl.ExecuteTemplate(w, "estadisticas", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}
