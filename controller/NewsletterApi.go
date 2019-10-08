package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../model"
	"../model/database"
	"../util"
)

// Pantalla de tratamiento de Newsletter
func Newsletter(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "newsletter", nil)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// NewsletterList - json con los datos de clientes
func NewsletterList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT newsletter.id, Email, Idtiponoticias FROM newsletter " + jtsort)
	if err != nil {
		util.ErrorApi(err.Error(), w, "Error en Select ")
	}
	news := model.Tnewsletter{}
	res := []model.Tnewsletter{}
	for selDB.Next() {

		err = selDB.Scan(&news.Id, &news.Email, &news.Idtiponoticias)
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando datos de newsletter")
		}
		res = append(res, news)
		i++
	}

	var vrecords model.NewsletterRecords
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

// NewsletterCreate Crear campos de Newsletter
func NewsletterCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	news := model.Tnewsletter{}
	if r.Method == "POST" {
		news.Email = r.FormValue("Email")
		news.Idtiponoticias, _ = strconv.Atoi(r.FormValue("Idtiponoticias"))
		insForm, err := db.Prepare("INSERT INTO newsletter(email, idtiponoticias) VALUES(?,?)")

		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Insertando datos de newsletter")
		}
		res, err1 := insForm.Exec(news.Email, news.Idtiponoticias)
		if err1 != nil {
			panic(err1.Error())
		}
		news.Id, err1 = res.LastInsertId()
		log.Printf("INSERT: id: %d | email: %s\n", news.Id, news.Email)

	}
	var vrecord model.NewsletterRecord
	vrecord.Result = "OK"
	vrecord.Record = news
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// NewsletterUpdate Actualiza el campo de Newsletter
func NewsletterUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	news := model.Tnewsletter{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("Id"))
		news.Id = int64(i)
		news.Email = r.FormValue("Email")
		news.Idtiponoticias, _ = strconv.Atoi(r.FormValue("Idtiponoticias"))
		insForm, err := db.Prepare("UPDATE newsletter SET email=?, idtiponoticias=? WHERE newsletter.id=?")
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Actualizando Base de Datos")
		}

		insForm.Exec(news.Email, news.Idtiponoticias)
		log.Println("UPDATE: id: %d  | email: %d\n", news.Id, news.Email)
	}
	defer db.Close()
	var vrecord model.NewsletterRecord
	vrecord.Result = "OK"
	vrecord.Record = news
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

///NewsletterDelete Borra de la DB
func NewsletterDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	news := r.FormValue("Id")
	delForm, err := db.Prepare("DELETE FROM newsletter WHERE id=?")
	log.Println(news)
	if err != nil {
		panic(err.Error())
	}
	_, err1 := delForm.Exec(news)
	if err1 != nil {
		util.ErrorApi(err.Error(), w, "Error borrando datos de newsletter")
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.NewsletterRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	http.Redirect(w, r, "/", 301)
}

func NewslettergetoptionsTipoNoticias(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT tiponoticias.id, tiponoticias.nombre from tiponoticias Order by tiponoticias.id")
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
