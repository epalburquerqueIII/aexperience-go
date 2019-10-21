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
	http.HandleFunc("/usuarios", controller.Usuarios)
	http.HandleFunc("/usuarios/list", controller.UsuariosList)
	http.HandleFunc("/usuarios/create", controller.UsuariosCreate)
	http.HandleFunc("/usuarios/update", controller.UsuariosUpdate)
	http.HandleFunc("/usuarios/delete", controller.UsuariosDelete)
	http.HandleFunc("/usuarios/register", controller.UsuariosRegister)
	http.HandleFunc("/usuarios/getoptions", controller.Usuariosgetoptions)

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
	http.HandleFunc("/autorizados/list", controller.AutorizadosList)
	http.HandleFunc("/autorizados/create", controller.AutorizadosCreate)
	http.HandleFunc("/autorizados/update", controller.AutorizadosUpdate)
	http.HandleFunc("/autorizados/delete", controller.AutorizadosDelete)
	http.HandleFunc("/autorizados/getoptions", controller.Autorizadosgetoptions)

	//Eventos
	http.HandleFunc("/eventos/getEventosmdtojson", controller.GetEventosmdtojson)

	// Apis reservas
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

	// Apis pagos pendientes
	http.HandleFunc("/pagospendientes", controller.PagosPendientes)
	http.HandleFunc("/pagospendientes/list", controller.PagosPendientesList)
	http.HandleFunc("/pagospendientes/create", controller.PagosPendientesCreate)
	http.HandleFunc("/pagospendientes/update", controller.PagosPendientesUpdate)
	http.HandleFunc("/pagospendientes/delete", controller.PagosPendientesDelete)
	//http.HandleFunc("/pagospendientes/getoptions", controller.TiposPagogetoptions)

	// Apis roles de usuario
	http.HandleFunc("/usuariosroles", controller.UsuariosRoles)
	http.HandleFunc("/usuariosroles/list", controller.UsuariosRolesList)
	http.HandleFunc("/usuariosroles/create", controller.UsuariosRolesCreate)
	http.HandleFunc("/usuariosroles/update", controller.UsuariosRolesUpdate)
	http.HandleFunc("/usuariosroles/delete", controller.UsuariosRolesDelete)
	http.HandleFunc("/usuariosroles/getoptions", controller.UsuariosRolesgetoptions)

	// Apis tiposPago
	http.HandleFunc("/tipospago", controller.TiposPago)
	http.HandleFunc("/tipospago/list", controller.TiposPagoList)
	http.HandleFunc("/tipospago/create", controller.TiposPagoCreate)
	http.HandleFunc("/tipospago/update", controller.TiposPagoUpdate)
	http.HandleFunc("/tipospago/delete", controller.TiposPagoDelete)
	http.HandleFunc("/tipospago/getoptions", controller.TiposPagogetoptions)

	//Apis menus
	http.HandleFunc("/menus", controller.Menus)
	http.HandleFunc("/menus/list", controller.MenusList)
	http.HandleFunc("/menus/create", controller.MenusCreate)
	http.HandleFunc("/menus/update", controller.MenusUpdate)
	http.HandleFunc("/menus/delete", controller.MenusDelete)
	http.HandleFunc("/menus/getoptions", controller.MenusgetoptionsMenuParent)

	// Apis tiposeventos
	http.HandleFunc("/tiposeventos", controller.Tiposeventos)
	http.HandleFunc("/tiposeventos/list", controller.TiposeventosList)
	http.HandleFunc("/tiposeventos/create", controller.TiposeventosCreate)
	http.HandleFunc("/tiposeventos/update", controller.TiposeventosUpdate)
	http.HandleFunc("/tiposeventos/delete", controller.TiposeventosDelete)
	http.HandleFunc("/tiposeventos/getoptions", controller.TiposeventosgetOptions)

	//Apis espacios
	http.HandleFunc("/espacios", controller.Espacios)
	http.HandleFunc("/espacios/list", controller.EspaciosList)
	http.HandleFunc("/espacios/create", controller.EspaciosCreate)
	http.HandleFunc("/espacios/update", controller.EspaciosUpdate)
	http.HandleFunc("/espacios/delete", controller.EspaciosDelete)
	http.HandleFunc("/espacios/getoptions", controller.Espaciosgetoptions)

	//Apis horarios
	http.HandleFunc("/horarios", controller.Horarios)
	http.HandleFunc("/horarios/list", controller.HorariosList)
	http.HandleFunc("/horarios/create", controller.HorariosCreate)
	http.HandleFunc("/horarios/update", controller.HorariosUpdate)
	http.HandleFunc("/horarios/delete", controller.HorariosDelete)

	//Apis menu roles
	http.HandleFunc("/menuroles", controller.MenuRoles)
	http.HandleFunc("/menuroles/list", controller.MenuRolesList)
	http.HandleFunc("/menuroles/create", controller.MenuRolesCreate)
	http.HandleFunc("/menuroles/update", controller.MenuRolesUpdate)
	http.HandleFunc("/menuroles/delete", controller.MenuRolesDelete)
	http.HandleFunc("/menuroles/getoptions", controller.MenuRolesGetOptions)

	//Apis newsletter
	http.HandleFunc("/newsletter", controller.Newsletter)
	http.HandleFunc("/newsletter/list", controller.NewsletterList)
	http.HandleFunc("/newsletter/create", controller.NewsletterCreate)
	http.HandleFunc("/newsletter/update", controller.NewsletterUpdate)
	http.HandleFunc("/newsletter/delete", controller.NewsletterDelete)
	http.HandleFunc("/newsletter/getoptions", controller.NewslettergetoptionsTipoNoticias)

	//Radio button emailnewsletter
	http.HandleFunc("/emailnewsletter", controller.EmailNewsletter)
	// Otras apis
	http.HandleFunc("/estadisticas", controller.Estadisticas)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/registro", controller.Registro)
	http.HandleFunc("/404", controller.Errorpag)
	http.HandleFunc("/recuperarcontrasena", controller.Recuperarcontrasena)
	http.HandleFunc("/paginavacia", controller.Paginavacia)
	http.HandleFunc("/iva", controller.Iva)

	http.ListenAndServe(":3000", nil)
}
