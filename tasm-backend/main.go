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
	"tasm-backend/middleware"
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

	allowedOrigin := os.Getenv("FRONTEND_ORIGIN")

	r.Use(func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if allowedOrigin == "" || origin == allowedOrigin {
			if origin != "" {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			} else if allowedOrigin != "" {
				c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			}
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if c.Request.Method == http.MethodOptions {
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

		// Auth Routes (Unprotected)
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
		}

		// Protected Routes
		protected := api.Group("/")
		protected.Use(middleware.AuthRequired())
		{
			protected.GET("/auth/me", handlers.GetMe)

			protected.GET("/assets", handlers.GetAssets)
			protected.GET("/assets/:id", handlers.GetAssetByID)
			protected.POST("/assets", handlers.CreateAsset)
			protected.PUT("/assets/:id", handlers.UpdateAsset)
			protected.DELETE("/assets/:id", handlers.DeleteAsset)

			protected.GET("/consumables", handlers.GetConsumables)
			protected.GET("/consumables/:id", handlers.GetConsumableByID)
			protected.PUT("/consumables/:id", handlers.UpdateConsumable)
			protected.DELETE("/consumables/:id", handlers.DeleteConsumable)
			protected.POST("/consumables", handlers.CreateConsumable)

			protected.GET("/contracts", handlers.GetContracts)
			protected.GET("/contracts/:id", handlers.GetContractByID)
			protected.PUT("/contracts/:id", handlers.UpdateContract)
			protected.DELETE("/contracts/:id", handlers.DeleteContract)
			protected.POST("/contracts", handlers.CreateContract)

			protected.GET("/work-orders", handlers.GetWorkOrders)
			protected.GET("/work-orders/:id", handlers.GetWorkOrderByID)
			protected.PUT("/work-orders/:id", handlers.UpdateWorkOrder)
			protected.DELETE("/work-orders/:id", handlers.DeleteWorkOrder)
			protected.POST("/work-orders", handlers.CreateWorkOrder)

			protected.GET("/audits", handlers.GetAudits)
			protected.GET("/audits/:id", handlers.GetAuditByID)
			protected.PUT("/audits/:id", handlers.UpdateAudit)
			protected.DELETE("/audits/:id", handlers.DeleteAudit)
			protected.POST("/audits", handlers.CreateAudit)
			protected.GET("/discrepancies", handlers.GetDiscrepancies)
			protected.GET("/discrepancies/:id", handlers.GetDiscrepancyByID)
			protected.PUT("/discrepancies/:id", handlers.UpdateDiscrepancy)
			protected.DELETE("/discrepancies/:id", handlers.DeleteDiscrepancy)

			protected.GET("/procurements", handlers.GetProcurements)
			protected.GET("/procurements/:id", handlers.GetProcurementByID)
			protected.PUT("/procurements/:id", handlers.UpdateProcurement)
			protected.DELETE("/procurements/:id", handlers.DeleteProcurement)
			protected.POST("/procurements", handlers.CreateProcurement)

			protected.GET("/reservations", handlers.GetReservations)
			protected.GET("/reservations/:id", handlers.GetReservationByID)
			protected.POST("/reservations", handlers.CreateReservation)
			protected.PUT("/reservations/:id", handlers.UpdateReservation)
			protected.DELETE("/reservations/:id", handlers.DeleteReservation)

			protected.GET("/locations", handlers.GetLocations)
			protected.GET("/locations/:id", handlers.GetLocationByID)
			protected.PUT("/locations/:id", handlers.UpdateLocation)
			protected.DELETE("/locations/:id", handlers.DeleteLocation)
			protected.POST("/locations", handlers.CreateLocation)

			// Financial APIs
			protected.GET("/ledgers", handlers.GetLedgers)
			protected.GET("/ledgers/:id", handlers.GetLedgerByID)
			protected.POST("/ledgers", handlers.CreateLedger)
			protected.PUT("/ledgers/:id", handlers.UpdateLedger)
			protected.DELETE("/ledgers/:id", handlers.DeleteLedger)
			protected.GET("/leases", handlers.GetLeases)
			protected.GET("/leases/:id", handlers.GetLeaseByID)
			protected.POST("/leases", handlers.CreateLease)
			protected.PUT("/leases/:id", handlers.UpdateLease)
			protected.DELETE("/leases/:id", handlers.DeleteLease)
			protected.GET("/depreciations", handlers.GetDepreciations)
			protected.GET("/depreciations/:id", handlers.GetDepreciationByID)
			protected.POST("/depreciations", handlers.CreateDepreciation)
			protected.PUT("/depreciations/:id", handlers.UpdateDepreciation)
			protected.DELETE("/depreciations/:id", handlers.DeleteDepreciation)
			protected.GET("/software-licenses", handlers.GetSoftwareLicenses)
			protected.GET("/software-licenses/:id", handlers.GetSoftwareLicenseByID)
			protected.POST("/software-licenses", handlers.CreateSoftwareLicense)
			protected.PUT("/software-licenses/:id", handlers.UpdateSoftwareLicense)
			protected.DELETE("/software-licenses/:id", handlers.DeleteSoftwareLicense)

			// User Management APIs
			protected.GET("/users", handlers.GetUsers)
			protected.GET("/users/:id", handlers.GetUserByID)
			protected.POST("/users", handlers.CreateUser)
			protected.PUT("/users/:id", handlers.UpdateUser)
			protected.DELETE("/users/:id", handlers.DeleteUser)
			protected.GET("/roles", handlers.GetRoles)
			protected.GET("/roles/:id", handlers.GetRoleByID)
			protected.POST("/roles", handlers.CreateRole)
			protected.PUT("/roles/:id", handlers.UpdateRole)
			protected.DELETE("/roles/:id", handlers.DeleteRole)

			// Settings & Dashboard APIs
			protected.GET("/alerts", handlers.GetAlerts)
			protected.GET("/alerts/:id", handlers.GetAlertByID)
			protected.PUT("/alerts/:id", handlers.UpdateAlert)
			protected.DELETE("/alerts/:id", handlers.DeleteAlert)
		}
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
