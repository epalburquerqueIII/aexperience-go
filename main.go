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

	error := tmpl.ExecuteTemplate(w, "index", menu)
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
	http.HandleFunc("/usuario/getoptions", controller.Usuariogetoptions)

	http.HandleFunc("/estadisticas", controller.Estadisticas)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/registro", controller.Registro)
	http.HandleFunc("/404", controller.Errorpag)
	http.HandleFunc("/olvido-contrasena", controller.Olvidocontrasena)
	http.HandleFunc("/paginavacia", controller.Paginavacia)
	http.HandleFunc("/iva", controller.Iva)

	// Apis consumoBonos
	http.HandleFunc("/consumoBonos", controller.ConsumoBonos)
	http.HandleFunc("/consumoBonos/list", controller.ConsumoBonosList)
	http.HandleFunc("/consumoBonos/create", controller.ConsumoBonosCreate)
	http.HandleFunc("/consumoBonos/update", controller.ConsumoBonosUpdate)

	// Apis bono
	http.HandleFunc("/bono", controller.Bonos)
	http.HandleFunc("/bono/list", controller.BonoList)
	http.HandleFunc("/bono/update", controller.BonoUpdate)

	// Apis autorizados
	http.HandleFunc("/autorizado", controller.Autorizado)
	http.HandleFunc("/autorizado/list", controller.AutorizadoList)
	http.HandleFunc("/autorizado/create", controller.AutorizadoCreate)
	http.HandleFunc("/autorizado/update", controller.AutorizadoUpdate)
	http.HandleFunc("/autorizado/delete", controller.AutorizadoDelete)
	http.HandleFunc("/autorizado/getoptions", controller.Autorizadogetoptions)

	http.ListenAndServe(":3000", nil)
}
