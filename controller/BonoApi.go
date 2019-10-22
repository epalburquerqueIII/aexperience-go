package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../model"
	"../model/database"
	"../util"
)

//Bonos Pantalla de tratamiento de Bonos
func Bonos(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)
	error := tmpl.ExecuteTemplate(w, "bonos", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// BonoList - json con los datos de clientes
func BonoList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT bonos.id, precio, sesiones FROM bonos " + jtsort)
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

		err = selDB.Scan(&bon.ID, &bon.Precio, &bon.Sesiones)

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
func BonoCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	bon := model.Tbono{}
	if r.Method == "POST" {
		bon.Precio, _ = strconv.Atoi(r.FormValue("Precio"))
		bon.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		insForm, err := db.Prepare("INSERT INTO bonos(precio, sesiones) VALUES(?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Bono"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(bon.Precio, bon.Sesiones)
		if err1 != nil {
			panic(err1.Error())
		}
		bon.ID, err1 = res.LastInsertId()
		log.Printf("INSERT: precio: %d | sesiones: %d\n", bon.Precio, bon.Sesiones)

	}
	var vrecord model.BonoRecord
	vrecord.Result = "OK"
	vrecord.Record = bon
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// BonoUpdate Actualiza el bono
func BonoUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	bon := model.Tbono{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("ID"))
		bon.ID = int64(i)
		bon.Precio, _ = strconv.Atoi(r.FormValue("Precio"))
		bon.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		insForm, err := db.Prepare("UPDATE bonos SET precio =?, sesiones=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(bon.Precio, bon.Sesiones, bon.ID)
		log.Printf("UPDATE: id: %d | precio: %d\n", bon.ID, bon.Precio)

	}
	defer db.Close()
	var vrecord model.BonoRecord
	vrecord.Result = "OK"
	vrecord.Record = bon
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}
