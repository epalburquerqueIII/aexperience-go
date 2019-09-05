package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../model"
	"../model/database"
)

// Menu
func Menu(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "menu", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// MenuList
func MenuList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT menus.id, icono, parent_id, menus.titulo, url FROM menus " + jtsort)
	if err != nil {
		var verror model.Resulterror
		verror.Result = "ERROR"
		verror.Error = "Error buscando datos"
		a, _ := json.Marshal(verror)
		w.Write(a)
		panic(err.Error())
	}
	menu := model.Tmenu{}
	res := []model.Tmenu{}
	for selDB.Next() {

		err = selDB.Scan(&menu.ID, &menu.Icono, &menu.ParentID, &menu.NomEnlace, &menu.Enlace)
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error Cargando menu"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res = append(res, menu)
		i++
	}
	// var vrecords model.UsuarioRecords
	// vrecords.Result = "OK"
	// vrecords.TotalRecordCount = i
	// vrecords.Records = res

	// create json response from struct
	// a, err := json.Marshal(vrecords)
	// Visualza
	// 	s := string(a)
	// 	fmt.Println(s)
	// 	w.Write(a)
	// 	defer db.Close()
}

// MenugetTitulo
func MenugetTitulo(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT menu_parent.id, menu_parent.titulo FROM menu_parent ORDER BY menu_parent.id")
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
