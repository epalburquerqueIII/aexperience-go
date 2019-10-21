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

// PagoBono Pantalla de tratamiento de ConsumoBonos
func PagosBonos(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)
	error := tmpl.ExecuteTemplate(w, "pagosbonos", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// PagosBonosList - json con los datos de clientes
func PagosBonosList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT pagosBonos.id, idUsuario, fechaCompra, pagosBonos.fechaPago, pagosBonos.sesiones FROM pagosBonos LEFT OUTER JOIN usuarios ON usuarios.id = idUsuario " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	pagoBono := model.Tpagobono{}
	res := []model.Tpagobono{}
	for selDB.Next() {

		err = selDB.Scan(&pagoBono.ID, &pagoBono.IDUsuario, &pagoBono.FechaCompra, &pagoBono.FechaPago, &pagoBono.Sesiones)
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Cargando registros de Pago de bonos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, pagoBono)
		i++
	}

	var vrecords model.PagoBonoRecords
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

// PagosBonosCreate Crear un Usuario
func PagosBonosCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	pagoBono := model.Tpagobono{}
	if r.Method == "POST" {
		pagoBono.IDUsuario, _ = strconv.Atoi(r.FormValue("IDUsuario"))
		pagoBono.FechaCompra = util.DateSql(r.FormValue("FechaCompra"))
		pagoBono.FechaPago = util.DateSql(r.FormValue("FechaPago"))
		pagoBono.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		insForm, err := db.Prepare("INSERT INTO pagosBonos(idUsuario, FechaCompra, FechaPago, Sesiones) VALUES(?,?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Consumo de bono"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(pagoBono.IDUsuario, pagoBono.FechaCompra, pagoBono.FechaPago, pagoBono.Sesiones)
		if err1 != nil {
			panic(err1.Error())
		}
		pagoBono.ID, err1 = res.LastInsertId()
		log.Printf("INSERT: fecha: %s | sesiones:  %d\n", pagoBono.FechaCompra, pagoBono.Sesiones)

	}
	var vrecord model.PagoBonoRecord
	vrecord.Result = "OK"
	vrecord.Record = pagoBono
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// PagosBonosUpdate Actualiza el consumo de bono
func PagosBonosUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	pagoBono := model.Tpagobono{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("ID"))
		pagoBono.ID = int64(i)
		pagoBono.IDUsuario, _ = strconv.Atoi(r.FormValue("IDUsuario"))
		pagoBono.FechaCompra = util.DateSql(r.FormValue("FechaCompra"))
		pagoBono.FechaPago = util.DateSql(r.FormValue("FechaPago"))
		pagoBono.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		insForm, err := db.Prepare("UPDATE pagosBonos SET IDUsuario=?, FechaCompra=?, FechaPago=?, Sesiones =? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(pagoBono.IDUsuario, pagoBono.FechaCompra, pagoBono.FechaPago, pagoBono.Sesiones, pagoBono.ID)
		log.Printf("UPDATE: fecha: %s | sesiones: %d", pagoBono.FechaCompra, pagoBono.Sesiones)
	}
	defer db.Close()
	var vrecord model.PagoBonoRecord
	vrecord.Result = "OK"
	vrecord.Record = pagoBono
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

// UsuarioDelete Borra usuario de la DB
// func UsuarioDelete(w http.ResponseWriter, r *http.Request) {
// 	db := database.DbConn()
// 	usu := r.FormValue("ID")
// 	delForm, err := db.Prepare("DELETE FROM usuarios WHERE id=?")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	_, err1 := delForm.Exec(usu)
// 	if err1 != nil {
// 		var verror model.Resulterror
// 		verror.Result = "ERROR"
// 		verror.Error = "Error Borrando usuario"
// 		a, _ := json.Marshal(verror)
// 		w.Write(a)
// 	}
// 	log.Println("DELETE")
// 	defer db.Close()
// 	var vrecord model.UsuarioRecord
// 	vrecord.Result = "OK"
// 	a, _ := json.Marshal(vrecord)
// 	w.Write(a)

//http.Redirect(w, r, "/", 301)
// }
