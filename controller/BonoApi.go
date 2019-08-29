//TODO dar de alta bonos dados de baja

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

var tmplbon = template.Must(template.ParseGlob("views/*.html"))

// BonoList - json con los datos de clientes
func BonoList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT precio, sesiones FROM bonos " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	bon := model.Tbono{}
	res := []model.Tbono{}
	for selDB.Next() {

		err = selDB.Scan(&bon.Precio, &bon.Sesiones)

		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Cargando registros de Bonos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, bon)
		i++
	}

	var vrecords model.BonoRecords
	vrecords.Result = "OK"
	vrecords.TotalRecordCount = i
	vrecords.Records = res
	// create json response from struct
	a, err := json.Marshal(vrecords)
	// Visualiza
	s := string(a)
	fmt.Println(s)
	w.Write(a)
	defer db.Close()
}

// BonoCreate Crear un Bono
// func BonoCreate(w http.ResponseWriter, r *http.Request) {

// 	db := database.DbConn()
// 	bon := model.Tbono{}
// 	if r.Method == "POST" {
// 		bon.Precio, _ = strconv.Atoi(r.FormValue("Precio"))
// 		bon.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
// 		insForm, err := db.Prepare("INSERT INTO bonos(precio, sesiones) VALUES(?,?)")
// 		if err != nil {
// 			var verror model.Resulterror
// 			verror.Result = "ERROR"
// 			verror.Error = "Error Insertando Bono"
// 			a, _ := json.Marshal(verror)
// 			w.Write(a)
// 			panic(err.Error())
// 		}
// 		res, err1 := insForm.Exec(bon.Precio, bon.Sesiones)
// 		if err1 != nil {
// 			panic(err1.Error())
// 		}
// 		bon.Precio, err1 = res.LastInsertId()
// 		log.Println("INSERT: precio: " + bon.Precio + " | sesiones: " + bon.Sesiones)

// 	}
// 	var vrecord model.BonoRecord
// 	vrecord.Result = "OK"
// 	vrecord.Record = bon
// 	a, _ := json.Marshal(vrecord)
// 	s := string(a)
// 	fmt.Println(s)

// 	w.Write(a)

// 	defer db.Close()
// 	//	http.Redirect(w, r, "/", 301)
// }

// BonoUpdate Actualiza el bono
func BonoUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	bon := model.Tbono{}
	if r.Method == "POST" {
		bon.Precio, _ = strconv.Atoi(r.FormValue("Precio"))
		bon.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		insForm, err := db.Prepare("UPDATE bonos SET sesiones=? WHERE precio=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(bon.Precio, bon.Sesiones)
		log.Println("UPDATE: precio: %d   | sesiones: %d\n", bon.Precio, bon.Sesiones)

	}
	defer db.Close()
	var vrecord model.BonoRecord
	vrecord.Result = "OK"
	vrecord.Record = bon
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

// BonoDelete Borra bono de la DB
// func BonoDelete(w http.ResponseWriter, r *http.Request) {
// 	db := database.DbConn()
// 	bon := r.FormValue("Precio")
// 	delForm, err := db.Prepare("DELETE FROM bonos WHERE id=?")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	_, err1 := delForm.Exec(bon)
// 	if err1 != nil {
// 		var verror model.Resulterror
// 		verror.Result = "ERROR"
// 		verror.Error = "Error Borrando bono"
// 		a, _ := json.Marshal(verror)
// 		w.Write(a)
// 	}
// 	log.Println("DELETE")
// 	defer db.Close()
// 	var vrecord model.BonoRecord
// 	vrecord.Result = "OK"
// 	a, _ := json.Marshal(vrecord)
// 	w.Write(a)

// 	// 	// 	http.Redirect(w, r, "/", 301)
// }

// 	// 	http.Redirect(w, r, "/", 301)

// BonogetoptionsRoles Roles de bono
// func BonogetoptionsRoles(w http.ResponseWriter, r *http.Request) {

// 	db := database.DbConn()
// 	selDB, err := db.Query("SELECT bonos_roles.precio, bonos_roles.sesiones from bonos_roles Order by bonos_roles.precio")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	elem := model.Option{}
// 	vtabla := []model.Option{}
// 	for selDB.Next() {
// 		err = selDB.Scan(&elem.Value, &elem.DisplayText)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		vtabla = append(vtabla, elem)
// 	}

// 	var vtab model.Options
// 	vtab.Result = "OK"
// 	vtab.Options = vtabla
// 	// create json response from struct
// 	a, err := json.Marshal(vtab)
// 	// Visualza
// 	s := string(a)
// 	fmt.Println(s)
// 	w.Write(a)
// 	defer db.Close()
// }
