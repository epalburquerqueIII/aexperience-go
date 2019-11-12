package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

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
