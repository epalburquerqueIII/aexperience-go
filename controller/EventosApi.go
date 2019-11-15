package controller

import (
	"fmt"
	"net/http"

	"../config"

	"../util/mdtojson"
)

// GetEventosmdtojson Obtiene los eventos de la Web en formato Json
func GetEventosmdtojson(w http.ResponseWriter, r *http.Request) {

	json, err := mdtojson.ProcessRepo(config.CmsHost+"/content/eventos/", "./util/dir")

	if json != "" {
		fmt.Printf(json)
	}
	if err != nil {
		panic(err.Error())

	}
	w.Write([]byte(json))
}
