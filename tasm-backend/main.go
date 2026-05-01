package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	api := r.Group("/api")
	{
		api.POST("/seed", handlers.SeedDatabase)

		api.GET("/assets", handlers.GetAssets)
		api.GET("/assets/:id", handlers.GetAssetByID)
		api.POST("/assets", handlers.CreateAsset)
		api.PUT("/assets/:id", handlers.UpdateAsset)
		api.DELETE("/assets/:id", handlers.DeleteAsset)

		api.GET("/consumables", handlers.GetConsumables)
		api.PUT("/consumables/:id", handlers.UpdateConsumable)
		api.DELETE("/consumables/:id", handlers.DeleteConsumable)
		api.POST("/consumables", handlers.CreateConsumable)

		api.GET("/contracts", handlers.GetContracts)
		api.PUT("/contracts/:id", handlers.UpdateContract)
		api.DELETE("/contracts/:id", handlers.DeleteContract)
		api.POST("/contracts", handlers.CreateContract)

		api.GET("/work-orders", handlers.GetWorkOrders)
		api.PUT("/work-orders/:id", handlers.UpdateWorkOrder)
		api.DELETE("/work-orders/:id", handlers.DeleteWorkOrder)
		api.POST("/work-orders", handlers.CreateWorkOrder)

		api.GET("/audits", handlers.GetAudits)
		api.PUT("/audits/:id", handlers.UpdateAudit)
		api.DELETE("/audits/:id", handlers.DeleteAudit)
		api.POST("/audits", handlers.CreateAudit)
		api.GET("/discrepancies", handlers.GetDiscrepancies)
		api.PUT("/discrepancies/:id", handlers.UpdateDiscrepancy)
		api.DELETE("/discrepancies/:id", handlers.DeleteDiscrepancy)

		api.GET("/procurements", handlers.GetProcurements)
		api.PUT("/procurements/:id", handlers.UpdateProcurement)
		api.DELETE("/procurements/:id", handlers.DeleteProcurement)
		api.POST("/procurements", handlers.CreateProcurement)

		api.GET("/reservations", handlers.GetReservations)
		api.PUT("/reservations/:id", handlers.UpdateReservation)
		api.DELETE("/reservations/:id", handlers.DeleteReservation)

		api.GET("/locations", handlers.GetLocations)
		api.PUT("/locations/:id", handlers.UpdateLocation)
		api.DELETE("/locations/:id", handlers.DeleteLocation)
		api.POST("/locations", handlers.CreateLocation)

		// Financial APIs
		api.GET("/ledgers", handlers.GetLedgers)
		api.PUT("/ledgers/:id", handlers.UpdateLedger)
		api.DELETE("/ledgers/:id", handlers.DeleteLedger)
		api.GET("/leases", handlers.GetLeases)
		api.PUT("/leases/:id", handlers.UpdateLease)
		api.DELETE("/leases/:id", handlers.DeleteLease)
		api.GET("/depreciations", handlers.GetDepreciations)
		api.PUT("/depreciations/:id", handlers.UpdateDepreciation)
		api.DELETE("/depreciations/:id", handlers.DeleteDepreciation)
		api.GET("/software-licenses", handlers.GetSoftwareLicenses)
		api.PUT("/software-licenses/:id", handlers.UpdateSoftwareLicense)
		api.DELETE("/software-licenses/:id", handlers.DeleteSoftwareLicense)

		// User Management APIs
		api.GET("/users", handlers.GetUsers)
		api.PUT("/users/:id", handlers.UpdateUser)
		api.DELETE("/users/:id", handlers.DeleteUser)
		api.GET("/roles", handlers.GetRoles)
		api.PUT("/roles/:id", handlers.UpdateRole)
		api.DELETE("/roles/:id", handlers.DeleteRole)

		// Settings & Dashboard APIs
		api.GET("/alerts", handlers.GetAlerts)
		api.PUT("/alerts/:id", handlers.UpdateAlert)
		api.DELETE("/alerts/:id", handlers.DeleteAlert)
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
