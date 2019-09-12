package database

import (
	"database/sql"
	"fmt"
)

// DbConn Conectar a db
func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "epalburquerque"
	dbPass := "ClasE0099_"
	//var host = "tcp(192.168.0.19)"
	var host = "tcp(127.0.0.1)"
	dbname := "dbaexperience"
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@%s/%s", dbUser, dbPass, host, dbname))
	if err != nil {
		panic(err.Error())
	}
	return db
}
