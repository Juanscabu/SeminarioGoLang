package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // adding sql driver support
)

// StartConn ...
func StartConn() *sql.DB {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "autos"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		panic(err.Error())
	}

	createDatabaseIfNotExist(db)
	createSchemaIfNotExists(db)
	return db
}

func createDatabaseIfNotExist(db *sql.DB) {
	s := "CREATE DATABASE IF NOT EXISTS autos;"
	_, err := db.Exec(s)
	if err != nil {
		panic(err.Error())
	}

	s = "USE autos;"
	_, err = db.Exec(s)
	if err != nil {
		panic(err.Error())
	}

	createSchemaIfNotExists(db)
}

func createSchemaIfNotExists(db *sql.DB) {
	schemaAuto := `CREATE TABLE IF NOT EXISTS auto (
		id_auto int NOT NULL AUTO_INCREMENT,
		modelo varchar(50) NOT NULL,
		marca varchar(50) NOT NULL,
		patente varchar(50) NOT NULL,
		CONSTRAINT PK_Auto PRIMARY KEY (id_auto)
	);`

	// execute a query on the server
	var _, err = db.Exec(schemaAuto)
	if err != nil {
		panic(err.Error())
	}
}
