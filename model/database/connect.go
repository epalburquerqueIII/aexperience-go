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
<<<<<<< HEAD
	//var host = "tcp(192.168.0.19)"
	//var host = "tcp(192.168.0.20)"
	var host = "tcp(127.0.0.1)"
=======
	//	var host = "tcp(192.168.0.19)"
	var host = "tcp(127.0.0.1)"

>>>>>>> 7fec213a7854bb8e649b13509c3d112584c54ec4
	dbname := "dbaexperience"
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@%s/%s", dbUser, dbPass, host, dbname))
	if err != nil {
		panic(err.Error())
	}
	return db
}
