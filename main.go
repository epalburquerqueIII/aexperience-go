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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", index)
	// Apis usuario

	http.HandleFunc("/usuario", controller.Usuario)
	http.HandleFunc("/usuario/list", controller.UsuarioList)
	http.HandleFunc("/usuario/create", controller.UsuarioCreate)
	http.HandleFunc("/usuario/update", controller.UsuarioUpdate)
	http.HandleFunc("/usuario/delete", controller.UsuarioDelete)
	http.HandleFunc("/usuario/getoptionsRoles", controller.UsuariogetoptionsRoles)
	http.HandleFunc("/estadisticas", controller.Estadisticas)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/registro", controller.Registro)
	http.HandleFunc("/404", controller.Errorpag)
<<<<<<< HEAD
	http.HandleFunc("/olvido-contrasena", controller.Olvidocontrasena)
=======
<<<<<<< HEAD
	http.HandleFunc("/registro", controller.Registro)
=======
	http.HandleFunc("/paginavacia", controller.Paginavacia)
>>>>>>> 165e681f905b7ad4496ccdedabc65131cee88610
>>>>>>> 211d78108e2a6c373f78b391749afc238de4c971

	http.ListenAndServe(":3000", nil)
}
