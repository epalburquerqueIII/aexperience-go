package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./controller"
	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.ParseGlob("views/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "index", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

func main() {
	log.Println("Server started on: http://localhost:3000")
	http.Handle("/static/css/", http.StripPrefix("/static/css", http.FileServer(http.Dir("static/css"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js", http.FileServer(http.Dir("static/js"))))
	http.Handle("/static/js/jtable", http.StripPrefix("/static/js/jtable", http.FileServer(http.Dir("static/js/jtable"))))
	http.HandleFunc("/", index)
	// Apis usuario

	http.HandleFunc("/usuario/list", controller.UsuarioList)
	http.HandleFunc("/usuario/create", controller.UsuarioCreate)
	http.HandleFunc("/usuario/update", controller.UsuarioUpdate)
	http.HandleFunc("/usuario/delete", controller.UsuarioDelete)
	http.HandleFunc("/usuario/getoptionsRoles", controller.UsuariogetoptionsRoles)

	http.ListenAndServe(":3000", nil)
}
