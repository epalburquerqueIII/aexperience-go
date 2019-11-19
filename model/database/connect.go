package database

import (
	"database/sql"
	"fmt"

	"../../config"
)

// DbConn Conectar a db
func DbConn() (db *sql.DB) {
	//var host = "tcp(192.168.0.14)"
	var host = "tcp(192.168.0.12)"
	// var host = "tcp(127.0.0.1)"
	dbname := "dbaexperience"
	db, err := sql.Open(config.DbDriver, fmt.Sprintf("%s:%s@%s/%s", config.DbUser, config.DbPass, host, dbname))
	if err != nil {
		panic(err.Error())
	}
	return db
}
