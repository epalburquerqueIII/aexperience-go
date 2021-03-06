package middleware

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"../../config"
	"../../controller"
	"../../model"
	"../../model/authdb"
	"../../server/middleware/myJwt"
	"../../server/templates"

	"../../util"
	"github.com/justinas/alice"
)

var user = model.User{"", "", "Default", 1, -1}
var uuid string

// NewHandler Punto de llamadas desde el servidor
func NewHandler() http.Handler {
	// this returns a handler that is a chain of all of our other handlers
	return alice.New(recoverHandler, authHandler).ThenFunc(logicHandler)
}

func recoverHandler(next http.Handler) http.Handler {
	// this catches any errors and returns an internal server error to the client
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Panic("Recovered! Panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func authHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var authTokenString, refreshTokenString, csrfSecret string
		// do auth stuff
		// include list of restricted paths, comma sep
		// I'm including logout here, bc I don't want a baddie forcing my users to logout
		activaRestricted := false

		requestCsrfToken := grabCsrfFromReq(r)
		if r.URL.Path == "/pagos/list" || r.URL.Path == "/pagos" {
			//				for _, cookie := range r.Cookies() {
			//					log.Printf("Found a cookie named: %s,%s\n", cookie.Name, cookie.Value)
			//				}
			log.Println("List")
			buf, _ := ioutil.ReadAll(r.Body)
			log.Println(buf)
		}
		// A M
		if (requestCsrfToken != "") && (requestCsrfToken[4] == '$') {
			activaRestricted = false
		}

		if activaRestricted {

			switch r.URL.Path {
			case "/restricted", "/deleteUser",
				//Pagos
				"/pagos",
				"/pagos/list",
				"/pagos/create",
				"/pagos/update",
				"/pagos/delete",
				//Usuarios
				"/usuarios",
				"/usuarios/list",
				"/usuarios/create",
				"/usuarios/update",
				//"/usuarios/delete",
				"/usuarios/getoptions",
				//Consumo bonos
				"/consumobonos",
				"/consumobonos/list",
				"/consumobonos/create",
				"/consumobonos/update",
				//Bonos
				"/bonos",
				"/bonos/list",
				"/bonos/create",
				"/bonos/update",
				"/bonos/delete",
				//Autorizados
				"/autorizados",
				"/autorizados/list",
				"/autorizados/create",
				"/autorizados/update",
				"/autorizados/delete",
				"/autorizados/getoptions",
				//Reservas
				"/reservas",
				"/reservas/list",
				"/reservas/create",
				"/reservas/update",
				"/reservas/delete",
				"/reservas/getoptions",
				"/reservas/comprarbono",
				//Pagos pendientes
				"/pagospendientes",
				"/pagospendientes/list",
				"/pagospendientes/confirmarpago",
				//Roles de usuario
				"/usuariosroles",
				"/usuariosroles/list",
				"/usuariosroles/create",
				"/usuariosroles/update",
				"/usuariosroles/delete",
				"/usuariosroles/getoptions",
				//Tipos de pago
				"/tipospagos",
				"/tipospagos/list",
				"/tipospagos/create",
				"/tipospagos/update",
				"/tipospagos/delete",
				"/tipospagos/getoptions",
				//Menus
				"/menus",
				"/menus/list",
				"/menus/create",
				"/menus/update",
				"/menus/delete",
				"/menus/getoptions",
				//Tipo eventos
				"/tiposeventos",
				"/tiposeventos/list",
				"/tiposeventos/create",
				"/tiposeventos/update",
				"/tiposeventos/delete",
				"/tiposeventos/getoptions",
				//Espacios
				"/espacios",
				"/espacios/list",
				"/espacios/create",
				"/espacios/update",
				"/espacios/delete",
				"/espacios/getoptions",
				//Horarios
				"/horarios",
				"/horarios/list",
				"/horarios/create",
				"/horarios/update",
				"/horarios/delete",
				//Menu roles
				"/menuroles",
				"/menuroles/list",
				"/menuroles/create",
				"/menuroles/update",
				"/menuroles/delete",
				"/menuroles/getoptions",
				//Newsletter
				"/newsletter",
				"/newsletter/list",
				"/newsletter/create",
				"/newsletter/update",
				"/newsletter/delete",
				//Tipo noticias
				"/tiponoticias",
				"/tiponoticias/list",
				//Horas del día
				"/reservapabellonpista",
				"/horasreservables",
				"/movilhorasreservables",
				"/reservapabellonpista/create",
				//Otras
				"/estadisticas",
				"/404",
				"/recuperarcontrasena",
				"/paginavacia",
				"/iva":
				//, "/logout"
				// Login desde otra plataforma o segmento de red, no acabada
				// fuerza el cambio de web de Hugo a privada en el 8088
				if r.URL.Path == "/pagos/list" || r.URL.Path == "/pagos" {
					//				for _, cookie := range r.Cookies() {
					//					log.Printf("Found a cookie named: %s,%s\n", cookie.Name, cookie.Value)
					//				}
					log.Println("List")
					buf, _ := ioutil.ReadAll(r.Body)
					log.Println(buf)
				}
				if r.URL.Path == "/restricted" && requestCsrfToken == "8088" {
					// grabar la cookie en el port 8088
					r.ParseForm()
					authTokenString = strings.Join(r.Form["data0"], "")
					refreshTokenString = strings.Join(r.Form["data1"], "")
					if authTokenString != "" {
						setAuthAndRefreshCookies(&w, authTokenString, refreshTokenString)
					}
					w.Header().Set("X-CSRF-Token", requestCsrfToken)
				}

				log.Println("In auth restricted section")

				// read cookies
				AuthCookie, authErr := r.Cookie("AuthToken")
				if authErr == http.ErrNoCookie {
					// set the cookies to these newly created jwt's

					log.Println("Unauthorized attempt! No auth cookie")
					nullifyTokenCookies(&w, r)
					// http.Redirect(w, r, "/login", 302)
					http.Error(w, http.StatusText(401), 401)
					return
				} else if authErr != nil {
					log.Panic("panic: %+v", authErr)
					nullifyTokenCookies(&w, r)
					http.Error(w, http.StatusText(500), 500)
					return
				}

				RefreshCookie, refreshErr := r.Cookie("RefreshToken")
				if refreshErr == http.ErrNoCookie {
					log.Println("Unauthorized attempt! No refresh cookie")
					nullifyTokenCookies(&w, r)
					http.Redirect(w, r, "/login", 302)
					return
				} else if refreshErr != nil {
					log.Panic("panic: %+v", refreshErr)
					nullifyTokenCookies(&w, r)
					http.Error(w, http.StatusText(500), 500)
					return
				}

				// grab the csrf token
				log.Println(requestCsrfToken)
				// Acceso desde otra plataforma o puerto, no acabado
				if requestCsrfToken != "8088" {
					// check the jwt's for validity
					var err error
					authTokenString, refreshTokenString, csrfSecret, err = myJwt.CheckAndRefreshTokens(AuthCookie.Value, RefreshCookie.Value, requestCsrfToken)
					if err != nil {
						if err.Error() == "Unauthorized" {
							log.Println("Unauthorized attempt! JWT's not valid!")
							// nullifyTokenCookies(&w, r)
							// http.Redirect(w, r, "/login", 302)
							http.Error(w, http.StatusText(401), 401)
							return
						} else {
							// @adam-hanna: do we 401 or 500, here?
							// it could be 401 bc the token they provided was messed up
							// or it could be 500 bc there was some error on our end
							log.Println("err not nil")
							log.Panic("panic: %+v", err)
							// nullifyTokenCookies(&w, r)
							http.Error(w, http.StatusText(500), 500)
							return
						}
					}
				} else {
					authTokenString = strings.Join(r.Form["data0"], "")
					refreshTokenString = strings.Join(r.Form["data1"], "")
					csrfSecret = strings.Join(r.Form["data2"], "")
				}

				log.Println("Successfully recreated jwts")

				// @adam-hanna: Change this. Only allow whitelisted origins! Also check referer header
				w.Header().Set("Access-Control-Allow-Origin", "*")

				// if we've made it this far, everything is valid!
				// And tokens have been refreshed if need-be
				setAuthAndRefreshCookies(&w, authTokenString, refreshTokenString)
				w.Header().Set("X-CSRF-Token", csrfSecret)

			default:
				// no jwt check necessary
			}
		} // if Activa seguridad

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// Enable CORS
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, X-CSRF-Token, Authorization, access-control-allow-headers")
	(*w).Header().Set("Access-Control-Request-Headers", "X-Requested-With, accept, content-type, X-CSRF-Token")
}

func logicHandler(w http.ResponseWriter, r *http.Request) {

	authweb := model.AuthWeb{grabCsrfFromReq(r), user.UserName}
	menu := util.Menus(user.Role)

	// @adam-hanna: I shouldn't be doing this in my middleware!!!!
	switch r.URL.Path {
	case "/restricted":
		w.Header().Set("X-CSRF-Token", authweb.CsrfSecret)
		templates.RenderTemplate(w, "restricted", &templates.RestrictedPage{authweb, menu})

	case "/comprarbonos":

		type Datos struct {
			IDUsuario int
			Bonos     []model.Tbono
		}
		datos := Datos{}
		datos.IDUsuario = 18
		datos.Bonos = controller.GetBonos()
		type Infobonos struct {
			AuthWeb model.AuthWeb
			Menus   []model.Tmenuconfig
			Params  Datos
		}
		templates.RenderTemplate(w, "comprarbonos", &Infobonos{authweb, menu, datos})

	//--- GESTIONES ---
	//Gestiona los pagos:
	case "/pagos":
		w.Header().Set("X-CSRF-Token", authweb.CsrfSecret)
		templates.RenderTemplate(w, "pagos", &templates.RestrictedPage{authweb, menu})
	case "/pagos/list":
		controller.PagosList(w, r)
	case "/pagos/create":
		controller.PagosCreate(w, r)
	case "/pagos/update":
		controller.PagosUpdate(w, r)
	case "/pagos/delete":
		controller.PagosDelete(w, r)

	//Gestiona los usuarios:
	case "/usuarios":
		templates.RenderTemplate(w, "usuarios", &templates.RestrictedPage{authweb, menu})
	case "/usuarios/list":
		controller.UsuariosList(w, r)
	case "/usuarios/create":
		controller.UsuariosCreate(w, r)
	case "/usuarios/update":
		controller.UsuariosUpdate(w, r)
	case "/usuarios/delete":
		controller.UsuariosDelete(w, r)
	case "/usuarios/register":
		controller.UsuariosRegister(w, r)
	case "/usuarios/getoptions":
		controller.Usuariosgetoptions(w, r)
	//Gestiona el consumo de bonos:
	case "/consumobonos":
		templates.RenderTemplate(w, "consumobonos", &templates.RestrictedPage{authweb, menu})
	case "/consumobonos/list":
		controller.ConsumoBonosList(w, r)
	case "/consumobonos/create":
		controller.ConsumoBonosCreate(w, r)
	case "/consumobonos/update":
		controller.ConsumoBonosUpdate(w, r)

	//Gestiona los bonos:
	case "/bonos":
		templates.RenderTemplate(w, "bonos", &templates.RestrictedPage{authweb, menu})
	case "/bonos/list":
		controller.BonoList(w, r)
	case "/bonos/create":
		controller.BonoCreate(w, r)
	case "/bonos/update":
		controller.BonoUpdate(w, r)
	case "/bonos/delete":
		controller.BonoDelete(w, r)

	//Gestiona los autorizados:
	case "/autorizados":
		templates.RenderTemplate(w, "autorizados", &templates.RestrictedPage{authweb, menu})
	case "/autorizados/list":
		controller.AutorizadosList(w, r)
	case "/autorizados/create":
		controller.AutorizadosCreate(w, r)
	case "/autorizados/update":
		controller.AutorizadosUpdate(w, r)
	case "/autorizados/delete":
		controller.AutorizadosDelete(w, r)
	case "/autorizados/getoptions":
		controller.Autorizadosgetoptions(w, r)

	//Gestiona los eventos:
	case "/eventos/getEventosmdtojson":
		controller.GetEventosmdtojson(w, r)

	//Gestiona las reservas:
	case "/reservas":
		templates.RenderTemplate(w, "reservas", &templates.RestrictedPage{authweb, menu})
	case "/reservas/list":
		controller.ReservasList(w, r)
	case "/reservas/create":
		controller.ReservasCreate(w, r)
	case "/reservas/update":
		controller.ReservasUpdate(w, r)
	case "/reservas/delete":
		controller.ReservasDelete(w, r)
	case "/reservas/getoptions":
		controller.Reservasgetoptions(w, r)
	case "/reservas/comprarbono":
		controller.ComprarBono(w, r)

	//Gestiona los pagos pendientes:
	case "/pagospendientes":
		templates.RenderTemplate(w, "pagospendientes", &templates.RestrictedPage{authweb, menu})
	case "/pagospendientes/list":
		controller.PagosPendientesList(w, r)
	case "/pagospendientes/confirmarpago":
		controller.Pagospendientesconfirmarpago(w, r)

	//Gestiona los roles de usuario:
	case "/usuariosroles":
		templates.RenderTemplate(w, "usuariosroles", &templates.RestrictedPage{authweb, menu})
	case "/usuariosroles/list":
		controller.UsuariosRolesList(w, r)
	case "/usuariosroles/create":
		controller.UsuariosRolesCreate(w, r)
	case "/usuariosroles/update":
		controller.UsuariosRolesUpdate(w, r)
	case "/usuariosroles/delete":
		controller.UsuariosRolesDelete(w, r)
	case "/usuariosroles/getoptions":
		controller.UsuariosRolesgetoptions(w, r)

	//Gestiona los tipos de pago:
	case "/tipospagos":
		templates.RenderTemplate(w, "tipospagos", &templates.RestrictedPage{authweb, menu})
	case "/tipospagos/list":
		controller.TiposPagoList(w, r)
	case "/tipospagos/create":
		controller.TiposPagoCreate(w, r)
	case "/tipospagos/update":
		controller.TiposPagoUpdate(w, r)
	case "/tipospagos/delete":
		controller.TiposPagoDelete(w, r)
	case "/tipospagos/getoptions":
		controller.TiposPagogetoptions(w, r)

	//Gestiona los menús:
	case "/menus":
		templates.RenderTemplate(w, "menus", &templates.RestrictedPage{authweb, menu})
	case "/menus/list":
		controller.MenusList(w, r)
	case "/menus/create":
		controller.MenusCreate(w, r)
	case "/menus/update":
		controller.MenusUpdate(w, r)
	case "/menus/delete":
		controller.MenusDelete(w, r)
	case "/menus/getoptions":
		controller.MenusgetoptionsMenuParent(w, r)

	//Gestiona los tipos de eventos:
	case "/tiposeventos":
		templates.RenderTemplate(w, "tiposeventos", &templates.RestrictedPage{authweb, menu})
	case "/tiposeventos/list":
		controller.TiposeventosList(w, r)
	case "/tiposeventos/create":
		controller.TiposeventosCreate(w, r)
	case "/tiposeventos/update":
		controller.TiposeventosUpdate(w, r)
	case "/tiposeventos/delete":
		controller.TiposeventosDelete(w, r)
	case "/tiposeventos/getoptions":
		controller.TiposeventosgetOptions(w, r)

	//Gestiona los espacios:
	case "/espacios":
		templates.RenderTemplate(w, "espacios", &templates.RestrictedPage{authweb, menu})
	case "/espacios/list":
		controller.EspacioList(w, r)
	case "/espacios/create":
		controller.EspacioCreate(w, r)
	case "/espacios/update":
		controller.EspacioUpdate(w, r)
	case "/espacios/delete":
		controller.EspacioDelete(w, r)
	case "/espacios/getoptions":
		controller.Espaciosgetoptions(w, r)

	//Gestiona los horarios:
	case "/horarios":
		templates.RenderTemplate(w, "horarios", &templates.RestrictedPage{authweb, menu})
	case "/horarios/list":
		controller.HorariosList(w, r)
	case "/horarios/create":
		controller.HorariosCreate(w, r)
	case "/horarios/update":
		controller.HorariosUpdate(w, r)
	case "/horarios/delete":
		controller.HorariosDelete(w, r)

	//Gestiona los roles de menú:
	case "/menuroles":
		templates.RenderTemplate(w, "menuroles", &templates.RestrictedPage{authweb, menu})
	case "/menuroles/list":
		controller.MenuRolesList(w, r)
	case "/menuroles/create":
		controller.MenuRolesCreate(w, r)
	case "/menuroles/update":
		controller.MenuRolesUpdate(w, r)
	case "/menuroles/delete":
		controller.MenuRolesDelete(w, r)
	case "/menuroles/getoptions":
		controller.MenuRolesGetOptions(w, r)

	//Gestiona el newsletter:
	case "/newsletter":
		templates.RenderTemplate(w, "newsletter", &templates.RestrictedPage{authweb, menu})
	case "/newsletter/list":
		controller.NewsletterList(w, r)
	case "/newsletter/create":
		controller.NewsletterCreate(w, r)
	case "/newsletter/update":
		controller.NewsletterUpdate(w, r)
	case "/newsletter/delete":
		controller.NewsletterDelete(w, r)
	case "/newsletter/getoptions":
		controller.NewslettergetoptionsTipoNoticias(w, r)
	case "/newsletter/newsletterguardar":
		controller.Newsletterguardar(w, r)

	//Gestiona el tipo de noticias:
	case "/tiponoticias":

		type datos struct {
			Email    string
			Noticias []model.TtipoNoticia
		}

		templates.RenderTemplate(w, "tiponoticias", &datos{r.FormValue("EMAIL"), controller.GetTipoNoticias()})
	case "/tiponoticias/list":
		controller.TipoNoticiasList(w, r)

		//Gestiona las horas del día:
	case "/movilhorasreservables":
		controller.MovilHorasReservables(w, r)
	case "/reservapabellonpista":
		templates.RenderTemplate(w, "reservapabellonpista", &templates.RestrictedPage{authweb, menu})
	case "/horasreservables":
		// Realizado con campos ocultos, se puede hace con cookie, este sistema es más rapido
		type Datos struct {
			Fechabusqueda  string
			Espacio        string
			Horasreservada []model.THorasdia
		}
		datos := Datos{}
		type HorasReservaPage struct {
			AuthWeb model.AuthWeb
			Menus   []model.Tmenuconfig
			Params  Datos
		}
		if r.FormValue("dia") == "1" {
			datos.Fechabusqueda = time.Now().Format("2006-01-02")
		} else {
			datos.Fechabusqueda = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
		}
		if r.FormValue("espacio") == "1" {
			datos.Espacio = "5"
		} else {
			datos.Espacio = "1"
		}
		datos.Horasreservada = controller.HorasReservables(datos.Fechabusqueda, datos.Espacio)
		templates.RenderTemplate(w, "horasreservables", &HorasReservaPage{authweb, menu, datos})

	case "/reservapabellonpista/create":
		controller.HorasDiaCreate(w, r)

	//Gestiona otras apis:
	case "/estadisticas":
		templates.RenderTemplate(w, "estadisticas", &templates.RestrictedPage{authweb, menu})
	case "/404":
		templates.RenderTemplate(w, "404", &templates.RestrictedPage{authweb, menu})
	case "/recuperarcontrasena":
		templates.RenderTemplate(w, "recuperarcontrasena", &templates.RestrictedPage{authweb, menu})
	case "/paginavacia":
		templates.RenderTemplate(w, "paginavacia", &templates.RestrictedPage{authweb, menu})
	case "/iva":
		templates.RenderTemplate(w, "iva", &templates.RestrictedPage{authweb, menu})

	//Gestiona el login:
	case "/login":
		switch r.Method {
		case "GET":
			templates.RenderTemplate(w, "login", &templates.LoginPage{false, ""})

			//			templates.RenderTemplate(w, "login", &templates.LoginPage{false, ""})

		case "POST":
			r.ParseForm()
			log.Println(r.Form)
			var user model.User
			var loginErr error
			user, uuid, loginErr = authdb.LogUserIn(strings.Join(r.Form["email"], ""), strings.Join(r.Form["password"], ""))
			log.Println(user, uuid, loginErr)
			if loginErr != nil {
				// login err
				// templates.RenderTemplate(w, "login", &templates.LoginPage{ true, "Login failed\n\nIncorrect email or password" })
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("X-CSRF-Token", "Acceso no autoricado")
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				// no login err
				// now generate cookies for this user
				authTokenString, refreshTokenString, csrfSecret, err := myJwt.CreateNewTokens(uuid, user.Role)
				if err != nil {
					http.Error(w, http.StatusText(500), 500)
				}

				// set the cookies to these newly created jwt's
				setAuthAndRefreshCookies(&w, authTokenString, refreshTokenString)
				w.Header().Set("X-CSRF-Token", csrfSecret)
				w.Header().Set("Access-Control-Allow-Origin", "*")
				result := model.ResponseToken{authTokenString, refreshTokenString, csrfSecret, user.UserName, user.UserID, user.Role}
				jsonResult, err := json.Marshal(result)
				w.Write(jsonResult)
				//w.WriteHeader(http.StatusOK)
			}
			return

		case "OPTIONS":
			//w.Header().Set("Content-Type", "text/html; charset=utf-8")
			setupResponse(&w, r)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	case "/usuarios/registerUI":
		switch r.Method {
		case "GET":
			// UsuariosUserRegister Pantalla para registrar un usuario
			templates.RenderTemplate(w, "userregister", nil)
		case "POST":
			r.ParseForm()
			log.Println(r.Form)
			var err error
			// check to see if the email is already taken
			user, uuid, err = authdb.FetchUserByEmail(strings.Join(r.Form["email"], ""))
			if err == nil {
				// templates.RenderTemplate(w, "register", &templates.RegisterPage{ true, "email not available!" })
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				// nope, now create this user

				uuid, err = authdb.StoreUser(strings.Join(r.Form["email"], ""), strings.Join(r.Form["password"], ""), user.Role, user.UserName, user.UserID)
				if err != nil {
					http.Error(w, http.StatusText(500), 500)
				}
				log.Println("uuid: " + uuid)

				// now generate cookies for this user
				authTokenString, refreshTokenString, csrfSecret, err := myJwt.CreateNewTokens(uuid, user.Role)
				if err != nil {
					http.Error(w, http.StatusText(500), 500)
				}

				// set the cookies to these newly created jwt's
				setAuthAndRefreshCookies(&w, authTokenString, refreshTokenString)
				w.Header().Set("X-Requested-With", "XMLHttpRequest")
				w.Header().Set("X-CSRF-Token", csrfSecret)

				w.WriteHeader(http.StatusOK)
			}

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	case "/logout":
		// remove this user's ability to make requests
		nullifyTokenCookies(&w, r)
		// use 302 to force browser to do GET request
		// Unifico el cierre de sesión con el borrado del usuario
		//	case "/deleteUser":
		log.Println("Deleting user")

		// grab auth cookie
		AuthCookie, authErr := r.Cookie("AuthToken")
		if authErr == http.ErrNoCookie {
			log.Println("Unauthorized attempt! No auth cookie")
			nullifyTokenCookies(&w, r)
			http.Redirect(w, r, "/login", 302)
			return
		} else if authErr != nil {
			log.Panic("panic: %+v", authErr)
			nullifyTokenCookies(&w, r)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		uuid, uuidErr := myJwt.GrabUUID(AuthCookie.Value)
		if uuidErr != nil {
			log.Panic("panic: %+v", uuidErr)
			nullifyTokenCookies(&w, r)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		authdb.DeleteUser(uuid)
		// remove this user's ability to make requests
		nullifyTokenCookies(&w, r)
		// use 302 to force browser to do GET request
		http.Redirect(w, r, config.CmsHost, 302)

	default:
		w.WriteHeader(http.StatusOK)
	}
}

func nullifyTokenCookies(w *http.ResponseWriter, r *http.Request) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(*w, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(*w, &refreshCookie)

	// if present, revoke the refresh cookie from our db
	RefreshCookie, refreshErr := r.Cookie("RefreshToken")
	if refreshErr == http.ErrNoCookie {
		// do nothing, there is no refresh cookie present
		return
	} else if refreshErr != nil {
		log.Panic("panic: %+v", refreshErr)
		http.Error(*w, http.StatusText(500), 500)
	}

	myJwt.RevokeRefreshToken(RefreshCookie.Value)
}

func setAuthAndRefreshCookies(w *http.ResponseWriter, authTokenString string, refreshTokenString string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authTokenString,
		HttpOnly: true,
	}

	http.SetCookie(*w, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshTokenString,
		HttpOnly: true,
	}

	http.SetCookie(*w, &refreshCookie)
}

func grabCsrfFromReq(r *http.Request) string {
	csrfFromFrom := r.FormValue("X-CSRF-Token")

	if csrfFromFrom != "" {
		return csrfFromFrom
	} else {
		return r.Header.Get("X-CSRF-Token")
	}
}
