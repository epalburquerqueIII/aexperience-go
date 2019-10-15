package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./controller"
	"./util"
	_ "github.com/go-sql-driver/mysql"
)

const usertype int = 0

var tmpl = template.Must(template.ParseGlob("views/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)

	error := tmpl.ExecuteTemplate(w, "index", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

func main() {
	log.Println("Server started on: http://localhost:3000")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", index)

	// Apis usuario
	http.HandleFunc("/usuarios", controller.Usuario)
	http.HandleFunc("/usuarios/list", controller.UsuarioList)
	http.HandleFunc("/usuarios/create", controller.UsuarioCreate)
	http.HandleFunc("/usuarios/update", controller.UsuarioUpdate)
	http.HandleFunc("/usuarios/delete", controller.UsuarioDelete)
	http.HandleFunc("/usuarios/getoptions", controller.Usuariogetoptions)

	http.HandleFunc("/estadisticas", controller.Estadisticas)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/registro", controller.Registro)
	http.HandleFunc("/404", controller.Errorpag)
	http.HandleFunc("/olvido-contrasena", controller.Olvidocontrasena)
	http.HandleFunc("/paginavacia", controller.Paginavacia)
	http.HandleFunc("/iva", controller.Iva)

	// Apis consumoBonos
	http.HandleFunc("/consumobonos", controller.ConsumoBonos)
	http.HandleFunc("/consumobonos/list", controller.ConsumoBonosList)
	http.HandleFunc("/consumobonos/create", controller.ConsumoBonosCreate)
	http.HandleFunc("/consumobonos/update", controller.ConsumoBonosUpdate)

	// Apis bono
	http.HandleFunc("/bonos", controller.Bonos)
	http.HandleFunc("/bonos/list", controller.BonoList)
	http.HandleFunc("/bonos/update", controller.BonoUpdate)

	// Apis autorizados
	http.HandleFunc("/autorizados", controller.Autorizado)
	http.HandleFunc("/autorizados/list", controller.AutorizadoList)
	http.HandleFunc("/autorizados/create", controller.AutorizadoCreate)
	http.HandleFunc("/autorizados/update", controller.AutorizadoUpdate)
	http.HandleFunc("/autorizados/delete", controller.AutorizadoDelete)
	http.HandleFunc("/autorizados/getoptions", controller.Autorizadogetoptions)

	http.ListenAndServe(":3000", nil)
}
