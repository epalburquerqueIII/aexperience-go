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

// PagosPendientesList - json con los datos de los pagos
func PagosPendientesList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT pagosPendientes.id, pagosPendientes.idreserva, reservas.fecha, pagosPendientes.fechapago, pagosPendientes.idtipopago," +
		"tipospago.nombre, pagosPendientes.numerotarjeta, pagospendientes.importe FROM pagosPendientes " +
		"LEFT OUTER JOIN Reservas on idreserva = Reservas.id " +
		"LEFT OUTER JOIN tiposPago ON (idTipopago = tiposPago.id) " +
		"ORDER BY pagosPendientes.id " + jtsort)
	if err != nil {
		util.ErrorApi(err.Error(), w, "Error en Select ")
	}
	pagopend := model.TpagoPendiente{}
	res := []model.TpagoPendiente{}
	for selDB.Next() {

		err = selDB.Scan(&pagopend.Id, &pagopend.IdReserva, &pagopend.FechaReserva, &pagopend.FechaPago, &pagopend.IdTipopago, &pagopend.TipopagoNombre, &pagopend.NumeroTarjeta, &pagopend.Importe)
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

// Pagospendientesconfirmarpago confirma un pago pendiente y lo pasa a pago
func Pagospendientesconfirmarpago(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	idPagopendiente, _ := strconv.Atoi(r.FormValue("Id"))

	var Sesiones int
	var idUsuario int

	//Obtener las sesiones de los usuarios de las reservas
	selDB, err := db.Query("SELECT Sesiones, idUsuario FROM pagospendientes LEFT OUTER JOIN reservas ON (idReserva = reservas.id) WHERE pagospendientes.id = ? ", idPagopendiente)
	if err != nil {
		panic(err.Error())
	} else {
		selDB.Next()
		err = selDB.Scan(&Sesiones, &idUsuario)
	}

	//Traspasar los registros de pagos pendientes a pagos
	sql := "INSERT INTO pagos (pagos.idReserva, pagos.fechapago, pagos.idTipopago, pagos.numeroTarjeta, pagos.importe) " +
		"(SELECT  pagospendientes.idReserva,  pagospendientes.fechapago, pagospendientes.idTipopago, pagospendientes.numeroTarjeta, pagospendientes.importe " +
		" FROM pagospendientes WHERE id = ? )"
	copia, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	} else {
		copia.Exec(idPagopendiente)
	}

	//Incrementar las sesiones al usuario con el IDReserva
	if Sesiones != 0 {
		// update	idusuario = sesiones + sesiones nuevas
		insForm, err := db.Prepare("UPDATE usuarios SET sesionesbono=sesionesbono + ? WHERE id=?")
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Actualizando Base de Datos")
		} else {
			insForm.Exec(Sesiones, idUsuario)
			log.Printf("UPDATE: sesiones  %d", Sesiones)
		}
	}

	//Eliminación de los registros de pagos pendientes que han pasado a pagos definitivamente
	delForm, err := db.Prepare("DELETE FROM pagosPendientes WHERE id=?")
	if err != nil {
		panic(err.Error())
	} else {
		delForm.Exec(idPagopendiente)
	}
}
