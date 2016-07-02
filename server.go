package main

import (
	"fmt"

	"./config"
	"./database"
	"./models/order"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Config
	err := config.Load("config/config.yaml")
	if err != nil {
		fmt.Println("Error: Failed to load configuration")
	}
	// Load Database
	err = database.Storage.Open(config.State.DatabasePath)
	if err != nil {
		fmt.Println("Error: Failed to load database")
	}
	// Initialize Gin HTTP Server
	r := gin.Default()
	r.Static("assets", "./assets")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", indexRoute)

	// REST API
	currentAPI := r.Group("/api/v" + config.State.APIVersion)
	orderEndpoint := currentAPI.Group("/order")
	{
		orderEndpoint.POST("/", order.Post)
		orderEndpoint.GET("/", order.List)
		orderEndpoint.GET("/:id", order.Get)
		orderEndpoint.PATCH("/:id", order.Patch)
		orderEndpoint.DELETE("/:id", order.Delete)
	}

	r.Run(":" + config.State.Port)
}

func indexRoute(c *gin.Context) {
	c.HTML(200, "guest.html",
		gin.H{
			"title": "Variable Title",
		})
}
