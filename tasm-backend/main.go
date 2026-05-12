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

			// Setup — auth required, Admin only for completion
			protected.POST("/setup/complete", handlers.CompleteSetup)

			// Dashboard
			protected.GET("/dashboard/stats", handlers.GetDashboardStats)

			// ---- Asset Management ----
			protected.GET("/assets", handlers.GetAssets)
			protected.GET("/assets/:id", handlers.GetAssetByID)
			protected.GET("/assets/:id/history", handlers.GetAssetHistory)
			protected.GET("/assets/by-tag/:tagId", handlers.GetAssetByTag)
			protected.POST("/assets", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.CreateAsset)
			protected.PUT("/assets/:id", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.UpdateAsset)
			protected.DELETE("/assets/:id", middleware.RequireRoles("System Admin"), handlers.DeleteAsset)
			protected.POST("/assets/:id/checkout", handlers.CheckoutAsset)
			protected.POST("/assets/:id/checkin", handlers.CheckinAsset)

			// ---- Consumables ----
			protected.GET("/consumables", handlers.GetConsumables)
			protected.GET("/consumables/:id", handlers.GetConsumableByID)
			protected.POST("/consumables", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.CreateConsumable)
			protected.PUT("/consumables/:id", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.UpdateConsumable)
			protected.DELETE("/consumables/:id", middleware.RequireRoles("System Admin"), handlers.DeleteConsumable)

			// ---- Maintenance Contracts ----
			protected.GET("/contracts", handlers.GetContracts)
			protected.GET("/contracts/:id", handlers.GetContractByID)
			protected.POST("/contracts", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.CreateContract)
			protected.PUT("/contracts/:id", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.UpdateContract)
			protected.DELETE("/contracts/:id", middleware.RequireRoles("System Admin"), handlers.DeleteContract)

			// ---- Work Orders ----
			protected.GET("/work-orders", handlers.GetWorkOrders)
			protected.GET("/work-orders/:id", handlers.GetWorkOrderByID)
			protected.POST("/work-orders", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.CreateWorkOrder)
			protected.PUT("/work-orders/:id", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.UpdateWorkOrder)
			protected.DELETE("/work-orders/:id", middleware.RequireRoles("System Admin"), handlers.DeleteWorkOrder)

			// ---- Audits ----
			protected.GET("/audits", handlers.GetAudits)
			protected.GET("/audits/:id", handlers.GetAuditByID)
			protected.POST("/audits", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.CreateAudit)
			protected.PUT("/audits/:id", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.UpdateAudit)
			protected.DELETE("/audits/:id", middleware.RequireRoles("System Admin"), handlers.DeleteAudit)
			// Scan asset in audit (any authenticated user — field auditors)
			protected.POST("/audits/:id/scan", handlers.ScanAssetInAudit)

			// ---- Discrepancies ----
			protected.GET("/discrepancies", handlers.GetDiscrepancies)
			protected.GET("/discrepancies/:id", handlers.GetDiscrepancyByID)
			protected.POST("/discrepancies", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.CreateDiscrepancy)
			protected.PUT("/discrepancies/:id", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.UpdateDiscrepancy)
			protected.DELETE("/discrepancies/:id", middleware.RequireRoles("System Admin"), handlers.DeleteDiscrepancy)
			protected.POST("/discrepancies/:id/resolve", middleware.RequireRoles("System Admin", "Facilities Head"), handlers.ResolveDiscrepancy)

			// ---- Procurements ----
			protected.GET("/procurements", handlers.GetProcurements)
			protected.GET("/procurements/:id", handlers.GetProcurementByID)
			protected.POST("/procurements", middleware.RequireRoles("System Admin", "Finance Manager"), handlers.CreateProcurement)
			protected.PUT("/procurements/:id", middleware.RequireRoles("System Admin", "Finance Manager"), handlers.UpdateProcurement)
			protected.DELETE("/procurements/:id", middleware.RequireRoles("System Admin"), handlers.DeleteProcurement)

			// ---- Reservations ----
			protected.GET("/reservations", handlers.GetReservations)
			protected.GET("/reservations/:id", handlers.GetReservationByID)
			protected.POST("/reservations", handlers.CreateReservation)
			protected.PUT("/reservations/:id", handlers.UpdateReservation)
			protected.DELETE("/reservations/:id", handlers.DeleteReservation)

			// ---- Locations ----
			protected.GET("/locations", handlers.GetLocations)
			protected.GET("/locations/:id", handlers.GetLocationByID)
			protected.POST("/locations", middleware.RequireRoles("System Admin"), handlers.CreateLocation)
			protected.PUT("/locations/:id", middleware.RequireRoles("System Admin"), handlers.UpdateLocation)
			protected.DELETE("/locations/:id", middleware.RequireRoles("System Admin"), handlers.DeleteLocation)

			// ---- Financial ----
			protected.GET("/ledgers", handlers.GetLedgers)
			protected.GET("/ledgers/:id", handlers.GetLedgerByID)
			protected.POST("/ledgers", middleware.RequireRoles("System Admin", "Finance Manager"), handlers.CreateLedger)
			protected.PUT("/ledgers/:id", middleware.RequireRoles("System Admin", "Finance Manager"), handlers.UpdateLedger)
			protected.DELETE("/ledgers/:id", middleware.RequireRoles("System Admin"), handlers.DeleteLedger)

			protected.GET("/leases", handlers.GetLeases)
			protected.GET("/leases/:id", handlers.GetLeaseByID)
			protected.POST("/leases", middleware.RequireRoles("System Admin", "Finance Manager"), handlers.CreateLease)
			protected.PUT("/leases/:id", middleware.RequireRoles("System Admin", "Finance Manager"), handlers.UpdateLease)
			protected.DELETE("/leases/:id", middleware.RequireRoles("System Admin"), handlers.DeleteLease)

			protected.GET("/depreciations", handlers.GetDepreciations)
			protected.GET("/depreciations/:id", handlers.GetDepreciationByID)
			protected.POST("/depreciations", middleware.RequireRoles("System Admin", "Finance Manager"), handlers.CreateDepreciation)
			protected.PUT("/depreciations/:id", middleware.RequireRoles("System Admin", "Finance Manager"), handlers.UpdateDepreciation)
			protected.DELETE("/depreciations/:id", middleware.RequireRoles("System Admin"), handlers.DeleteDepreciation)

			protected.GET("/software-licenses", handlers.GetSoftwareLicenses)
			protected.GET("/software-licenses/:id", handlers.GetSoftwareLicenseByID)
			protected.POST("/software-licenses", middleware.RequireRoles("System Admin", "Finance Manager"), handlers.CreateSoftwareLicense)
			protected.PUT("/software-licenses/:id", middleware.RequireRoles("System Admin", "Finance Manager"), handlers.UpdateSoftwareLicense)
			protected.DELETE("/software-licenses/:id", middleware.RequireRoles("System Admin"), handlers.DeleteSoftwareLicense)

			// ---- User Management (Admin only) ----
			protected.GET("/users", middleware.RequireRoles("System Admin"), handlers.GetUsers)
			protected.GET("/users/:id", middleware.RequireRoles("System Admin"), handlers.GetUserByID)
			protected.POST("/users", middleware.RequireRoles("System Admin"), handlers.CreateUser)
			protected.PUT("/users/:id", middleware.RequireRoles("System Admin"), handlers.UpdateUser)
			protected.DELETE("/users/:id", middleware.RequireRoles("System Admin"), handlers.DeleteUser)

			protected.GET("/roles", handlers.GetRoles)
			protected.GET("/roles/:id", handlers.GetRoleByID)
			protected.POST("/roles", middleware.RequireRoles("System Admin"), handlers.CreateRole)
			protected.PUT("/roles/:id", middleware.RequireRoles("System Admin"), handlers.UpdateRole)
			protected.DELETE("/roles/:id", middleware.RequireRoles("System Admin"), handlers.DeleteRole)

			// ---- Alerts ----
			protected.GET("/alerts", handlers.GetAlerts)
			protected.GET("/alerts/:id", handlers.GetAlertByID)
			protected.PUT("/alerts/:id", handlers.UpdateAlert)
			protected.DELETE("/alerts/:id", middleware.RequireRoles("System Admin"), handlers.DeleteAlert)
			protected.POST("/alerts/mark-all-read", handlers.MarkAllAlertsRead)
			protected.POST("/alerts/generate", middleware.RequireRoles("System Admin"), handlers.GenerateAlerts)
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
