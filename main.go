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
	http.HandleFunc("/usuarios", controller.Usuarios)
	http.HandleFunc("/usuarios/list", controller.UsuariosList)
	http.HandleFunc("/usuarios/create", controller.UsuariosCreate)
	http.HandleFunc("/usuarios/update", controller.UsuariosUpdate)
	http.HandleFunc("/usuarios/delete", controller.UsuariosDelete)
	http.HandleFunc("/usuarios/getoptions", controller.Usuariosgetoptions)

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
	http.HandleFunc("/autorizados", controller.Autorizados)
	http.HandleFunc("/autorizados/list", controller.AutorizadoList)
	http.HandleFunc("/autorizados/create", controller.AutorizadoCreate)
	http.HandleFunc("/autorizados/update", controller.AutorizadoUpdate)
	http.HandleFunc("/autorizados/delete", controller.AutorizadoDelete)
	http.HandleFunc("/autorizados/getoptions", controller.Autorizadogetoptions)

	http.HandleFunc("/reservas", controller.Reservas)
	http.HandleFunc("/reservas/list", controller.ReservasList)
	http.HandleFunc("/reservas/create", controller.ReservasCreate)
	http.HandleFunc("/reservas/update", controller.ReservasUpdate)
	http.HandleFunc("/reservas/delete", controller.ReservasDelete)
	http.HandleFunc("/reservas/getoptions", controller.Reservasgetoptions)
	// Apis pagos
	http.HandleFunc("/pagos", controller.Pagos)
	http.HandleFunc("/pagos/list", controller.PagosList)
	http.HandleFunc("/pagos/create", controller.PagosCreate)
	http.HandleFunc("/pagos/update", controller.PagosUpdate)
	http.HandleFunc("/pagos/delete", controller.PagosDelete)
	// Apis roles de usuario
	http.HandleFunc("/usuariosroles", controller.UsuarioRoles)
	http.HandleFunc("/usuariosroles/list", controller.UsuarioRolesList)
	http.HandleFunc("/usuariosroles/create", controller.UsuarioRolesCreate)
	http.HandleFunc("/usuariosroles/update", controller.UsuarioRolesUpdate)
	http.HandleFunc("/usuariosroles/delete", controller.UsuarioRolesDelete)
	http.HandleFunc("/usuariosroles/getoptions", controller.UsuarioRolesgetoptions)

	// Apis tiposPago
	http.HandleFunc("/tipospagos", controller.TiposPago)
	http.HandleFunc("/tipospagos/list", controller.TiposPagoList)
	http.HandleFunc("/tipospagos/create", controller.TiposPagoCreate)
	http.HandleFunc("/tipospagos/update", controller.TiposPagoUpdate)
	http.HandleFunc("/tipospagos/delete", controller.TiposPagoDelete)
	http.HandleFunc("/tipospagos/getoptions", controller.TiposPagogetoptions)

	//Apis menus
	http.HandleFunc("/menus", controller.Menus)
	http.HandleFunc("/menus/list", controller.MenusList)
	http.HandleFunc("/menus/create", controller.MenusCreate)
	http.HandleFunc("/menus/update", controller.MenusUpdate)
	http.HandleFunc("/menus/delete", controller.MenusDelete)
	// sumministra los nombres de los menus
	http.HandleFunc("/menus/getoptions", controller.MenusgetoptionsMenuParent)

	// Apis tiposevento
	http.HandleFunc("/tiposeventos", controller.Tiposevento)
	http.HandleFunc("/tiposeventos/list", controller.TiposeventoList)
	http.HandleFunc("/tiposeventos/create", controller.TiposeventoCreate)
	http.HandleFunc("/tiposeventos/update", controller.TiposeventoUpdate)
	http.HandleFunc("/tiposeventos/delete", controller.TiposeventoDelete)
	http.HandleFunc("/tiposeventos/getoptions", controller.TiposeventogetOptions)

	//Apis espacios
	http.HandleFunc("/espacios", controller.Espacio)
	http.HandleFunc("/espacios/list", controller.EspacioList)
	http.HandleFunc("/espacios/create", controller.EspacioCreate)
	http.HandleFunc("/espacios/update", controller.EspacioUpdate)
	http.HandleFunc("/espacios/delete", controller.EspaciosDelete)
	http.HandleFunc("/espacios/getoptions", controller.Espaciosgetoptions)

	//Apis horarios
	http.HandleFunc("/horarios", controller.Horarios)
	http.HandleFunc("/horarios/list", controller.HorariosList)
	http.HandleFunc("/horarios/create", controller.HorariosCreate)
	http.HandleFunc("/horarios/update", controller.HorariosUpdate)
	http.HandleFunc("/horarios/delete", controller.HorariosDelete)

	http.ListenAndServe(":3000", nil)
}
