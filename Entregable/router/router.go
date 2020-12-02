package router

import (
	"net/http"

	"github.com/Juanscabu/SeminarioGoLang/Entregable/controller/agenciaController"
	autoController "github.com/Juanscabu/SeminarioGoLang/Entregable/controller/autoController"
	"github.com/dimfeld/httptreemux/v5"
)

// Start ...
func Start() {
	router := httptreemux.New()

	router.POST("/agencias", agenciaController.SaveAgenciaHandler)
	router.GET("/agencias/:id", agenciaController.FindByIdAgenciaHandler)
	router.GET("/agencias", agenciaController.FindAllAgenciasHandler)
	router.PUT("/agencias/:id", agenciaController.UpdateAgenciaHandler)
	router.DELETE("/agencias/:id", agenciaController.RemoveAgenciaHandler)

	router.POST("/agencias/autos", autoController.SaveAutoHandler)
	router.GET("/agencias/autos/:id", autoController.FindByIDAutoHandler)
	router.GET("/agencias/autos", autoController.FindAllAutosHandler)
	router.GET("/agencias/:idAgencia/autos", autoController.FindAllAutosByAgenciaHandler)
	router.PUT("/agencias/autos/:id", autoController.UpdateAutoHandler)
	router.DELETE("/agencias/autos/:id", autoController.RemoveAutoHandler)

	http.ListenAndServe(":8080", router)
}
