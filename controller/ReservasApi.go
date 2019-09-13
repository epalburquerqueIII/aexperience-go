package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"../util"
	"../model"
	"../model/database"
	"../util"
)

// Reservas Pantalla de tratamiento de Reservas
func Reservas(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "reservas", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// ReservasList - json con los datos de clientes
func ReservasList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT reservas.id, reservas.fecha, reservas.fechaPago, reservas.hora, usuarios.id, espacios.id, autorizados.id FROM reservas LEFT OUTER JOIN usuarios ON (usuarios.id = reservas.idUsuario) LEFT OUTER JOIN espacios ON (espacios.id = reservas.idEspacio) LEFT OUTER JOIN autorizados ON (autorizados.id = reservas.idAutorizado) " + jtsort)
	if err != nil {
		util.ErrorApi(err.Error(),w,"Error en Select ")
	}
	reser := model.Treservas{}
	res := []model.Treservas{}
	for selDB.Next() {

		err = selDB.Scan(&reser.Id, &reser.Fecha, &reser.FechaPago, &reser.Hora, &reser.IdUsuario, &reser.IdEspacio, &reser.IdAutorizado)
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Cargando registros de Reservas"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, reser)
		i++
	}

	var vrecords model.ReservasRecords
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

// ReservasCreate Crear un Usuario
func ReservasCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	reser := model.Treservas{}
	if r.Method == "POST" {
		reser.Fecha = util.DateSql(r.FormValue("Fecha"))
		reser.FechaPago = r.FormValue("FechaPago")
		reser.Hora, _ = strconv.Atoi(r.FormValue("Hora"))
		reser.IdUsuario, _ = strconv.Atoi(r.FormValue("IdUsuario"))
		reser.IdEspacio, _ = strconv.Atoi(r.FormValue("IdEspacio"))
		reser.IdAutorizado, _ = strconv.Atoi(r.FormValue("IdAutorizado"))

		insForm, err := db.Prepare("INSERT INTO reservas(fecha, fechaPago, hora, idUsuario, idEspacio, idAutorizado) VALUES(?,CURDATE(),?,?,?,?)")

		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Pago"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		res, err1 := insForm.Exec(reser.Fecha, reser.Hora, reser.IdUsuario, reser.IdEspacio, reser.IdAutorizado)

		if err1 != nil {
			panic(err1.Error())
		}
		reser.Id, err1 = res.LastInsertId()
		log.Printf("INSERT: fecha: %s | fechaPago: %s | hora:  %d\n ", reser.Fecha, reser.FechaPago, reser.Hora)

	}
	var vrecord model.ReservasRecord
	vrecord.Result = "OK"
	vrecord.Record = reser
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// ReservasUpdate Actualiza las reservas
func ReservasUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	reser := model.Treservas{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("Id"))
		reser.Id = int64(i)
		reser.Fecha = util.DateSql(r.FormValue("Fecha"))


		reser.FechaPago = util.DateSql(r.FormValue("FechaPago"))
		reser.Hora, _ = strconv.Atoi(r.FormValue("Hora"))
		reser.IdUsuario, _ = strconv.Atoi(r.FormValue("IdUsuario"))
		reser.IdEspacio, _ = strconv.Atoi(r.FormValue("IdEspacio"))
		reser.IdAutorizado, _ = strconv.Atoi(r.FormValue("IdAutorizado"))

		insForm, err := db.Prepare("UPDATE reservas SET fecha=?, fechaPago=CURDATE(), hora=?, idUsuario=?, idEspacio =?, idAutorizado=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(reser.Fecha, reser.Hora, reser.IdUsuario, reser.IdEspacio, reser.IdAutorizado, reser.Id)
		log.Printf("UPDATE: fecha: %s  | fechaPago: %s | hora:  %d\n", reser.Fecha, reser.FechaPago, reser.Hora)

	}
	defer db.Close()
	var vrecord model.ReservasRecord
	vrecord.Result = "OK"
	vrecord.Record = reser
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

// ReservasDelete Borra reservas de la DB
func ReservasDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	reser := r.FormValue("Id")
	delForm, err := db.Prepare("DELETE FROM reservas WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(reser)
	if err1 != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error Borrando reservas"
		a, _ := json.Marshal(verror)
		w.Write(a)
	}
	log.Println("Borrar")
	defer db.Close()
	var vrecord model.ReservasRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	// 	// 	// 	http.Redirect(w, r, "/", 301)
}

//ReservasgetoptionsRoles Roles de usuario
func ReservasgetoptionsRoles(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT usuarios.id, usuarios.nombre from usuarios Order by usuarios.id")
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

//ReservasgetoptionsRoles tabala de espacios
func ReservasgetoptionsEspacios(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT espacios.id, espacios.descripcion from espacios Order by espacios.id")
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

//Reservasgetoptions tabla de autorizados
func ReservasgetoptionsAutorizado(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT autorizados.id, autorizados.nombreAutorizado from autorizados Order by autorizados.id")
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
