package entity

import "github.com/gin-gonic/gin"

// Agencia entity
type Agencia struct {
	ID     int64  `json:"id"`
	Nombre string `json:"nombre"`
}

// ToJson ...
func (a Agencia) ToJson() gin.H {
	return gin.H{
		"id":     a.ID,
		"nombre": a.Nombre,
	}
}
