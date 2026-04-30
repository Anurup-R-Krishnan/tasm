package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"tasm-backend/database"
	"tasm-backend/handlers"
)

func main() {
	// Load .env if exists
	_ = godotenv.Load()

	// Initialize DB
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}

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

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		api.GET("/assets", handlers.GetAssets)
		api.POST("/assets", handlers.CreateAsset)

		api.GET("/consumables", handlers.GetConsumables)
		api.POST("/consumables", handlers.CreateConsumable)

		api.GET("/contracts", handlers.GetContracts)
		api.POST("/contracts", handlers.CreateContract)

		api.GET("/work-orders", handlers.GetWorkOrders)
		api.POST("/work-orders", handlers.CreateWorkOrder)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := ":" + port
	log.Printf("Starting Go backend on %s", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("server startup failed: %v", err)
	}
}
