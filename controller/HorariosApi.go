//TODO dar de alta los usuarios que están de baja

package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"../model"
	"../model/database"
)

var tmpl = template.Must(template.ParseGlob("views/*.html"))

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
	selDB, err := db.Query("SELECT horarios.id, idEspacio, descripcion, fechaInicio, fechaFin, hora FROM horarios " + jtsort)
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
		//} else {)
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
		insForm, err := db.Prepare("INSERT INTO Horarios(idEspacio,descripcion,fechaInicio,fechaFin,hora) VALUES(?,?,?,?,?)")
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
		log.Println("INSERT: IdEspacio: " + h.IdEspacio + " | Descripcion: " + h.Descripcion)

	}
	var vrecord model.UsuarioRecord
	vrecord.Result = "OK"
	vrecord.Record = usu
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// UsuarioUpdate Actualiza el usuario
func UsuarioUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	usu := model.Tusuario{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("ID"))
		usu.ID = int64(i)
		usu.Nombre = r.FormValue("Nombre")
		usu.Nif = r.FormValue("Nif")
		usu.Email = r.FormValue("Email")
		usu.Tipo, _ = strconv.Atoi(r.FormValue("Tipo"))
		usu.Telefono = r.FormValue("Telefono")
		usu.SesionesBonos, _ = strconv.Atoi(r.FormValue("SesionesBonos"))
		usu.Newsletter, _ = strconv.Atoi(r.FormValue("Newsletter"))
		usu.FechaBaja = r.FormValue("FechaBaja")
		insForm, err := db.Prepare("UPDATE usuarios SET nombre=?, nif=?, email=?, tipo =?, telefono=?, sesionesBonos=?, newsletter=?, fechaBaja=? WHERE id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(usu.Nombre, usu.Nif, usu.Email, usu.Tipo, usu.Telefono, usu.SesionesBonos, usu.Newsletter, usu.FechaBaja, usu.ID)
		log.Println("UPDATE: nombre: " + usu.Nombre + " | nif: " + usu.Nif)
	}
	defer db.Close()
	var vrecord model.UsuarioRecord
	vrecord.Result = "OK"
	vrecord.Record = usu
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

//UsuarioBaja da de baja al usuario
func UsuarioBaja(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	usu := r.FormValue("ID")
	delForm, err := db.Prepare("UPDATE usuarios SET fechaBaja=CURDATE() WHERE id=?")
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
	var vrecord model.UsuarioRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	// 	// 	http.Redirect(w, r, "/", 301)
}

// UsuariogetoptionsRoles Roles de usuario
func UsuariogetoptionsRoles(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT usuarios_roles.id, usuarios_roles.nombre from usuarios_roles Order by usuarios_roles.id")
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