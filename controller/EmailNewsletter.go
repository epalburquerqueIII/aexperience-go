package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"../model"
	"../model/database"
	"../util"
)

// Select tipo de noticias y email
func TipoNoticias(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()

	selDB, err := db.Query("SELECT id, nombre FROM tiponoticias order by tiponoticias.id")
	if err != nil {
		util.ErrorApi(err.Error(), w, "Error en Select ")
	}
	mceEmail := r.FormValue("EMAIL")
	tipo := model.TtipoNoticia{}
	noticias := []model.TtipoNoticia{}
	type datos struct {
		Email    string
		Noticias []model.TtipoNoticia
	}

	for selDB.Next() {

		err = selDB.Scan(&tipo.Id, &tipo.Nombre)
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando datos de tiponoticias")
		}
		noticias = append(noticias, tipo)
	}
	parametros := datos{mceEmail, noticias}

	error := tmpl.ExecuteTemplate(w, "emailnewsletter", &parametros)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}

// Guardar tipo de noticias y email
func Newsletterguardar(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	if r.Method == "POST" {
		for i := 1; i <= 10; i++ {
			mceEmail := r.FormValue("EMAIL")
			if r.FormValue("ch"+strconv.Itoa(i)) == "1" {
				insForm, err := db.Prepare("INSERT INTO newsletter(email, idtiponoticias) VALUES(?,?)")
				if err != nil {
					panic(err.Error())
				}
				_, err = insForm.Exec(mceEmail, i)
				if err != nil {
					panic(err.Error())
				}
			}
		}
	}
	http.Redirect(w, r, "http://192.168.0.3:1313/", 301)

}
