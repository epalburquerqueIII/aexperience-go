package util

import (
	"../model"
	"../model/database"
)

func menus(usertype int) Tmenu {
	db := database.DbConn()
	selDB, err := db.Query("SELECT menuparent.titulo AS Categoría, menuparent.tipo, menus.titulo, icono, url FROM menus LEFT OUTER JOIN menuparent ON menuparent.id = parentId ")
	if err != nil {
		panic(err.Error())
	}
	menu := model.Tmenu{}
	res := []model.Tmenu{}

	desp := model.Tmenudesplegable{}
	rst := []model.Tmenudesplegable{}
	for selDB.Next() {

		err = selDB.Scan(&menu.ParentTitle, &menu.Icono, &menu.Despliega, &desp.NomEnlace, &desp.Enlace)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, menu)
		rst = append(rst, desp)
	}
	defer db.Close()
}
