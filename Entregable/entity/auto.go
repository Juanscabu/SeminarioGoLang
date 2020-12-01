package entity

import (
	"github.com/gin-gonic/gin"
)

// Auto ...
type Auto struct {
	ID              int64  `json:"id"`
	Modelo          string `json:"modelo"`
	Marca           string `json:"marca"`
	Patente         string `json:"patente"`
	IDConcesionaria int64  `json:"idConcesionaria"`
}

func (a Auto) ToJson() gin.H {
	return gin.H{
		"id":      a.ID,
		"modelo":  a.Modelo,
		"marca":   a.Marca,
		"patente": a.Patente,
	}
}
