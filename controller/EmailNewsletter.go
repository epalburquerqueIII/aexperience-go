package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../model"
	"../model/database"
	"../util"
)

// NoticiasList
func TipoNoticiasList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT id, nombre FROM tiponoticias " + jtsort)
	if err != nil {
		util.ErrorApi(err.Error(), w, "Error en Select ")
	}
	tipo := model.TtipoNoticia{}
	res := []model.TtipoNoticia{}
	for selDB.Next() {

		err = selDB.Scan(&tipo.Id, &tipo.Nombre)
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando datos de tiponoticias")
		}
		res = append(res, tipo)
		i++
	}

	var vrecords model.TipoNoticiaRecords
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

	error := tmpl.ExecuteTemplate(w, "tiponoticias", &parametros)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}

}

// Guardar tipo de noticias y email
func Newsletterguardar(w http.ResponseWriter, r *http.Request) {
	mceEmail := r.FormValue("EMAIL")
	db := database.DbConn()
	delForm, err := db.Prepare("DELETE FROM newsletter WHERE email = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = delForm.Exec(mceEmail)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "POST" {
		for i := 1; i <= 10; i++ {
			mceEmail := r.FormValue("EMAIL")
			if r.FormValue("nwch"+strconv.Itoa(i)) == "1" {
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
	// TODO cambiar address por variable global
	http.Redirect(w, r, "http://192.168.0.3:1313/", 301)
}
