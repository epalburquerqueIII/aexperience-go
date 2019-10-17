package controller

import (
	"fmt"
	"net/http"

	"../util"
)

// Pantalla Pagina vacia
func Paginavacia(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)
	error := tmpl.ExecuteTemplate(w, "paginavacia", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}
