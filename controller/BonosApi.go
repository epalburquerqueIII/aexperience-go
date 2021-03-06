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

// GetBonos -
func GetBonos() []model.Tbono {
	db := database.DbConn()
	selDB, err := db.Query("SELECT bonos.id, precio, sesiones FROM bonos")
	if err != nil {
		log.Println("Error en Select")
	}
	bon := model.Tbono{}
	res := []model.Tbono{}
	for selDB.Next() {
		err = selDB.Scan(&bon.ID, &bon.Precio, &bon.Sesiones)
		if err != nil {
			log.Println("Error al insertar datos")
		}
		res = append(res, bon)

	}
	return res
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
		i, _ := strconv.Atoi(r.FormValue("ID"))
		bon.ID = int64(i)
		bon.Precio, _ = strconv.ParseFloat(r.FormValue("Precio"), 64)
		bon.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		insForm, err := db.Prepare("INSERT INTO bonos(id, precio, sesiones) VALUES(?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Bono"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(bon.ID, bon.Precio, bon.Sesiones)
		if err1 != nil {
			panic(err1.Error())
		}
		bon.ID, err1 = res.LastInsertId()
		log.Printf("INSERT: precio: %d | sesiones: %.2f\n", bon.Precio, bon.Sesiones)

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
		bon.Precio, _ = strconv.ParseFloat(r.FormValue("Precio"), 64)
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
		log.Printf("UPDATE: id: %d | precio: %.2f\n", bon.ID, bon.Precio)

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
func BonoDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	usu := r.FormValue("ID")
	delForm, err := db.Prepare("DELETE FROM bonos WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(usu)
	if err1 != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error Borrando bono"
		a, _ := json.Marshal(verror)
		w.Write(a)
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.BonoRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	http.Redirect(w, r, "/", 301)
}
