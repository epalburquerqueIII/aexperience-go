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
	http.Handle("/static/vendor/fontawesome-free/css", http.StripPrefix("/static/vendor/fontawesome-free/css", http.FileServer(http.Dir("static/vendor/fontawesome-free/css"))))
	http.Handle("/static/vendor/jquery-easing", http.StripPrefix("/static/vendor/jquery-easing", http.FileServer(http.Dir("/static/vendor/jquery-easing"))))
	http.Handle("/static/vendor/chart.js", http.StripPrefix("static/vendor/chart.js", http.FileServer(http.Dir("static/vendor/chart.js"))))
	http.Handle("/static/vendor/jquery", http.StripPrefix("/static/vendor/jquery", http.FileServer(http.Dir("/static/vendor/jquery"))))
	http.Handle("/static/vendor/bootstrap/js", http.StripPrefix("/static/vendor/bootstrap/js", http.FileServer(http.Dir("/static/vendor/bootstrap/js"))))
	http.Handle("/static/js/demo", http.StripPrefix("/static/js/demo/", http.FileServer(http.Dir("/static/js/demo"))))
	http.Handle("/static/img", http.StripPrefix("/static/img", http.FileServer(http.Dir("/static/img"))))
	http.HandleFunc("/", index)
	// Apis usuario

	http.HandleFunc("/usuario", controller.Usuario)
	http.HandleFunc("/usuario/list", controller.UsuarioList)
	http.HandleFunc("/usuario/create", controller.UsuarioCreate)
	http.HandleFunc("/usuario/update", controller.UsuarioUpdate)
	http.HandleFunc("/usuario/delete", controller.UsuarioDelete)
	http.HandleFunc("/usuario/getoptionsRoles", controller.UsuariogetoptionsRoles)

	http.ListenAndServe(":3000", nil)
}
