package controller

import (
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
	db := database.DbConn()
	selDB, err := db.Query("Select parent_id, titulo, icono, url FROM menus ")
	if err != nil {
		panic(err.Error())
	}
	menu := model.Tmenu{}
	res := []model.Tmenu{}
	for selDB.Next() {

		err = selDB.Scan(&menu.ID, &menu.Icono, &menu.ParentID, &menu.NomEnlace, &menu.Enlace)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, menu)
	}
	defer db.Close()
}

// MenugetTitulo - para obtener el titulo del parent
func MenugetTitulo(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT menu_parent.titulo FROM menu_parent ORDER BY menu_parent.id")
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
	defer db.Close()
}

//MenugetTipo saber si es usuario o administrador
func MenugetTipo(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT menu_parent.tipo FROM menu_parent ")
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
	defer db.Close()
}
