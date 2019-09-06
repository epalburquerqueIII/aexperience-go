//TODO dar de alta los usuarios que están de baja

package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"../model"
	"../model/database"
)

var tmpl = template.Must(template.ParseGlob("views/*.html"))

// Menus Pantalla de tratamiento de menus
func Menu(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "menus", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

//MenusList
func MenusList(w http.ResponseWriter, r *http.Request) {

	var i int
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT menus.id, parent_id, orden, titulo, icono, url, handleFunc FROM menus " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	men := model.Tmenu{}
	res := []model.Tmenu{}
	for selDB.Next() {

		err = selDB.Scan(&men.ID, &men.ParentID, &men.Orden, &men.Titulo, &men.Icono, &men.Url, &men.HandleFunc)
		/*//Si no hay fecha de baja, este campo aparece como activo
		if usu.FechaBaja == "0000-00-00" {
			usu.FechaBaja = "Activo"
		} else {
			//Formato de fecha en español cuando está de baja
			t, _ := time.Parse("2006-01-02", usu.FechaBaja)
			usu.FechaBaja = t.Format("02-01-2006")

		}*/
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Cargando registros de Menús"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, men)
		i++
	}

	var vrecords model.MenuRecords
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

// MenusCreate - Crear un Menus
func MenusCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	men := model.Tmenu{}
	if r.Method == "POST" {
		men.ParentID, _ = strconv.Atoi(r.FormValue("Parent_id"))
		men.Orden = r.FormValue("Orden")
		men.Titulo = r.FormValue("Titulo")
		men.Icono = r.FormValue("Icono")
		men.Url = r.FormValue("Url")
		men.HandleFunc = r.FormValue("HandleFunc")

		insForm, err := db.Prepare("INSERT INTO menus(parent_id, orden, titulo, icono, url, handleFunc) VALUES(?,?,?,?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Menus"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(men.ParentID, men.Orden, men.Titulo, men.Icono, men.Url, men.HandleFunc)
		if err1 != nil {
			panic(err1.Error())
		}
		men.ID, err1 = res.LastInsertId()
		log.Println("INSERT: parent id: " + men.ParentID)

	}
	var vrecord model.MenusRecords
	vrecord.Result = "OK"
	vrecord.Record = men
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// MenusUpdate Actualiza el menus
func MenusUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	men := model.Tmenu{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("ID"))
		men.ID = int64(i)
		men.ParentID, _ = strconv.Atoi(r.FormValue("Parent_id"))
		men.Orden = r.FormValue("Orden")
		men.Titulo = r.FormValue("Titulo")
		men.Icono = r.FormValue("Icono")
		men.Url = r.FormValue("Url")
		men.HandleFunc = r.FormValue("handleFunc")
		insForm, err := db.Prepare("UPDATE menus SET parent_id=?, orden=?, titulo=?, icono=?, url=?, handleFunc=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(men.ParentID, men.Orden, men.Titulo, men.Icono, men.Url, men.HandleFunc, men.ID)
		log.Println("UPDATE: nombre: " + men.Parent_id)
	}
	defer db.Close()
	var vrecord model.MenuRecord
	vrecord.Result = "OK"
	vrecord.Record = men
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

// UsuarioDelete Borra usuario de la DB
// func UsuarioDelete(w http.ResponseWriter, r *http.Request) {
// 	db := database.DbConn()
// 	usu := r.FormValue("ID")
// 	delForm, err := db.Prepare("DELETE FROM usuarios WHERE id=?")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	_, err1 := delForm.Exec(usu)
// 	if err1 != nil {
// 		var verror model.Resulterror
// 		verror.Result = "ERROR"
// 		verror.Error = "Error Borrando usuario"
// 		a, _ := json.Marshal(verror)
// 		w.Write(a)
// 	}
// 	log.Println("DELETE")
// 	defer db.Close()
// 	var vrecord model.UsuarioRecord
// 	vrecord.Result = "OK"
// 	a, _ := json.Marshal(vrecord)
// 	w.Write(a)

// 	// 	// 	http.Redirect(w, r, "/", 301)
// }

/*//UsuarioBaja da de baja al usuario
func MenuBaja(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	men := r.FormValue("ID")
	delForm, err := db.Prepare("UPDATE menus SET fechaBaja=CURDATE() WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(men)
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

	// 	// 	http.Redirect(w, r, "/", 301)
}*/

// UsuariogetoptionsRoles Roles de usuario
func UsuariogetoptionsRoles(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT usuarios_roles.id, usuarios_roles.nombre from usuarios_roles Order by usuarios_roles.id")
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
