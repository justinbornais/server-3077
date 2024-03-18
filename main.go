package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justinbornais/server-3077/handlers"
	"github.com/justinbornais/server-3077/utilities"
)

func main() {
	router := gin.Default()
	db := utilities.InitDB()
	defer db.Close()

	getUser := func(c *gin.Context) {
		handlers.GetUser(c, db)
	}

	createUser := func(c *gin.Context) {
		handlers.CreateUser(c, db)
	}

	router.GET("/users/:id", getUser)
	router.GET("/users/", createUser)

	router.Run(":1450")
}
