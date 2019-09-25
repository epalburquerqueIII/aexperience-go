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

// Horarios Pantalla de tratamiento de los horarios
func Horarios(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "horarios", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// HorariosList - json con los horarios
func HorariosList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT horarios.id, idEspacio, descripcion, fechaInicio, fechaFin , hora FROM horarios " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	h := model.Thorarios{}
	res := []model.Thorarios{}
	for selDB.Next() {

		err = selDB.Scan(&h.ID, &h.IdEspacio, &h.Descripcion, &h.Fechainicio, &h.Fechafinal, &h.Hora)
		//Si no hay fecha de baja, este campo aparece como activo
		//if h.FechaBaja == "0000-00-00" {
		//	usu.FechaBaja = "Activo"
		//} else {
		//Formato de fecha en español cuando está de baja
		//	t, _ := time.Parse("2006-01-02", usu.FechaBaja)
		//	usu.FechaBaja = t.Format("02-01-2006"

		//}
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error al cargar horarios"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, h)
		i++
	}

	var vrecords model.HorariosRecords
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

// HorariosCreate - Crear un Usuario
func HorariosCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	h := model.Thorarios{}
	if r.Method == "POST" {
		h.IdEspacio, _ = strconv.Atoi(r.FormValue("IdEspacio"))
		h.Descripcion = r.FormValue("Descripcion")
		h.Fechainicio = r.FormValue("Fechainicio")
		h.Fechafinal = r.FormValue("Fechafinal")
		h.Hora, _ = strconv.Atoi(r.FormValue("Hora"))
		insForm, err := db.Prepare("INSERT INTO horarios(idEspacio,descripcion,fechaInicio,fechaFin,hora) VALUES(?,?,?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando horario"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(h.IdEspacio, h.Descripcion, h.Fechainicio, h.Fechafinal, h.Hora)
		if err1 != nil {
			panic(err1.Error())
		}
		h.ID, err1 = res.LastInsertId()
		log.Printf("INSERT: idEspacio: %d   | hora: %d\n", h.IdEspacio, h.Hora)

	}
	var vrecord model.HorariosRecord
	vrecord.Result = "OK"
	vrecord.Record = h
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// HorariosUpdate Actualiza el usuario
func HorariosUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	h := model.Thorarios{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("ID"))
		h.ID = int64(i)
		h.IdEspacio, _ = strconv.Atoi(r.FormValue("idEspacio"))
		h.Descripcion = r.FormValue("Descripcion")
		h.Fechainicio = r.FormValue("Fechainicio")
		h.Fechafinal = r.FormValue("Fechafinal")
		h.Hora, _ = strconv.Atoi(r.FormValue("Hora"))
		insForm, err := db.Prepare("UPDATE horarios SET idEspacio=?, descripcion=?, fechaInicio=?, fechaFin =?,hora=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(h.IdEspacio, h.Descripcion, h.Fechainicio, h.Fechafinal, h.Hora, h.ID)
		log.Printf("INSERT: IdEspacio: %d   | Hora: %d\n", h.IdEspacio, h.Hora)
	}
	defer db.Close()
	var vrecord model.HorariosRecord
	vrecord.Result = "OK"
	vrecord.Record = h
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

//HorariosDelete Borra usuario de la DB
func HorariosDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	h := r.FormValue("ID")
	delForm, err := db.Prepare("DELETE FROM horarios WHERE id=?")
	if err != nil {

		panic(err.Error())
	}
	_, err1 := delForm.Exec(h)
	if err1 != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error Borrando usuario"
		a, _ := json.Marshal(verror)
		w.Write(a)
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.UsuarioRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)
	http.Redirect(w, r, "/", 301)
}

//HorariosBaja da de baja al usuario
//func HorariosBaja(w http.ResponseWriter, r *http.Request) {
////	h := r.FormValue("ID")
//delForm, err := db.Prepare("UPDATE horarios SET fechaBaja=CURDATE() WHERE id=?")
//if err != nil {

//	panic(err.Error())
//}
//_, err1 := delForm.Exec(usu)
//if err1 != nil {
//	var verror model.Resulterror
//	verror.Result = "ERROR"
//	verror.Error = "Error dando de baja al usuario"
//	a, _ := json.Marshal(verror)
//	w.Write(a)
//}
//log.Println("BAJA")
//defer db.Close()
//var vrecord model.UsuarioRecord
//vrecord.Result = "OK"
//a, _ := json.Marshal(vrecord)
//w.Write(a)

// 	// 	http.Redirect(w, r, "/", 301)
//}

// UsuariogetoptionsRoles Roles de usuario
//func UsuariogetoptionsRoles(w http.ResponseWriter, r *http.Request) {

//	db := database.DbConn()
//	selDB, err := db.Query("SELECT usuarios_roles.id, usuarios_roles.nombre from usuarios_roles Order by usuarios_roles.id")
//	if err != nil {
//		panic(err.Error())
//	}
//	elem := model.Option{}
//	vtabla := []model.Option{}
//	for selDB.Next() {
//		if err != nil {
//		}
//		vtabla = append(vtabla, elem)
//	}

//	var vtab model.Options
//	vtab.Result = "OK"
//	vtab.Options = vtabla
// create json response from struct
//	a, err := json.Marshal(vtab)
// Visualza
//	s := string(a)
//	fmt.Println(s)
//	w.Write(a)
//	defer db.Close()
//}
