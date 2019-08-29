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

// ConsumoBonosList - json con los datos de clientes
func ConsumoBonosList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT consumoBonos.id, consumoBonos.fecha, consumoBonos.sesiones, usuarios.nombre, espacios.descripcion, autorizados.nombreAutorizado FROM consumoBonos LEFT OUTER JOIN usuarios ON (usuarios.id = consumoBonos.idUsuario) LEFT OUTER JOIN espacios ON (espacios.id = consumoBonos.idEspacio) LEFT OUTER JOIN autorizados ON (autorizados.id = consumoBonos.idAutorizado) " + jtsort)
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

		err = selDB.Scan(&consu.ID, &consu.Fecha, &consu.Sesiones, &consu.IDUsuario, &consu.IDEspacio, &consu.IDAutorizado)
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
	usu := model.Tusuario{}
	if r.Method == "POST" {
		usu.Nombre = r.FormValue("Nombre")
		usu.Nif = r.FormValue("Nif")
		usu.Email = r.FormValue("Email")
		usu.Tipo, _ = strconv.Atoi(r.FormValue("Tipo"))
		usu.Telefono = r.FormValue("Telefono")
		usu.SesionesBonos, _ = strconv.Atoi(r.FormValue("SesionesBonos"))
		usu.Newsletter, _ = strconv.Atoi(r.FormValue("Newsletter"))
		usu.FechaBaja = r.FormValue("FechaBaja")
		insForm, err := db.Prepare("INSERT INTO usuarios(nombre, nif, email, tipo, telefono, sesionesBonos, newsletter, fechaBaja) VALUES(?,?,?,?,?,?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Insertando Usuario"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(usu.Nombre, usu.Nif, usu.Email, usu.Tipo, usu.Telefono, usu.SesionesBonos, usu.Newsletter, usu.FechaBaja)
		if err1 != nil {
			panic(err1.Error())
		}
		usu.ID, err1 = res.LastInsertId()
		log.Println("INSERT: nombre: " + usu.Nombre + " | nif: " + usu.Nif)

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

// UsuariogetoptionsRoles Roles de usuario
// func UsuariogetoptionsRoles(w http.ResponseWriter, r *http.Request) {

// 	db := database.DbConn()
// 	selDB, err := db.Query("SELECT usuarios_roles.id, usuarios_roles.nombre from usuarios_roles Order by usuarios_roles.id")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	elem := model.Option{}
// 	vtabla := []model.Option{}
// 	for selDB.Next() {
// 		err = selDB.Scan(&elem.Value, &elem.DisplayText)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		vtabla = append(vtabla, elem)
// 	}

// 	var vtab model.Options
// 	vtab.Result = "OK"
// 	vtab.Options = vtabla
// 	// create json response from struct
// 	a, err := json.Marshal(vtab)
// 	// Visualza
// 	s := string(a)
// 	fmt.Println(s)
// 	w.Write(a)
// 	defer db.Close()
// }
