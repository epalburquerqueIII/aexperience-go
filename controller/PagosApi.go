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

//Pagos Pantalla de tratamiento de Pagos
func Pagos(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)
	error := tmpl.ExecuteTemplate(w, "pagos", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// PagosList - json con los datos de los pagos
func PagosList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT pagos.id, reservas.id, pagos.fechaPago, tipospago.id, numeroTarjeta FROM pagos LEFT OUTER JOIN reservas ON (idReserva = reservas.id) LEFT OUTER JOIN tiposPago ON (idTipopago = tiposPago.id)" + jtsort)
	if err != nil {
		util.ErrorApi(err.Error(), w, "Error en Select ")
	}
	pag := model.Tpago{}
	res := []model.Tpago{}
	for selDB.Next() {

		err = selDB.Scan(&pag.Id, &pag.IdReserva, &pag.FechaPago, &pag.IdTipopago, &pag.NumeroTarjeta)
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando el registros de los Pagos")
		}
		res = append(res, pag)
		i++
	}

	var vrecords model.PagoRecords
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
func PagosCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	pag := model.Tpago{}
	if r.Method == "POST" {
		pag.IdReserva, _ = strconv.Atoi(r.FormValue("IdReserva"))
		pag.FechaPago = r.FormValue("FechaPago")
		pag.IdTipopago, _ = strconv.Atoi(r.FormValue("IdTipopago"))
		pag.NumeroTarjeta = r.FormValue("NumeroTarjeta")
		insForm, err := db.Prepare("INSERT INTO pagos(idReserva, fechaPago, idTipopago, numeroTarjeta) VALUES(?,CURDATE(),?,?)")
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Insertando Pago")
		}
		res, err1 := insForm.Exec(pag.IdReserva, pag.IdTipopago, pag.NumeroTarjeta)
		if err1 != nil {
			panic(err1.Error())
		}
		pag.Id, err1 = res.LastInsertId()
		log.Printf("INSERT: fechaPago: %s | idTipopago:  %d\n", pag.FechaPago, pag.IdTipopago)

	}
	var vrecord model.PagoRecord
	vrecord.Result = "OK"
	vrecord.Record = pag
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()

}

// PagosUpdate Actualiza los pagos
func PagosUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	pag := model.Tpago{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("Id"))
		pag.Id = int64(i)
		pag.IdReserva, _ = strconv.Atoi(r.FormValue("IdReserva"))
		pag.FechaPago = util.DateSql(r.FormValue("FechaPago"))
		pag.IdTipopago, _ = strconv.Atoi(r.FormValue("IdTipopago"))
		pag.NumeroTarjeta = r.FormValue("NumeroTarjeta")
		insForm, err := db.Prepare("UPDATE pagos SET idReserva=?, fechaPago=?, idTipopago=?, numeroTarjeta =? WHERE id=?")
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Actualizando Base de Datos")
		}

		insForm.Exec(pag.IdReserva, pag.FechaPago, pag.IdTipopago, pag.NumeroTarjeta, pag.Id)
		log.Printf("UPDATE: fechaPago: %s | idTipopago:  %d\n", pag.FechaPago, pag.IdTipopago)
	}
	defer db.Close()
	var vrecord model.PagoRecord
	vrecord.Result = "OK"
	vrecord.Record = pag
	a, _ := json.Marshal(vrecord)
	w.Write(a)

}

//PagosDelete Borra pagos de la DB
func PagosDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	pag := r.FormValue("Id")
	delForm, err := db.Prepare("DELETE FROM pagos WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(pag)
	if err1 != nil {
		util.ErrorApi(err.Error(), w, "Error borrando pago")
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.PagoRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

}
