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

// TiposeventosList - json con los datos de clientes
func TiposeventosList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT tiposevento.ID  , tiposevento.Nombre FROM tiposevento " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error no hay eventos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	tipo := model.Ttiposevento{}
	res := []model.Ttiposevento{}
	for selDB.Next() {

		err = selDB.Scan(&tipo.ID, &tipo.Nombre)
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error buscando eventos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, tipo)
		i++
	}

	var vrecords model.TiposeventoRecords
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

// TiposeventosCreate Crear un Autorizado
func TiposeventosCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	tipo := model.Ttiposevento{}
	if r.Method == "POST" {
		tipo.Nombre = r.FormValue("Nombre")
		insForm, err := db.Prepare("INSERT INTO tiposevento(id, nombre) VALUES(?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Eventos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(tipo.ID, tipo.Nombre)
		if err1 != nil {
			panic(err1.Error())
		}
		tipo.ID, err1 = res.LastInsertId()
		log.Println("INSERT: nombre: " + tipo.Nombre)

	}
	var vrecord model.TiposeventoRecord
	vrecord.Result = "OK"
	vrecord.Record = tipo
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// TiposeventosUpdate Actualiza el Autorizado
func TiposeventosUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	tipo := model.Ttiposevento{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("ID"))
		tipo.ID = int64(i)
		tipo.Nombre = r.FormValue("Nombre")
		insForm, err := db.Prepare("UPDATE tiposevento SET nombre=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando los Eventos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(tipo.Nombre, tipo.ID)
		log.Println("UPDATE: nombre: " + tipo.Nombre)
	}
	defer db.Close()
	var vrecord model.TiposeventoRecord
	vrecord.Result = "OK"
	vrecord.Record = tipo
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

//TiposeventosDelete Borra Autorizado de la DB
func TiposeventosDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	tipo := r.FormValue("ID")
	delForm, err := db.Prepare("DELETE FROM tiposevento WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(tipo)
	if err1 != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error Borrando Eventos"
		a, _ := json.Marshal(verror)
		w.Write(a)
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.TiposeventoRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	// 	// 	http.Redirect(w, r, "/", 301)
}

// TiposeventosgetOptions nombre tipo evento
func TiposeventosgetOptions(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT id, nombre from tiposevento Order by id")
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
