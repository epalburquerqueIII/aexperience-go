//TODO dar de alta los usuarios que están de baja

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
)

// Espacio Pantalla de tratamiento de Espacio
func Espacio(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "espacios", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// Espacio - json con los datos de Espacio
func EspacioList(w http.ResponseWriter, r *http.Request) {

	var i int
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT espacios.id, descripcion, estado, modo, precio, tiposevento.id, aforo, fecha, numeroReservaslimite FROM espacios LEFT OUTER JOIN tiposevento on (tiposevento.id = idTipoevento) " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	esp := model.Tespacios{}
	res := []model.Tespacios{}
	for selDB.Next() {

		err = selDB.Scan(&esp.ID, &esp.Descripcion, &esp.Estado, &esp.Modo, &esp.Precio, &esp.IDTipoevento, &esp.Aforo, &esp.Fecha, &esp.NumeroReservaslimite)
		//Si no hay fecha de baja, este campo aparece como activo
		if esp.Fecha == "0000-00-00" {
			esp.Fecha = "Activo"
		} else {
			//Formato de fecha en español cuando está de baja
			t, _ := time.Parse("2006-01-02", esp.Fecha)
			esp.Fecha = t.Format("02-01-2006")

		}
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Cargando registros de Usuarios"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, esp)
		i++
	}

	var vrecords model.EspacioRecords
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

//EspacioCreate - Crear un Espacio
func EspacioCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	esp := model.Tespacios{}
	if r.Method == "POST" {
		esp.Descripcion = r.FormValue("Descripcion")
		esp.Estado = r.FormValue("Estado")
		esp.Modo = r.FormValue("Modo")
		esp.Precio, _ = strconv.Atoi(r.FormValue("Precio"))
		esp.IDTipoevento, _ = strconv.Atoi(r.FormValue("IdTipoevento"))
		esp.Aforo, _ = strconv.Atoi(r.FormValue("Aforo"))
		esp.Fecha = r.FormValue("Fecha")
		esp.NumeroReservaslimite, _ = strconv.Atoi(r.FormValue("NumeroReservaslimite"))
		insForm, err := db.Prepare("INSERT INTO espacios(descripcion, estado, modo, precio, idTipoevento, aforo, fecha, numeroReservaslimite) VALUES(?,?,?,?,?,?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Usuario"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(esp.Descripcion, esp.Estado, esp.Modo, esp.Precio, esp.IDTipoevento, esp.Aforo, esp.Fecha, esp.NumeroReservaslimite)
		if err1 != nil {
			panic(err1.Error())
		}
		esp.ID, err1 = res.LastInsertId()
		log.Println("INSERT: descripcion: " + esp.Descripcion + " | estado: " + esp.Estado)

	}
	var vrecord model.EspacioRecord
	vrecord.Result = "OK"
	vrecord.Record = esp
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// EspacioUpdate Actualiza el Espacio
func EspacioUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	esp := model.Tespacios{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("ID"))
		esp.ID = int64(i)
		esp.Descripcion = r.FormValue("Descripcion")
		esp.Estado = r.FormValue("Estado")
		esp.Modo = r.FormValue("Modo")
		esp.Precio, _ = strconv.Atoi(r.FormValue("Precio"))
		esp.IDTipoevento, _ = strconv.Atoi(r.FormValue("IdTipoevento"))

		esp.Aforo, _ = strconv.Atoi(r.FormValue("Aforo"))
		esp.Fecha = r.FormValue("Fecha")
		esp.NumeroReservaslimite, _ = strconv.Atoi(r.FormValue("numeroReservaslimite"))
		insForm, err := db.Prepare("UPDATE espacios SET descripcion=?, estado=?, modo=?, precio =?, idTipoevento=?, aforo=?, fecha=?, numeroReservaslimite=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(esp.Descripcion, esp.Estado, esp.Modo, esp.Precio, esp.IDTipoevento, esp.Aforo, esp.Fecha, esp.NumeroReservaslimite, esp.ID)
		log.Println("UPDATE: descripcion: " + esp.Descripcion + " | estado: " + esp.Estado)
	}
	defer db.Close()
	var vrecord model.EspacioRecord
	vrecord.Result = "OK"
	vrecord.Record = esp
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

//EspaciosBaja da de baja al usuario
/* func EspaciosBaja(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	usu := r.FormValue("ID")
	delForm, err := db.Prepare("UPDATE espacios SET fecha=CURDATE() WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(usu)
	if err1 != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error dando de baja al usuario"
		a, _ := json.Marshal(verror)
		w.Write(a)
	}
	log.Println("BAJA")
	defer db.Close()
	var vrecord model.EspacioRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	// 	// 	http.Redirect(w, r, "/", 301)
} */

// EspaciosgetoptionsRoles Roles de usuario
func Espaciosgetoptionsespacios(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT tiposevento.id, tiposevento.nombre from tiposevento Order by tiposevento.id")
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
