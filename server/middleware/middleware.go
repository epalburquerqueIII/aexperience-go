package middleware

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
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
		activaRestricted := true

		requestCsrfToken := grabCsrfFromReq(r)
		if r.URL.Path == "/autorizados/list" || r.URL.Path == "/pagos" {
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
				"/pagos", //Pagos
				"/pagos/list",
				"/pagos/create",
				"/pagos/update",
				"/pagos/delete",
				//"/usuarios", //Usuarios
				"/usuarios/list",
				"/usuarios/create",
				"/usuarios/update",
				//"/usuarios/delete",
				"/usuarios/register",
				"/usuarios/getoptions",
				"/usuarios/registerUI",
				"/consumobonos", //Consumo bonos
				"/consumobonos/list",
				"/consumobonos/create",
				"/consumobonos/update",
				"/bonos", //Bonos
				"/bonos/list",
				"/bonos/create",
				"/bonos/update",
				"/bonos/delete",
				"/autorizados", //Autorizados
				"/autorizados/list",
				"/autorizados/create",
				"/autorizados/update",
				"/autorizados/delete",
				"/autorizados/getoptions",
				"/eventos/getEventosmdtojson", //Eventos
				"/reservas",                   //Reservas
				"/reservas/list",
				"/reservas/create",
				"/reservas/update",
				"/reservas/delete",
				"/reservas/getoptions",
				"/reservas/comprarbono",
				"/pagospendientes", //Pagos pendientes
				"/pagospendientes/list",
				"/pagospendientes/getoptions",
				"/usuariosroles", //Roles de usuario
				"/usuariosroles/list",
				"/usuariosroles/create",
				"/usuariosroles/update",
				"/usuariosroles/delete",
				"/usuariosroles/getoptions",
				"/tipospago", //Tipos de pago
				"/tipospago/list",
				"/tipospago/create",
				"/tipospago/update",
				"/tipospago/delete",
				"/tipospago/getoptions",
				"/menus", //Menus
				"/menus/list",
				"/menus/create",
				"/menus/update",
				"/menus/delete",
				"/menus/getoptions",
				"/tiposeventos", //Tipo eventos
				"/tiposeventos/list",
				"/tiposeventos/create",
				"/tiposeventos/update",
				"/tiposeventos/delete",
				"/tiposeventos/getoptions",
				"/espacios", //Espacios
				"/espacios/list",
				"/espacios/create",
				"/espacios/update",
				"/espacios/delete",
				"/espacios/getoptions",
				"/horarios", //Horarios
				"/horarios/list",
				"/horarios/create",
				"/horarios/update",
				"/horarios/delete",
				"/menuroles", //Menu roles
				"/menuroles/list",
				"/menuroles/create",
				"/menuroles/update",
				"/menuroles/delete",
				"/menuroles/getoptions",
				"/newsletter", //Newsletter
				"/newsletter/list",
				"/newsletter/create",
				"/newsletter/update",
				"/newsletter/delete",
				"/newsletter/getoptions",
				"/newsletter/newsletterguardar",
				"/tiponoticias", //Tipo noticias
				"/tiponoticias/list",
				"/horasdia", //Horas del día
				"/horasdia/list",
				"/horasdia/create",
				"/horasdia/update",
				"/estadisticas", //Otras
				"/login",
				"/404",
				"/recuperarcontrasena",
				"/paginavacia",
				"/iva",
				"/usuarios":
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

	var tmpl = template.Must(template.ParseGlob("./views/*.html"))

	authweb := model.AuthWeb{grabCsrfFromReq(r), user.UserName}
	menu := util.Menus(user.Role)

	// @adam-hanna: I shouldn't be doing this in my middleware!!!!
	switch r.URL.Path {
	case "/restricted":
		w.Header().Set("X-CSRF-Token", authweb.CsrfSecret)
		error := tmpl.ExecuteTemplate(w, "restricted", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}

	//--- GESTIONES ---
	//Gestiona los pagos:
	case "/pagos":
		w.Header().Set("X-CSRF-Token", authweb.CsrfSecret)
		error := tmpl.ExecuteTemplate(w, "pagos", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
		error := tmpl.ExecuteTemplate(w, "usuarios", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
	case "/usuarios/registerUI":
		controller.UsuariosUIRegister(w, r)

	//Gestiona el consumo de bonos:
	case "/consumobonos":
		error := tmpl.ExecuteTemplate(w, "consumobonos", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
	case "/consumobonos/list":
		controller.ConsumoBonosList(w, r)
	case "/consumobonos/create":
		controller.ConsumoBonosCreate(w, r)
	case "/consumobonos/update":
		controller.ConsumoBonosUpdate(w, r)

	//Gestiona los bonos:
	case "/bonos":
		error := tmpl.ExecuteTemplate(w, "bonos", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
		error := tmpl.ExecuteTemplate(w, "autorizados", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
	case "/getEventosmdtojson":
		controller.GetEventosmdtojson(w, r)

	//Gestiona las reservas:
	case "/reservas":
		error := tmpl.ExecuteTemplate(w, "reservas", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
		error := tmpl.ExecuteTemplate(w, "pagospendientes", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
	case "/pagospendientes/list":
		controller.PagosPendientesList(w, r)
	case "/pagospendientes/getoptions":
		controller.Pagospendientesgetoptions(w, r)

	//Gestiona los roles de usuario:
	case "/usuariosroles":
		error := tmpl.ExecuteTemplate(w, "usuariosroles", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
	case "/tipospago":
		error := tmpl.ExecuteTemplate(w, "tipospago", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
	case "/tipospago/list":
		controller.TiposPagoList(w, r)
	case "/tipospago/create":
		controller.TiposPagoCreate(w, r)
	case "/tipospago/update":
		controller.TiposPagoUpdate(w, r)
	case "/tipospago/delete":
		controller.TiposPagoDelete(w, r)
	case "/tipospago/getoptions":
		controller.TiposPagogetoptions(w, r)

	//Gestiona los menús:
	case "/menus":
		error := tmpl.ExecuteTemplate(w, "menus", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
		error := tmpl.ExecuteTemplate(w, "tiposeventos", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
		error := tmpl.ExecuteTemplate(w, "espacios", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
		error := tmpl.ExecuteTemplate(w, "horarios", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
		error := tmpl.ExecuteTemplate(w, "menuroles", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
		error := tmpl.ExecuteTemplate(w, "newsletter", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
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
		error := tmpl.ExecuteTemplate(w, "tiponoticias", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
	case "/tiponoticias/list":
		controller.TipoNoticiasList(w, r)

	//Gestiona las horas del día:
	case "/horasdia":
		error := tmpl.ExecuteTemplate(w, "horasdia", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
	case "/horasdia/list":
		controller.HorasDiaList(w, r)
	case "/horasdia/create":
		controller.HorasDiaCreate(w, r)
	case "/horasdia/update":
		controller.HorasDiaUpdate(w, r)

	//Gestiona otras apis:
	case "/estadisticas":
		error := tmpl.ExecuteTemplate(w, "estadisticas", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
	case "/404":
		error := tmpl.ExecuteTemplate(w, "404", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
	case "/recuperarcontrasena":
		error := tmpl.ExecuteTemplate(w, "recuperarcontrasena", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
	case "/paginavacia":
		error := tmpl.ExecuteTemplate(w, "paginavacia", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}
	case "/iva":
		error := tmpl.ExecuteTemplate(w, "iva", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}

	//Gestiona el login:
	case "/login":
		switch r.Method {
		case "GET":
			tmpl.ExecuteTemplate(w, "login", &templates.LoginPage{false, ""})

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
	case "/register":
		switch r.Method {
		case "GET":
			templates.RenderTemplate(w, "register", &templates.RegisterPage{false, ""})

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
