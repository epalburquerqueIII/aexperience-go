package server

import (
	"log"
	"net/http"

	"../server/middleware"
)

func StartServer(hostname string, port string) error {

	// http.Handler
	host := hostname + ":" + port

	log.Printf("Listening on: %s", host)

	handler := middleware.NewHandler()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.Handle("/", handler)
	return http.ListenAndServe(host, nil)
}
