package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../model"
	"../model/database"
)

// ConsumoBonos Pantalla de tratamiento de ConsumoBonos
func ConsumoBonos(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "consumobonos", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// ConsumoBonosList - json con los datos de clientes
func ConsumoBonosList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT consumoBonos.id, consumoBonos.fecha, consumoBonos.sesiones, idUsuario, usuarios.nombre, idEspacio, idAutorizado FROM consumoBonos LEFT OUTER JOIN usuarios ON usuarios.id = idUsuario " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	consu := model.Tconsumo{}
	res := []model.Tconsumo{}
	for selDB.Next() {

		err = selDB.Scan(&consu.ID, &consu.Fecha, &consu.Sesiones, &consu.IDUsuario, &consu.NombreUsuario, &consu.IDEspacio, &consu.IDAutorizado)
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Cargando registros de Consumo de bonos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, consu)
		i++
	}

	var vrecords model.ConsumoBonosRecords
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

// ConsumoBonosCreate Crear un Usuario
func ConsumoBonosCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	consu := model.Tconsumo{}
	if r.Method == "POST" {
		consu.Fecha = r.FormValue("Fecha")
		consu.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		consu.IDUsuario, _ = strconv.Atoi(r.FormValue("IDUsuario"))
		consu.IDEspacio, _ = strconv.Atoi(r.FormValue("IDEspacio"))
		consu.IDAutorizado, _ = strconv.Atoi(r.FormValue("IDAutorizado"))
		insForm, err := db.Prepare("INSERT INTO consumoBonos(fecha, sesiones, idUsuario, idEspacio, idAutorizado) VALUES(?,?,?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Consumo de bono"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(consu.Fecha, consu.Sesiones, consu.IDUsuario, consu.IDEspacio, consu.IDAutorizado)
		if err1 != nil {
			panic(err1.Error())
		}
		consu.ID, err1 = res.LastInsertId()
		log.Printf("INSERT: fecha: %s | sesiones:  %d\n", consu.Fecha, consu.Sesiones)

	}
	var vrecord model.ConsumoBonosRecord
	vrecord.Result = "OK"
	vrecord.Record = consu
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// ConsumoBonosUpdate Actualiza el consumo de bono
func ConsumoBonosUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	consu := model.Tconsumo{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("ID"))
		consu.ID = int64(i)
		consu.Fecha = r.FormValue("Fecha")
		consu.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		consu.IDUsuario, _ = strconv.Atoi(r.FormValue("IDUsuario"))
		consu.IDEspacio, _ = strconv.Atoi(r.FormValue("IDEspacio"))
		consu.IDAutorizado, _ = strconv.Atoi(r.FormValue("IDAutorizado"))

		insForm, err := db.Prepare("UPDATE consumoBonos SET fecha=?, sesiones=?, idUsuario=?, idEspacio =?, idAutorizado=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(consu.Fecha, consu.Sesiones, consu.IDUsuario, consu.IDEspacio, consu.IDAutorizado, consu.ID)
		log.Printf("UPDATE: fecha: %s | sesiones: %d", consu.Fecha, consu.Sesiones)
	}
	defer db.Close()
	var vrecord model.ConsumoBonosRecord
	vrecord.Result = "OK"
	vrecord.Record = consu
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

// 	// 	// 	http.Redirect(w, r, "/", 301)
// }
