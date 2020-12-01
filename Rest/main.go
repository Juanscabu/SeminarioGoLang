package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/users/:name", getUsersHandler)
	r.POST("/users", addUsersHandler)
	r.Run()
}

func getUsersHandler(c *gin.Context) {
	name := c.Param("name")
	lastname := c.Query("lastname")
	c.JSON(200, gin.H{
		"name": name + " " + lastname,
	})
}

func addUsersHandler(c *gin.Context) {
	c.JSON(201, gin.H{
		"name":     "jhon",
		"lastname": "doe",
	})
}
