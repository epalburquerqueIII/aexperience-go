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

// Pantalla de tratamiento de Menus
func Menus(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "menus", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// MenusList - json con los datos de clientes
func MenusList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT menus.id, parent_id, orden, titulo, icono, url, hanledFunc FROM menus " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	menu := model.Tmenus{}
	res := []model.Tmenus{}
	for selDB.Next() {

		err = selDB.Scan(&menu.Id, &menu.Parent_id, &menu.Orden, &menu.Titulo, &menu.Icono, &menu.Url, &menu.HanledFunc)
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error cargando menus"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, menu)
		i++
	}

	var vrecords model.MenusRecords
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

// MenusCreate Crear campos de Menus
func MenusCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	menu := model.Tmenus{}
	if r.Method == "POST" {
		menu.Parent_id, _ = strconv.Atoi(r.FormValue("Parent_id"))
		menu.Orden, _ = strconv.Atoi(r.FormValue("Orden"))
		menu.Titulo = r.FormValue("Titulo")
		menu.Icono = r.FormValue("Icono")
		menu.Url = r.FormValue("Url")
		menu.HanledFunc = r.FormValue("HanledFunc")
		insForm, err := db.Prepare("INSERT INTO menus(parent_id, orden, titulo, icono, url, hanledFunc) VALUES(?,?,?,?,?,?)")

		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error al insertar campo del menu"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(menu.Parent_id, menu.Orden, menu.Titulo, menu.Icono, menu.Url, menu.HanledFunc)
		if err1 != nil {
			panic(err1.Error())
		}
		menu.Id, err1 = res.LastInsertId()
		log.Printf("INSERT: id: %d | parent_id: %d\n", menu.Id, menu.Parent_id)

	}
	var vrecord model.MenusRecord
	vrecord.Result = "OK"
	vrecord.Record = menu
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// MenusUpdate Actualiza el campo de Menus
func MenusUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	menu := model.Tmenus{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("Id"))
		menu.Id = int64(i)
		menu.Parent_id, _ = strconv.Atoi(r.FormValue("Parent_id"))
		menu.Orden, _ = strconv.Atoi(r.FormValue("Orden"))
		menu.Titulo = r.FormValue("Titulo")
		menu.Icono = r.FormValue("Icono")
		menu.Url = r.FormValue("Url")
		menu.HanledFunc = r.FormValue("HanledFunc")
		insForm, err := db.Prepare("UPDATE menus SET parent_id=?, orden=?, titulo=?, icono=?, url=?, hanledFunc=? WHERE menus.id=?")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Actualizando Base de Datos"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}

		insForm.Exec(menu.Parent_id, menu.Orden, menu.Titulo, menu.Icono, menu.Url, menu.HanledFunc, menu.Id)
		log.Println("UPDATE: id: %d  | parent_id: %d\n", menu.Id, menu.Parent_id)
	}
	defer db.Close()
	var vrecord model.MenusRecord
	vrecord.Result = "OK"
	vrecord.Record = menu
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

///MenusDelete Borra de la DB
func MenusDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	menu := r.FormValue("Id")
	delForm, err := db.Prepare("DELETE FROM menus WHERE id=?")
	log.Println(menu)
	if err != nil {
		panic(err.Error())
	}
	_, err1 := delForm.Exec(menu)
	if err1 != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error Borrado"
		a, _ := json.Marshal(verror)
		w.Write(a)
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.MenusRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	http.Redirect(w, r, "/", 301)
}
