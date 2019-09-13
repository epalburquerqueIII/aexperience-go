package util

import (
	"../model"
	"../model/database"
)

// Menus Estructura de menu para template
func Menus(usertype int) []model.Tmenu {
	db := database.DbConn()
	selDB, err := db.Query("SELECT menuparent.titulo, icono, menus.titulo, url, menuparent.tipo FROM menus LEFT OUTER JOIN menuparent ON menuparent.id = parentId")
	if err != nil {
		panic(err.Error())
	}
	menu := model.Tmenu{}
	res := []model.Tmenu{}

	desp := model.Tmenudesplegable{}
	rst := []model.Tmenudesplegable{}
	for selDB.Next() {

		err = selDB.Scan(&menu.ParentTitle, &menu.Icono, &desp.NomEnlace, &desp.Enlace, &menu.Despliega)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, menu)
		rst = append(rst, desp)
	}
	defer db.Close()
	return res
}
