package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"../model"
	"../model/database"
	"../util"
)

type datocarga struct {
	fechabusqueda  string
	espacioposible string
}

var fechafinal string
var espaciofinal string

// ReservaPabellonPista Pantalla de tratamiento de Espacio
func ReservaPabellonPista(w http.ResponseWriter, r *http.Request) {

	error := tmpl.ExecuteTemplate(w, "reservapabellonpista", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// HorasDia  Pantalla de tratamiento de Espacio
/*
func HorasDia(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT horarios.id, horarios.idespacio,horarios.hora, IF(isnull(reservas.hora),0,1) as reservado from horarios left outer join reservas on reservas.fecha = '2019-12-30' and reservas.hora = horarios.hora and horarios.idespacio = reservas.idespacio where horarios.fechaInicio >= '2019-10-01' and horarios.fechaFin <= '2019-12-31' and horarios.idespacio = 1 order by hora")

	hora := model.THorasdia{}
	horas := []model.THorasdia{}

	for selDB.Next() {

		err = selDB.Scan(&hora.ID, &hora.IDEspacio, &hora.Hora, &hora.Reservado)

		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando datos de horas dia")
		}
		horas = append(horas, hora)
	}

	error := tmpl.ExecuteTemplate(w, "horasdia", &horas)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}
*/
func HorasReservables(w http.ResponseWriter, r *http.Request) {
	datos := datocarga{}
	if r.FormValue("dia") == "1" {
		datos.fechabusqueda = time.Now().Format("2006-01-02")
	} else {
		datos.fechabusqueda = "0"
	}
	if r.FormValue("espacio") == "1" {
		datos.espacioposible = "5"
	} else {
		datos.espacioposible = "1"
	}
	fechafinal = datos.fechabusqueda
	espaciofinal = datos.espacioposible
	db := database.DbConn()
	selDB, err := db.Query("SELECT horarios.id, horarios.idespacio,horarios.hora, IF(isnull(reservas.hora),0,1) as reservado from horarios left outer join reservas on reservas.fecha = '" + fechafinal + "' and reservas.hora = horarios.hora and horarios.idespacio = reservas.idespacio where horarios.fechaInicio >= '2019-10-01' and horarios.fechaFin <= '2019-12-31' and horarios.idespacio = '" + espaciofinal + "' order by hora")

	horareservada := model.THorasdia{}
	horasreservada := []model.THorasdia{}

	ID := 0
	for selDB.Next() {

		err = selDB.Scan(&ID, &horareservada.IDEspacio, &horareservada.Hora, &horareservada.Reservado)

		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando datos de horas dia")
		}
		horasreservada = append(horasreservada, horareservada)
	}

	error := tmpl.ExecuteTemplate(w, "horasreservables", &horasreservada)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}

// EspaciosList - json con los datos de Espacio
/*
func HorasDiaList(w http.ResponseWriter, r *http.Request) {

	var i int
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}

	db := database.DbConn()
	selDB, err := db.Query("SELECT horarios.id, horarios.idespacio, horarios.fechaFin,horarios.hora, IF(isnull(reservas.hora), 'No',' Si') as reservado from horarios left outer join reservas on reservas.fecha = '2019-11-11' and reservas.hora = horarios.hora and horarios.idespacio = reservas.idespacio where horarios.fechaInicio <= '2019-10-11' and horarios.fechaFin >= '2019-10-12' and horarios.idespacio = 1 " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	esp := model.THorasdia{}
	res := []model.THorasdia{}
	for selDB.Next() {

		err = selDB.Scan(&esp.ID, &esp.IDEspacio, &esp.Fecha, &esp.Hora, &esp.Reservado)
		//Si no hay fecha de baja, este campo aparece como activo

		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Cargando registros de Espacios"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, esp)
		i++
	}

	var vrecords model.HorariodiaRecords
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
*/
// HorasDiaCreate - json con los datos de Espacio
func HorasDiaCreate(w http.ResponseWriter, r *http.Request) {
	//Falta saltar si y esta reservada y modificar el aspecto
	var hora string
	usuariosimulcro := 1
	rolsimulacro := 1
	sesionessimulacro := 0
	db := database.DbConn()
	reser := model.Treserva{}
	if r.Method == "POST" {
		guardado := 0
		reser.Fecha = fechafinal

		for i := 0; i <= 8; i++ {
			a := "ch" + string(i)
			fmt.Printf(a)
			hora = r.FormValue("ch" + strconv.Itoa(i))
			if hora != "" {
				guardado = 1
			}
			if guardado == 1 {
				guardado = 0
				reser.Hora, _ = strconv.Atoi(hora)
				reser.IdEspacio, _ = strconv.Atoi(espaciofinal)
				reser.IdUsuario = usuariosimulcro
				reser.Sesiones = sesionessimulacro
				reser.IdAutorizado = rolsimulacro
				insForm, err := db.Prepare("INSERT INTO reservas (fecha, Sesiones, hora, idUsuario, idEspacio, idAutorizado) VALUES(?,?,?,?,?,?)")
				if err != nil {
					util.ErrorApi(err.Error(), w, "Error Insertando Pago")
				}
				res, err1 := insForm.Exec(reser.Fecha, reser.Sesiones, reser.Hora, reser.IdUsuario, reser.IdEspacio, reser.IdAutorizado)
				if err1 != nil {
					panic(err1.Error())
					util.ErrorApi(err.Error(), w, "")
				}
				defer db.Close()
				reser.Id, err1 = res.LastInsertId()
				log.Printf("INSERT: fecha: %s | hora:  %d \n ", reser.Fecha, reser.Hora)
			}
		}

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
