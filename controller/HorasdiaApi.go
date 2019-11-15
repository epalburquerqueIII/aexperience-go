package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"../model"
	"../model/database"
	"../util"
)

func MovilHorasReservables(w http.ResponseWriter, r *http.Request) {

	var fechabusqueda, espacio string
	if r.FormValue("dia") == "1" {
		fechabusqueda = time.Now().Format("2006-01-02")
	} else {
		fechabusqueda = "0"
	}
	if r.FormValue("espacio") == "1" {
		espacio = "5"
	} else {
		espacio = "1"
	}
	horasreservada := HorasReservables(fechabusqueda, espacio)

	var vrecords model.HorariodiaRecords
	vrecords.Result = "OK"
	vrecords.TotalRecordCount = 1
	vrecords.Records = horasreservada
	// create json response from struct
	a, _ := json.Marshal(vrecords)
	// Visualza
	s := string(a)
	fmt.Println(s)
	w.Write(a)
}

func HorasReservables(fechabusqueda string, espacio string) []model.THorasdia {

	db := database.DbConn()
	selDB, err := db.Query("SELECT horarios.id, horarios.idespacio,horarios.hora, " +
		"IF(isnull(reservas.hora),0,1) as reservado from horarios " +
		"left outer join reservas on " +
		"reservas.fecha = '" + fechabusqueda + "' and " +
		"reservas.hora = horarios.hora and " +
		"horarios.idespacio = reservas.idespacio " +
		"where horarios.fechaInicio >= '2019-10-01' " +
		"and horarios.fechaFin <= '2019-12-31' and " +
		"horarios.idespacio = '" + espacio + "' " +
		" order by hora")

	horareservada := model.THorasdia{}
	horasreservada := []model.THorasdia{}

	ID := 0
	for selDB.Next() {

		err = selDB.Scan(&ID, &horareservada.IDEspacio, &horareservada.Hora, &horareservada.Reservado)

		if err != nil {
			util.ErrorApi(err.Error(), nil, "Error Cargando datos de horas dia")
		}
		horasreservada = append(horasreservada, horareservada)
	}

	return horasreservada

}

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
		reser.Fecha = r.FormValue("fechabusqueda")

		for i := 0; i <= 8; i++ {
			a := "ch" + string(i)
			fmt.Printf(a)
			hora = r.FormValue("ch" + strconv.Itoa(i))
			if hora != "0" {
				guardado = 1
			}
			if guardado == 1 {
				guardado = 0
				horas := strings.TrimSpace(hora[:2])
				reser.Hora, _ = strconv.Atoi(horas)
				reser.IdEspacio, _ = strconv.Atoi(r.FormValue("espacio"))
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
