package util

import (
	"encoding/json"
	"net/http"

	"../model"
)

// ErrorApi Centraliza la gestion de errores de API's
func ErrorApi(textoPanic string, w http.ResponseWriter, textoError string) {

	var verror model.Resulterror
	verror.Result = "ERROR"
	verror.Error = textoError + "<--->" + textoPanic
	a, _ := json.Marshal(verror)
	w.Write(a)
	panic(textoPanic)
}
