package util

import (
	"../model"
	"../model/database"
)

const menudesplegable = 1

// Menus Estructura de menu para template
func Menus(usertype int) []model.Tmenuconfig {
	var ant string = ""
	var i int = -1
	db := database.DbConn()
	selDB, err := db.Query("SELECT menuparent.id, menuparent.titulo, menus.icono, menus.orden, menus.titulo, url, menuparent.tipo FROM menus LEFT OUTER JOIN menuparent ON menuparent.id = parentId ORDER BY menuparent.id, orden")
	if err != nil {
		panic(err.Error())
	}
	menu := model.Tmenuconfig{}
	res := []model.Tmenuconfig{}

	submenu := model.Tmenudesplegable{}

	for selDB.Next() {

		err = selDB.Scan(&menu.ID, &menu.ParentTitle, &menu.Icono, &submenu.Orden, &submenu.NomEnlace, &submenu.Enlace, &menu.Despliega)
		if err != nil {
			panic(err.Error())
		}
		if ant != menu.ParentTitle {
			res = append(res, menu)
			i++
			res[i].Options = nil
			if menu.Despliega == menudesplegable {
				//res[i].Options = append(res[i].Options, model.Tmenudesplegable{"HeaderMenu", ""})
				res[i].Options = append(res[i].Options, submenu)
			} else {
				res[i].Options = append(res[i].Options, submenu)
			}
		} else {
			if menu.Despliega == menudesplegable {
				res[i].Options = append(res[i].Options, submenu)
			}
		}
		ant = menu.ParentTitle
	}
	defer db.Close()
	return res
}
