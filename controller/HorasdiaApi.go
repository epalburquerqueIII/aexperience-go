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

// Espacios Pantalla de tratamiento de Espacio
func HorasDia(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)
	error := tmpl.ExecuteTemplate(w, "horasdia", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// EspaciosList - json con los datos de Espacio
func HorasDiaList(w http.ResponseWriter, r *http.Request) {

	var i int
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}

	db := database.DbConn()
	selDB, err := db.Query("SELECT horarios.id, horarios.idespacio, horarios.fechaFin,horarios.hora, IF(isnull(reservas.hora), '0', '1') as reservado from horarios left outer join reservas on reservas.fecha = '2019-11-11' and reservas.hora = horarios.hora and horarios.idespacio = reservas.idespacio where horarios.fechaInicio <= '2019-10-11' and horarios.fechaFin >= '2019-10-11' and horarios.idespacio = 1 " + jtsort)
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

// HorasDiaCreate - json con los datos de Espacio

func HorasDiaCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	esp := model.THorasdia{}
	if r.Method == "POST" {

		esp.Fecha = r.FormValue("Fecha")
		esp.Hora, _ = strconv.Atoi(r.FormValue("Hora"))
		esp.IDEspacio, _ = strconv.Atoi(r.FormValue("IDEspacio"))
		esp.IDUsuario, _ = strconv.Atoi(r.FormValue("IDUsuario"))
		esp.IDAutorizado, _ = strconv.Atoi(r.FormValue("IDAutorizado"))

		insForm, err := db.Prepare("INSERT INTO reservas(fecha, hora, idUsuario, idEspacio, idAutorizado) VALUES(?,?,?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Espacio"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(esp.Fecha, esp.Hora, esp.IDUsuario, esp.IDEspacio, esp.IDAutorizado)
		if err1 != nil {
			panic(err1.Error())
		}
		esp.ID, err1 = res.LastInsertId()
		log.Println("INSERT: Reservado: " + esp.Fecha)

	}
	var vrecord model.HorariosdiasRecord
	vrecord.Result = "OK"
	vrecord.Record = esp
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// HorasDiaUpdate - json con los datos de Espacio

func HorasDiaUpdate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	esp := model.THorasdia{}
	if r.Method == "POST" {

		esp.IDEspacio, _ = strconv.Atoi(r.FormValue("IDEspacio"))
		esp.Fecha = r.FormValue("Fecha")
		esp.Hora, _ = strconv.Atoi(r.FormValue("Hora"))
		esp.Reservado, _ = strconv.Atoi(r.FormValue("Reservado"))

		insForm, err := db.Prepare("UPDATE espacios SET descripcion=?, estado=?, modo=?, precio =?, idTipoevento=?, aforo=?, fecha=?, numeroReservaslimite=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Espacio"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(esp.ID, esp.IDEspacio, esp.Fecha, esp.Hora, esp.Reservado)
		if err1 != nil {
			panic(err1.Error())
		}
		esp.ID, err1 = res.LastInsertId()
		log.Println("INSERT: Reservado: " + esp.Fecha)

	}
	var vrecord model.HorariosdiasRecord
	vrecord.Result = "OK"
	vrecord.Record = esp
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}
