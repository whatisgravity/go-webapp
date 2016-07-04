package main

import (
	"fmt"
	"net/http"

	"./config"
	"./database"
	"./models/order"

	"github.com/gin-gonic/gin"
)

func main() {
	// CONFIG
	err := config.Load("config/config.yaml")
	if err != nil {
		fmt.Println("Error: Failed to load configuration")
	}
	// DATABASE
	fmt.Println("Opening database at...", config.State.DatabasePath)
	_, err = database.Storage.Open(config.State.DatabasePath)
	if err != nil {
		fmt.Println("Error: Failed to load database")
	}
	defer database.Storage.Close()
	// HTTP Server
	if !config.State.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// WEB UI
	r.Static("assets", "./assets")
	r.GET("/", index)
	// REST API
	currentAPI := r.Group("/api/v" + config.State.APIVersion)
	orderEndpoint := currentAPI.Group("/order")
	{
		orderEndpoint.GET("/", order.List)
		orderEndpoint.GET("/:id", order.Get)
		orderEndpoint.GET("/:id/delete", order.Delete)
		orderEndpoint.POST("/", order.Post)
		orderEndpoint.POST("/:id/patch", order.Patch)
	}
	r.Run(":" + config.State.Port)
}

func index(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "templates/js/index.html")
}
