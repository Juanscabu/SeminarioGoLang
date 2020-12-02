package main

import (
	"github.com/Juanscabu/SeminarioGoLang/Entregable/controller/agenciaController"
	"github.com/Juanscabu/SeminarioGoLang/Entregable/controller/autoController"
	"github.com/Juanscabu/SeminarioGoLang/Entregable/database"
	"github.com/Juanscabu/SeminarioGoLang/Entregable/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.StartConn()
	defer db.Close()

	agenciaController.Start(db)
	autoController.Start(db)
	router.Start()

}
