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

// Autorizados Pantalla de tratamiento de Autorizados
func Autorizados(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "autorizado", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// AutorizadoList - json con los datos de clientes
func AutorizadoList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT autorizados.id, autorizados.idUsuario, nombreAutorizado, autorizados.nif FROM autorizados " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	auto := model.Tautorizado{}
	res := []model.Tautorizado{}
	for selDB.Next() {

		err = selDB.Scan(&auto.IDUsuario, &auto.NombreAutorizado, &auto.Nif)
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Cargando registros de Autorizados"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, auto)
		i++
	}

	var vrecords model.AutorizadoRecords
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

// AutorizadoCreate Crear un Autorizado
func AutorizadoCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	auto := model.Tautorizado{}
	if r.Method == "POST" {
		auto.IDUsuario, _ = strconv.Atoi(r.FormValue("idUsuario"))
		auto.NombreAutorizado = r.FormValue("nombreAutorizado")
		auto.Nif = r.FormValue("nif")
		insForm, err := db.Prepare("INSERT INTO autorizados(idUsuario, nombreAutorizado, nif) VALUES(?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Autorizado"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(auto.IDUsuario, auto.NombreAutorizado, auto.Nif)
		if err1 != nil {
			panic(err1.Error())
		}
		auto.ID, err1 = res.LastInsertId()
		log.Println("INSERT: nombreAutorizado: " + auto.NombreAutorizado + " | nif: " + auto.Nif)

	}
	var vrecord model.AutorizadoRecord
	vrecord.Result = "OK"
	vrecord.Record = auto
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// AutorizadoUpdate Actualiza el Autorizado
func AutorizadoUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	auto := model.Tautorizado{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("ID"))
		auto.ID = int64(i)
		auto.IDUsuario, _ = strconv.Atoi(r.FormValue("idUsuario"))
		auto.NombreAutorizado = r.FormValue("nombreAutorizado")
		auto.Nif = r.FormValue("nif")
		insForm, err := db.Prepare("UPDATE autorizados SET nombreAutorizado=?, autorizados.nif=? WHERE autorizados.id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(auto.IDUsuario, auto.NombreAutorizado, auto.Nif, auto.ID)
		log.Println("UPDATE: nombreAutorizado: " + auto.NombreAutorizado + " | nif: " + auto.Nif)
	}
	defer db.Close()
	var vrecord model.AutorizadoRecord
	vrecord.Result = "OK"
	vrecord.Record = auto
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

//AutorizadoDelete Borra Autorizado de la DB
func AutorizadoDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	auto := r.FormValue("ID")
	delForm, err := db.Prepare("DELETE FROM autorizados WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(auto)
	if err1 != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error Borrando Autorizado"
		a, _ := json.Marshal(verror)
		w.Write(a)
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.AutorizadoRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	// 	// 	http.Redirect(w, r, "/", 301)
}

// //AutorizadoBaja da de baja al Autorizado
// func AutorizadoBaja(w http.ResponseWriter, r *http.Request) {
// 	db := database.DbConn()
// 	usu := r.FormValue("ID")
// 	delForm, err := db.Prepare("UPDATE usuarios SET fechaBaja=CURDATE() WHERE id=?")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	_, err1 := delForm.Exec(usu)
// 	if err1 != nil {
// 		var verror model.Resulterror
// 		verror.Result = "ERROR"
// 		verror.Error = "Error dando de baja al usuario"
// 		a, _ := json.Marshal(verror)
// 		w.Write(a)
// 	}
// 	log.Println("BAJA")
// 	defer db.Close()
// 	var vrecord model.UsuarioRecord
// 	vrecord.Result = "OK"
// 	a, _ := json.Marshal(vrecord)
// 	w.Write(a)

// 	// 	// 	http.Redirect(w, r, "/", 301)
// }

// AutorizadogetNombreUsuario - Obtener nombres de usuarios para la tabla de autorizados
func AutorizadogetNombreUsuario(w http.ResponseWriter, r *http.Request) {

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
