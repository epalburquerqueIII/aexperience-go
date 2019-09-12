package controller

import (
	"fmt"
	"net/http"

	"../model"
	"../model/database"
)

// Menu -
func Menu(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "menu", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// MenuList -
func MenuList(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	selDB, err := db.Query("SELECT parentId, icono, titulo, url FROM menus ")
	if err != nil {
		panic(err.Error())
	}
	menu := model.Tmenu{}
	res := []model.Tmenu{}

	desp := model.Tmenudesplegable{}
	rst := []model.Tmenudesplegable{}
	for selDB.Next() {

		err = selDB.Scan(&menu.ParentTitle, &menu.Icono, &menu.)
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
	selDB, err := db.Query("SELECT menuParent.titulo FROM menuParent ORDER BY menuParent.id")
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
	selDB, err := db.Query("SELECT menuParent.tipo FROM menuParent ")
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
