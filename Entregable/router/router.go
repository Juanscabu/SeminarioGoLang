package router

import (
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
)

func Start() {
	router := httptreemux.New()

	router.POST("/autos", autoController.SaveAutoHandler)
	router.GET("/autos/:id", autoController.FindByIdAutoHandler)
	router.GET("/autos", autoController.FindAllAutoHandler)
	router.PUT("/autos/:id", autosController.UpdateAutoHandler)
	router.DELETE("/autos/:id", autosController.RemoveAutoHandler)

	http.ListenAndServe(":8080", router)
}
