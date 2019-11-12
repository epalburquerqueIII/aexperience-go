package main

import (
	"log"

	"./model/authdb"
	"./server"
	"./server/middleware/myJwt"
	_ "github.com/go-sql-driver/mysql"
)

var host = "192.168.0.21"
var port = "8088"

func main() {
	// init the DB
	authdb.InitDB()

	// init the JWTs
	jwtErr := myJwt.InitJWT()
	if jwtErr != nil {
		log.Println("Error initializing the JWT's!")
		log.Fatal(jwtErr)
	}

	// start the server
	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error starting server!")
		log.Fatal(serverErr)
	}
}
