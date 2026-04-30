package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"tasm-backend/database"
	"tasm-backend/handlers"
)

func main() {
	// Load .env if exists
	_ = godotenv.Load()

	// Initialize DB
	database.ConnectDB()

	// Setup Router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.GET("/assets", handlers.GetAssets)
		api.POST("/assets", handlers.CreateAsset)
		
		api.GET("/consumables", handlers.GetConsumables)
		api.POST("/consumables", handlers.CreateConsumable)
	}

	log.Println("Starting Go backend on :8080")
	r.Run(":8080")
}
