package controller

import (
	"fmt"
	"net/http"

	"../util/mdtojson"
)

func GetEventosmdtojson(w http.ResponseWriter, r *http.Request) {

	json, err := mdtojson.ProcessRepo("http://localhost:1313/content/eventos/", "./dir")

	if json != "" {
		fmt.Printf(json)
	}
	if err != nil {
		panic(err.Error())

	}
	w.Write([]byte(json))
}
