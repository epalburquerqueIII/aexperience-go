package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../model"
	"../model/database"
)

// UsuarioRoles Pantalla de tratamiento de usuario
func UsuarioRoles(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "usuariosRoles", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// UsuarioRolesList - json con los datos de clientes
func UsuarioRolesList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT id, nombre FROM usuariosRoles " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error. Buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	usuR := model.TusuariosRoles{}
	res := []model.TusuariosRoles{}
	for selDB.Next() {

		err = selDB.Scan(&usuR.ID, &usuR.Nombre)
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error. Cargando registros de roles de usuario"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, usuR)
		i++
	}

	var vrecords model.UsuariosRolesRecords
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

// UsuarioRolesCreate - Crear un Rol Usuario
func UsuarioRolesCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	usuR := model.TusuariosRoles{}
	if r.Method == "POST" {
		usuR.Nombre = r.FormValue("Nombre")
		insForm, err := db.Prepare("INSERT INTO usuariosRoles(nombre) VALUES(?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error. Insertando el rol de usuario"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(usuR.Nombre)
		if err1 != nil {
			panic(err1.Error())
		}
		usuR.ID, err1 = res.LastInsertId()
		log.Println("INSERT: nombre: " + usuR.Nombre)

	}
	var vrecord model.UsuariosRolesRecord
	vrecord.Result = "OK"
	vrecord.Record = usuR
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// UsuarioRolesUpdate Actualiza el rol de usuario
func UsuarioRolesUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	usuR := model.TusuariosRoles{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("Id"))
		usuR.ID = int64(i)
		usuR.Nombre = r.FormValue("Nombre")
		insForm, err := db.Prepare("UPDATE usuariosRoles SET nombre=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error. Actualizando la base de datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(usuR.Nombre, usuR.ID)
		log.Println("UPDATE: nombre: " + usuR.Nombre)
	}
	defer db.Close()
	var vrecord model.UsuariosRolesRecord
	vrecord.Result = "OK"
	vrecord.Record = usuR
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

// UsuarioRolesDelete Borra rol de usuario de la DB
func UsuarioRolesDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	usuR := r.FormValue("Id")
	delForm, err := db.Prepare("DELETE FROM usuariosRoles WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	_, err1 := delForm.Exec(usuR)
	if err1 != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error. Borrando usuario"
		a, _ := json.Marshal(verror)
		w.Write(a)
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.UsuariosRolesRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	http.Redirect(w, r, "/", 301)
}

/* UsuariogetoptionsRoles Roles de usuario
func UsuariogetoptionsRoles(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT usuariosRoles.id, usuariosRoles.nombre from usuariosRoles Order by usuariosRoles.id")
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
}*/
