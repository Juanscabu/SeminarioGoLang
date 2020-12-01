package router

import (
	"net/http"

	"github.com/Juanscabu/SeminarioGoLang/Entregable/controller/autoController"
	"github.com/dimfeld/httptreemux/v5"
)

// Start ...
func Start() {
	router := httptreemux.New()

	router.POST("/autos", autoController.SaveAutoHandler)
	router.GET("/autos/:id", autoController.FindByIDAutoHandler)
	router.GET("/autos", autoController.FindAllAutoHandler)
	router.PUT("/autos/:id", autoController.UpdateAutoHandler)
	router.DELETE("/autos/:id", autoController.RemoveAutoHandler)

	http.ListenAndServe(":8080", router)
}
