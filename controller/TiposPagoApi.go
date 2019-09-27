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

// TiposPago Pantalla de tratamiento de TiposPagos
func TiposPago(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "tiposPago", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// TiposPagoList - json con los datos de clientes
func TiposPagoList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT id, nombre FROM tiposPago " + jtsort)
	if err != nil {
		util.ErrorApi(err.Error(), w, "Error en Select ")
	}
	tip := model.TtiposPago{}
	res := []model.TtiposPago{}
	for selDB.Next() {

		err = selDB.Scan(&tip.Id, &tip.Nombre)
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando el registros de los tipos de pagos")
		}
		res = append(res, tip)
		i++
	}

	var vrecords model.TiposPagoRecords
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

// TiposPagoCreate Crear un Tipo de Pago
func TiposPagoCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	tip := model.TtiposPago{}
	if r.Method == "POST" {
		tip.Nombre = r.FormValue("Nombre")
		insForm, err := db.Prepare("INSERT INTO tiposPago(nombre) VALUES(?)")
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Insertando tipo de pago")
		}
		res, err1 := insForm.Exec(tip.Nombre)
		if err1 != nil {
			panic(err1.Error())
		}
		tip.Id, err1 = res.LastInsertId()
		log.Printf("INSERT: nombre: " + tip.Nombre)

	}
	var vrecord model.TiposPagoRecord
	vrecord.Result = "OK"
	vrecord.Record = tip
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()

	http.Redirect(w, r, "/", 301)
}

// TiposPagoUpdate Actualiza el tipo de pago
func TiposPagoUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	tip := model.TtiposPago{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("Id"))
		tip.Id = int64(i)
		tip.Nombre = r.FormValue("Nombre")

		insForm, err := db.Prepare("UPDATE tiposPago SET nombre=? WHERE id=?")
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Actualizando Base de Datos")
		}

		insForm.Exec(tip.Nombre, tip.Id)
		log.Printf("UPDATE: nombre: " + tip.Nombre)
	}
	defer db.Close()
	var vrecord model.TiposPagoRecord
	vrecord.Result = "OK"
	vrecord.Record = tip
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	http.Redirect(w, r, "/", 301)
}

// TiposPagoDelete Borra tipo de pago de la DB
func TiposPagoDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	tip := r.FormValue("Id")
	delForm, err := db.Prepare("DELETE FROM tiposPago WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(tip)
	if err1 != nil {
		util.ErrorApi(err.Error(), w, "Error borrando tipo de pago")
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.TiposPagoRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	http.Redirect(w, r, "/", 301)
}

// TiposPagogetoptions tipos de pago para tabla pagos
func TiposPagogetoptions(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT  tipospago.id, tipospago.nombre from tipospago Order by tipospago.id")
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
