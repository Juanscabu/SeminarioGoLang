package main

//_ "github.com/mattn/go-sqlite3"
import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var db *sqlx.DB
	db, err := sqlx.Open("sqlite", ":memory")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	createSchema(db)

}

// user...
type user struct {
	id   int    "db:`id`"
	Name string "db:`name`"
}

func createSchema(db *sqlx.DB) {
	schema := "create table if not exists `user` (" +
		"id integer primary key not null auto_increment," +
		"name varchar(255),"
	db.Exec(schema)

	db.MustExec("insert into user (name) VALUES (?)", "Jane Doe")
	rows, err := db.Query("Select id, name From user")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var u user
		rows.Scan(&u.Name)
		fmt.Println(u)
	}
}
