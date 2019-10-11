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

//PagoPendientes Pantalla de tratamiento de Pagos
func PagoPendientes(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "pagosPendientes", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// PagoPendientesList - json con los datos de los pagos
func PagoPendientesList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT pagospendientes.id, reservas.id, pagospendientes.fechaPago, tipospago.id, pagospendientes.numeroTarjeta FROM pagosPendientes LEFT OUTER JOIN reservas ON (idReserva = reservas.id) LEFT OUTER JOIN tiposPago ON (idTipopago = tiposPago.id)" + jtsort)
	if err != nil {
		util.ErrorApi(err.Error(), w, "Error en Select ")
	}
	pagopend := model.TpagosPendientes{}
	res := []model.TpagosPendientes{}
	for selDB.Next() {

		err = selDB.Scan(&pagopend.Id, &pagopend.IdReserva, &pagopend.FechaPago, &pagopend.IdTipopago, &pagopend.NumeroTarjeta)
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando el registros de los Pagos")
		}
		res = append(res, pagopend)
		i++
	}

	var vrecords model.PagosPendientesRecords
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

// PagosCreate Crear un Pago
func PagoPendientesCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	pagopend := model.TpagosPendientes{}
	if r.Method == "POST" {
		pagopend.IdReserva, _ = strconv.Atoi(r.FormValue("IdReserva"))
		pagopend.FechaPago = r.FormValue("FechaPago")
		pagopend.IdTipopago, _ = strconv.Atoi(r.FormValue("IdTipopago"))
		pagopend.NumeroTarjeta = r.FormValue("NumeroTarjeta")
		insForm, err := db.Prepare("INSERT INTO pagosPendientes(idReserva, fechaPago, idTipopago, numeroTarjeta) VALUES(?,CURDATE(),?,?)")
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Insertando Pago")
		}
		res, err1 := insForm.Exec(pagopend.IdReserva, pagopend.IdTipopago, pagopend.NumeroTarjeta)
		if err1 != nil {
			panic(err1.Error())
		}
		pagopend.Id, err1 = res.LastInsertId()
		log.Printf("INSERT: fechaPago: %s | idTipopago:  %d\n", pagopend.FechaPago, pagopend.IdTipopago)

	}
	var vrecord model.PagosPendientesRecord
	vrecord.Result = "OK"
	vrecord.Record = pagopend
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()

}

// PagoPendienteUpdate Actualiza los pagos pendientes
func PagoPendientesUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	pagopend := model.TpagosPendientes{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("Id"))
		pagopend.Id = int64(i)
		pagopend.IdReserva, _ = strconv.Atoi(r.FormValue("IdReserva"))
		pagopend.FechaPago = util.DateSql(r.FormValue("FechaPago"))
		pagopend.IdTipopago, _ = strconv.Atoi(r.FormValue("IdTipopago"))
		pagopend.NumeroTarjeta = r.FormValue("NumeroTarjeta")
		insForm, err := db.Prepare("UPDATE pagosPendientes SET idReserva=?, fechaPago=?, idTipopago=?, numeroTarjeta =? WHERE id=?")
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Actualizando Base de Datos")
		}

		insForm.Exec(pagopend.IdReserva, pagopend.FechaPago, pagopend.IdTipopago, pagopend.NumeroTarjeta, pagopend.Id)
		log.Printf("UPDATE: fechaPago: %s | idTipopago:  %d\n", pagopend.FechaPago, pagopend.IdTipopago)
	}
	defer db.Close()
	var vrecord model.PagosPendientesRecord
	vrecord.Result = "OK"
	vrecord.Record = pagopend
	a, _ := json.Marshal(vrecord)
	w.Write(a)

}

//PagosDelete Borra pagos de la DB
func PagoPendientesDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	pagopend := r.FormValue("Id")
	delForm, err := db.Prepare("DELETE FROM pagosPendientes WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(pagopend)
	if err1 != nil {
		util.ErrorApi(err.Error(), w, "Error borrando pago")
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.PagosPendientesRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

}
