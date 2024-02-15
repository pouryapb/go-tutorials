package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pouryapb/go-tutorials/rest-api/db"
	"github.com/pouryapb/go-tutorials/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
