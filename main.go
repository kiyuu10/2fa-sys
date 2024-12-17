package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kiyuu10/2fa-sys/config"
	"github.com/kiyuu10/2fa-sys/routes"
)

func main() {
	config.LoadConfig()
	// connect database
	config.ConnectDatabase()

	// init route
	router := gin.Default()

	// define routes
	routes.AuthRoutes(router)

	// run server
	router.Run(":8080")
}
