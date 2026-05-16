package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
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

	// Performance: GZIP Compression
	r.Use(gzip.Gzip(gzip.DefaultCompression))

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

		// Security Headers
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")

		c.Next()
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	api := r.Group("/api")
	{
		// Public seed (dev only) + auth
		api.POST("/seed", handlers.SeedDatabase)
		api.GET("/seed", handlers.SeedDatabase)

		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
			auth.POST("/register", handlers.Register)
			auth.POST("/create-admin", handlers.CreateAdmin) // first-run only: bootstrap
		}

		// Setup status — unauthenticated
		api.GET("/setup/status", handlers.GetSetupStatus)

		// ----------------------------------------------------------------
		// All routes below require a valid JWT
		// ----------------------------------------------------------------
		protected := api.Group("")
		protected.Use(middleware.AuthRequired())
		{
			protected.GET("/auth/me", handlers.GetMe)

			protected.GET("/config", handlers.GetSystemConfig)
			protected.PUT("/config", handlers.UpdateSystemConfig)

			// Setup — auth required, Admin only for completion
			protected.POST("/setup/complete", handlers.CompleteSetup)

			// Dashboard
			protected.GET("/dashboard/stats", handlers.GetDashboardStats)

			// ---- Asset Management ----
			protected.GET("/assets", handlers.GetAssets)
			protected.GET("/assets/:id", handlers.GetAssetByID)
			protected.GET("/assets/:id/history", handlers.GetAssetHistory)
			protected.GET("/assets/by-tag/:tagId", handlers.GetAssetByTag)
			protected.POST("/assets", handlers.CreateAsset)
			protected.PUT("/assets/:id", handlers.UpdateAsset)
			protected.DELETE("/assets/:id", handlers.DeleteAsset)
			protected.POST("/assets/:id/checkout", handlers.CheckoutAsset)
			protected.POST("/assets/:id/checkin", handlers.CheckinAsset)
			protected.POST("/assets/:id/transfer", handlers.TransferAsset)

			// ---- Consumables ----
			protected.GET("/consumables", handlers.GetConsumables)
			protected.GET("/consumables/:id", handlers.GetConsumableByID)
			protected.POST("/consumables", handlers.CreateConsumable)
			protected.PUT("/consumables/:id", handlers.UpdateConsumable)
			protected.DELETE("/consumables/:id", handlers.DeleteConsumable)

			// ---- Maintenance Contracts ----
			protected.GET("/contracts", handlers.GetContracts)
			protected.GET("/contracts/:id", handlers.GetContractByID)
			protected.POST("/contracts", handlers.CreateContract)
			protected.PUT("/contracts/:id", handlers.UpdateContract)
			protected.DELETE("/contracts/:id", handlers.DeleteContract)

			// ---- Work Orders ----
			protected.GET("/work-orders", handlers.GetWorkOrders)
			protected.GET("/work-orders/:id", handlers.GetWorkOrderByID)
			protected.POST("/work-orders", handlers.CreateWorkOrder)
			protected.PUT("/work-orders/:id", handlers.UpdateWorkOrder)
			protected.DELETE("/work-orders/:id", handlers.DeleteWorkOrder)

			// ---- Audits ----
			protected.GET("/audits", handlers.GetAudits)
			protected.GET("/audits/:id", handlers.GetAuditByID)
			protected.POST("/audits", handlers.CreateAudit)
			protected.PUT("/audits/:id", handlers.UpdateAudit)
			protected.DELETE("/audits/:id", handlers.DeleteAudit)
			// Scan asset in audit (any authenticated user — field auditors)
			protected.POST("/audits/:id/scan", handlers.ScanAssetInAudit)

			// ---- Discrepancies ----
			protected.GET("/discrepancies", handlers.GetDiscrepancies)
			protected.GET("/discrepancies/:id", handlers.GetDiscrepancyByID)
			protected.POST("/discrepancies", handlers.CreateDiscrepancy)
			protected.PUT("/discrepancies/:id", handlers.UpdateDiscrepancy)
			protected.DELETE("/discrepancies/:id", handlers.DeleteDiscrepancy)
			protected.POST("/discrepancies/:id/resolve", handlers.ResolveDiscrepancy)

			// ---- Procurements ----
			protected.GET("/procurements", handlers.GetProcurements)
			protected.GET("/procurements/:id", handlers.GetProcurementByID)
			protected.POST("/procurements", handlers.CreateProcurement)
			protected.PUT("/procurements/:id", handlers.UpdateProcurement)
			protected.DELETE("/procurements/:id", handlers.DeleteProcurement)

			// ---- Reservations ----
			protected.GET("/reservations", handlers.GetReservations)
			protected.GET("/reservations/:id", handlers.GetReservationByID)
			protected.POST("/reservations", handlers.CreateReservation)
			protected.PUT("/reservations/:id", handlers.UpdateReservation)
			protected.DELETE("/reservations/:id", handlers.DeleteReservation)

			// ---- Locations ----
			protected.GET("/locations", handlers.GetLocations)
			protected.GET("/locations/:id", handlers.GetLocationByID)
			protected.POST("/locations", handlers.CreateLocation)
			protected.PUT("/locations/:id", handlers.UpdateLocation)
			protected.DELETE("/locations/:id", handlers.DeleteLocation)

			// ---- Financial ----
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

			// ---- User Management ----
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

			// ---- Alerts ----
			protected.GET("/alerts", handlers.GetAlerts)
			protected.GET("/alerts/:id", handlers.GetAlertByID)
			protected.PUT("/alerts/:id", handlers.UpdateAlert)
			protected.DELETE("/alerts/:id", handlers.DeleteAlert)
			protected.POST("/alerts/mark-all-read", handlers.MarkAllAlertsRead)
			protected.POST("/alerts/generate", handlers.GenerateAlerts)
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
