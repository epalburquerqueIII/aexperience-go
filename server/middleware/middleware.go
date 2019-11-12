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
		activa_restricted := true

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
			activa_restricted = false
		}

		if activa_restricted {

			switch r.URL.Path {
			case "/restricted", "/deleteUser",
				"/pagos",
				"/pagos/list",
				"/pagos/create",
				"/pagos/update",
				"/pagos/delete",
				"/autorizados/list",
				"/autorizados/create",
				"/autorizados/update",
				"/autorizados/delete",
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
		// Gestiona los pagos
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
	case "/autorizados/list":
		controller.AutorizadosList(w, r)
	case "/autorizados/create":
		controller.AutorizadosCreate(w, r)
	case "/autorizados/update":
		controller.AutorizadosUpdate(w, r)
	case "/autorizados/delete":
		controller.AutorizadosDelete(w, r)
	case "/usuarios":
		error := tmpl.ExecuteTemplate(w, "usuarios", &templates.RestrictedPage{authweb, menu})
		if error != nil {
			log.Println("Error ", error.Error)
		}

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
