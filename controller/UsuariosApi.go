//TODO dar de alta los usuarios que están de baja
package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"../model"
	"../model/database"
	"../util"
)

var tmpl = template.Must(template.ParseGlob("views/*.html"))

// Usuario Pantalla de tratamiento de usuario
func Usuarios(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)
	error := tmpl.ExecuteTemplate(w, "usuarios", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// UsuarioList - json con los datos de clientes
func UsuariosList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT usuarios.id, nombre, nif, email, fechaNacimiento, idusuariorol, telefono, password, sesionesbonos, newsletter, fechaBaja FROM usuarios " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	usu := model.Tusuario{}
	res := []model.Tusuario{}
	for selDB.Next() {

		err = selDB.Scan(&usu.ID, &usu.Nombre, &usu.Nif, &usu.Email, &usu.FechaNacimiento, &usu.IDUsuarioRol, &usu.Telefono, &usu.Password, &usu.SesionesBonos, &usu.Newsletter, &usu.FechaBaja)
		//Si no hay fecha de baja, este campo aparece como activo
		if usu.FechaBaja == "0000-00-00" {
			usu.FechaBaja = "Activo"
		} else {
			//Formato de fecha en español cuando está de baja
			t, _ := time.Parse("2006-01-02", usu.FechaBaja)
			usu.FechaBaja = t.Format("02-01-2006")
		}

		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Cargando registros de Usuarios"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, usu)
		i++
	}

	var vrecords model.UsuarioRecords
	vrecords.Result = "OK"
	vrecords.TotalRecordCount = i
	vrecords.Records = res
	// create json response from struct
	a, err := json.Marshal(vrecords)
	// Visualza
	s := string(a)
	fmt.Println(s)
	w.Write(a)
	defer db.Close()
}

// UsuarioCreate - Crear un Usuario
func UsuariosCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	usu := model.Tusuario{}
	if r.Method == "POST" {
		usu.Nombre = r.FormValue("Nombre")
		usu.Nif = r.FormValue("Nif")
		usu.Email = r.FormValue("Email")
		usu.FechaNacimiento = util.DateSql(r.FormValue("FechaNacimiento"))
		usu.IDUsuarioRol, _ = strconv.Atoi(r.FormValue("idUsuarioRol"))
		usu.Telefono = r.FormValue("Telefono")
		usu.Password = r.FormValue("Password")
		usu.SesionesBonos, _ = strconv.Atoi(r.FormValue("SesionesBonos"))
		usu.Newsletter, _ = strconv.Atoi(r.FormValue("Newsletter"))
		usu.FechaBaja = r.FormValue("FechaBaja")
		insForm, err := db.Prepare("INSERT INTO usuarios(nombre, nif, email, fechaNacimiento, idusuariorol, telefono, password, sesionesBonos, newsletter, fechaBaja) VALUES(?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Usuario"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(usu.Nombre, usu.Nif, usu.Email, usu.FechaNacimiento, usu.IDUsuarioRol, usu.Telefono, usu.Password, usu.SesionesBonos, usu.Newsletter, usu.FechaBaja)
		if err1 != nil {
			panic(err1.Error())
		}
		usu.ID, err1 = res.LastInsertId()
		log.Println("INSERT: nombre: " + usu.Nombre + " | nif: " + usu.Nif)

	}
	var vrecord model.UsuarioRecord
	vrecord.Result = "OK"
	vrecord.Record = usu
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// UsuarioUpdate Actualiza el usuario
func UsuariosUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	usu := model.Tusuario{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("ID"))
		usu.ID = int64(i)
		usu.Nombre = r.FormValue("Nombre")
		usu.Nif = r.FormValue("Nif")
		usu.Email = r.FormValue("Email")
		usu.FechaNacimiento = util.DateSql(r.FormValue("FechaNacimiento"))
		usu.IDUsuarioRol, _ = strconv.Atoi(r.FormValue("idUsuarioRol"))
		usu.Telefono = r.FormValue("Telefono")
		usu.Password = r.FormValue("Password")
		usu.SesionesBonos, _ = strconv.Atoi(r.FormValue("SesionesBonos"))
		usu.Newsletter, _ = strconv.Atoi(r.FormValue("Newsletter"))
		usu.FechaBaja = r.FormValue("FechaBaja")
		insForm, err := db.Prepare("UPDATE usuarios SET nombre=?, nif=?, email=?, fechanacimiento =?, idusuariorol =?, telefono=?, password=?, sesionesBonos=?, newsletter=?, fechaBaja=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(usu.Nombre, usu.Nif, usu.Email, usu.FechaNacimiento, usu.IDUsuarioRol, usu.Telefono, usu.Password, usu.SesionesBonos, usu.Newsletter, usu.FechaBaja, usu.ID)
		log.Println("UPDATE: nombre: " + usu.Nombre + " | nif: " + usu.Nif)
	}
	defer db.Close()
	var vrecord model.UsuarioRecord
	vrecord.Result = "OK"
	vrecord.Record = usu
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

// UsuariosUserRegister Pantalla para registrar un usuario
func UsuariosUIRegister(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "userregister", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// UsuarioRegister - registra un Usuario
func UsuariosRegister(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	usu := model.Tusuario{}
	if r.Method == "POST" {
		usu.Nombre = r.FormValue("Nombre")
		usu.Nif = r.FormValue("Nif")
		usu.Email = r.FormValue("Email")
		usu.FechaNacimiento = util.DateSql(r.FormValue("FechaNacimiento"))
		usu.Telefono = r.FormValue("Telefono")
		usu.Password = r.FormValue("Password")

		usu.IDUsuarioRol = 1
		usu.SesionesBonos = 0
		usu.Newsletter = 1
		usu.FechaBaja = util.DateSql("00-00-0000")

		insForm, err := db.Prepare("INSERT INTO usuarios(nombre, nif, email, fechaNacimiento, idusuariorol, telefono, password,sesionesbonos,newsletter,fechabaja) VALUES(?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error registrando Usuario"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(usu.Nombre, usu.Nif, usu.Email, usu.FechaNacimiento, usu.IDUsuarioRol, usu.Telefono, usu.Password, usu.SesionesBonos, usu.Newsletter, usu.FechaBaja)
		if err1 != nil {
			panic(err1.Error())
		}
		usu.ID, err1 = res.LastInsertId()
		log.Println("INSERT: nombre: " + usu.Nombre + " | nif: " + usu.Nif + "| password:" + usu.Password)

	}
	var vrecord model.UsuarioRecord
	vrecord.Result = "OK"
	vrecord.Record = usu
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

//UsuariosDelete da de baja al usuario
func UsuariosDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	usu := r.FormValue("ID")
	delForm, err := db.Prepare("UPDATE usuarios SET fechaBaja=CURDATE() WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(usu)
	if err1 != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error dando de baja al usuario"
		a, _ := json.Marshal(verror)
		w.Write(a)
	}
	log.Println("BAJA")
	defer db.Close()
	var vrecord model.UsuarioRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

// Usuariosgetoptions - Obtener nombres de usuarios para la tabla de autorizados
func Usuariosgetoptions(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT usuarios.id, usuarios.nombre FROM usuarios ORDER BY usuarios.nombre")
	if err != nil {
		panic(err.Error())
	}
	elem := model.Option{}
	vtabla := []model.Option{}
	for selDB.Next() {
		err = selDB.Scan(&elem.Value, &elem.DisplayText)
		if err != nil {
			panic(err.Error())
		}
		vtabla = append(vtabla, elem)
	}

	var vtab model.Options
	vtab.Result = "OK"
	vtab.Options = vtabla
	// create json response from struct
	a, err := json.Marshal(vtab)
	// Visualza
	s := string(a)
	fmt.Println(s)
	w.Write(a)
	defer db.Close()
}

// GetUsuario Get user to login
func GetUsuario(email string) (bool, int, string, string, int) {

	var tipo, userid int
	var passwd, username string
	var found bool
	db := database.DbConn()
	selDB, err := db.Query("SELECT ID,nombre,password,idusuariorol FROM usuarios where email = ?", email)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		panic(err.Error())
	}
	for selDB.Next() {
		found = true
		err = selDB.Scan(&userid, &username, &passwd, &tipo)
	}
	defer db.Close()
	if found {
		return found, userid, username, passwd, tipo
	} else {
		return false, -1, "", "", 1
	}
}
