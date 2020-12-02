package entity

import (
	"github.com/gin-gonic/gin"
)

// Auto ...
type Auto struct {
	ID        int64  `json:"id"`
	Modelo    string `json:"modelo"`
	Marca     string `json:"marca"`
	Patente   string `json:"patente"`
	IdAgencia int64  `json:"idAgencia"`
}

// ToJSON ...
func (a Auto) ToJSON() gin.H {
	return gin.H{
		"id":        a.ID,
		"modelo":    a.Modelo,
		"marca":     a.Marca,
		"patente":   a.Patente,
		"idAgencia": a.IdAgencia,
	}
}
