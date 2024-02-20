package main

import (
	"gabefraser/minepass/handlers"
	"gabefraser/minepass/middleware"
	"gabefraser/minepass/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	password := os.Getenv("MP_UI_PASSWORD")
	if password == "" {
		utils.Logger("Missing UI password. Please add a password using MP_UI_PASSWORD environment variable, exiting...")
		os.Exit(1)
	}

	// Set up some defaults
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Not found"})
	})

	// API Routes
	v1 := r.Group("/api", middleware.GuardMiddleware, middleware.RconMiddleware)
	{
		v1.POST("/whitelist/add", handlers.WhitelistAdd)
		v1.POST("/whitelist/remove", handlers.WhitelistRemove)
		v1.POST("/validate", handlers.Validate)
	}

	// HTML Routes
	r.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": os.Getenv("MP_TITLE"),
		})
	})

	utils.Logger("Listening and serving MinePass on port :8080")

	r.Run(":8080")
}
