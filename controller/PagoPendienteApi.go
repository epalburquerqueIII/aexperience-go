package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../model"
	"../model/database"
	"../util"
)

//PagosPendientes Pantalla de tratamiento de Pagos
func PagosPendientes(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)
	error := tmpl.ExecuteTemplate(w, "pagosPendientes", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// PagosPendientesList - json con los datos de los pagos
func PagosPendientesList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT pagospendientes.id, reservas.id,reservas.fecha, pagospendientes.fechaPago, tipospago.id, tipospago.nombre, pagospendientes.numeroTarjeta, pagospendientes.importe FROM pagosPendientes LEFT OUTER JOIN reservas ON (idReserva = reservas.id) LEFT OUTER JOIN tiposPago ON (idTipopago = tiposPago.id)" + jtsort)
	if err != nil {
		util.ErrorApi(err.Error(), w, "Error en Select ")
	}
	pagopend := model.TpagoPendiente{}
	res := []model.TpagoPendiente{}
	for selDB.Next() {

		err = selDB.Scan(&pagopend.Id, &pagopend.IdReserva, &pagopend.ReservaNombre, &pagopend.FechaPago, &pagopend.IdTipopago, &pagopend.TipopagoNombre, &pagopend.NumeroTarjeta, &pagopend.Importe)
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando el registros de los Pagos")
		}
		res = append(res, pagopend)
		i++
	}

	var vrecords model.PagoPendienteRecords
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

// Pagospendientesgetoptions - Obtener pagos de pagos pendientes para la tabla de pagos
func Pagospendientesgetoptions(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT pagosPendientes.id, pagosPendientes.idreserva, pagosPendientes.fechapago, pagosPendientes.idtipopago, pagosPendientes.numerotarjeta, pagospendientes.importe from pagosPendientes Order by pagosPendientes.id")
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
