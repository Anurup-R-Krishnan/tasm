package main

import (
	"log"
	"net/http"
	"os"

	"tasm-backend/database"
	"tasm-backend/handlers"
	"tasm-backend/middleware"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
		// Auth
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
			protected.POST("/setup/complete", middleware.RequireRoles("Admin"), handlers.CompleteSetup)

			// Dashboard
			protected.GET("/dashboard/stats", handlers.GetDashboardStats)

			// ---- Asset Management ----
			protected.GET("/assets", handlers.GetAssets)
			protected.GET("/assets/:id", handlers.GetAssetByID)
			protected.GET("/assets/:id/history", handlers.GetAssetHistory)
			protected.GET("/assets/:id/lifecycle", handlers.GetAssetFullLifecycle)
			protected.GET("/assets/:id/depreciation", handlers.GetAssetDepreciationSchedule)
			protected.GET("/assets/:id/transfers", handlers.GetAssetTransfers)
			protected.GET("/assets/:id/receipt", handlers.GetReceipt)
			protected.GET("/assets/by-tag/:tagId", handlers.GetAssetByTag)
			protected.POST("/assets", middleware.RequireRoles("Admin", "Asset Manager"), handlers.CreateAsset)
			protected.PUT("/assets/:id", middleware.RequireRoles("Admin", "Asset Manager"), handlers.UpdateAsset)
			protected.DELETE("/assets/:id", middleware.RequireRoles("Admin"), handlers.DeleteAsset)
			protected.POST("/assets/:id/checkout", handlers.CheckoutAsset)
			protected.POST("/assets/:id/checkin", handlers.CheckinAsset)
			protected.POST("/assets/:id/transfer", middleware.RequireRoles("Admin", "Asset Manager"), handlers.TransferAsset)
			protected.POST("/assets/:id/retire", middleware.RequireRoles("Admin", "Asset Manager"), handlers.RetireAsset)
			protected.POST("/assets/:id/dispose", middleware.RequireRoles("Admin", "Asset Manager"), handlers.DisposeAsset)
			protected.POST("/assets/:id/transfers/:transferId/revert", middleware.RequireRoles("Admin", "Asset Manager"), handlers.RevertTransfer)
			protected.POST("/assets/:id/receipt", middleware.RequireRoles("Admin", "Asset Manager"), handlers.UploadReceipt)

			// ---- Consumables ----
			protected.GET("/consumables", handlers.GetConsumables)
			protected.GET("/consumables/:id", handlers.GetConsumableByID)
			protected.POST("/consumables", middleware.RequireRoles("Admin", "Asset Manager"), handlers.CreateConsumable)
			protected.PUT("/consumables/:id", middleware.RequireRoles("Admin", "Asset Manager"), handlers.UpdateConsumable)
			protected.DELETE("/consumables/:id", middleware.RequireRoles("Admin"), handlers.DeleteConsumable)

			// ---- Maintenance ----
			maintenance := protected.Group("/maintenance")
			{
				maintenance.GET("/contracts", handlers.GetContracts)
				maintenance.GET("/contracts/:id", handlers.GetContractByID)
				maintenance.POST("/contracts", middleware.RequireRoles("Admin", "Asset Manager"), handlers.CreateContract)
				maintenance.PUT("/contracts/:id", middleware.RequireRoles("Admin", "Asset Manager"), handlers.UpdateContract)
				maintenance.DELETE("/contracts/:id", middleware.RequireRoles("Admin"), handlers.DeleteContract)

				maintenance.GET("/work-orders", handlers.GetWorkOrders)
				maintenance.GET("/work-orders/:id", handlers.GetWorkOrderByID)
				maintenance.POST("/work-orders", middleware.RequireRoles("Admin", "Asset Manager"), handlers.CreateWorkOrder)
				maintenance.PUT("/work-orders/:id", middleware.RequireRoles("Admin", "Asset Manager"), handlers.UpdateWorkOrder)
				maintenance.DELETE("/work-orders/:id", middleware.RequireRoles("Admin"), handlers.DeleteWorkOrder)
			}

			// ---- Audits ----
			protected.GET("/audits", handlers.GetAudits)
			protected.GET("/audits/:id", handlers.GetAuditByID)
			protected.POST("/audits", middleware.RequireRoles("Admin", "Auditor"), handlers.CreateAudit)
			protected.PUT("/audits/:id", middleware.RequireRoles("Admin", "Auditor"), handlers.UpdateAudit)
			protected.DELETE("/audits/:id", middleware.RequireRoles("Admin"), handlers.DeleteAudit)
			// Scan asset in audit (any authenticated user — field auditors)
			protected.POST("/audits/:id/scan", handlers.ScanAssetInAudit)

			// ---- Discrepancies ----
			protected.GET("/discrepancies", handlers.GetDiscrepancies)
			protected.GET("/discrepancies/:id", handlers.GetDiscrepancyByID)
			protected.POST("/discrepancies", middleware.RequireRoles("Admin", "Auditor"), handlers.CreateDiscrepancy)
			protected.PUT("/discrepancies/:id", middleware.RequireRoles("Admin", "Auditor"), handlers.UpdateDiscrepancy)
			protected.DELETE("/discrepancies/:id", middleware.RequireRoles("Admin"), handlers.DeleteDiscrepancy)
			protected.POST("/discrepancies/:id/resolve", middleware.RequireRoles("Admin", "Auditor"), handlers.ResolveDiscrepancy)

			// ---- Procurements ----
			protected.GET("/procurements", handlers.GetProcurements)
			protected.GET("/procurements/:id", handlers.GetProcurementByID)
			protected.POST("/procurements", middleware.RequireRoles("Admin", "Asset Manager"), handlers.CreateProcurement)
			protected.PUT("/procurements/:id", middleware.RequireRoles("Admin", "Asset Manager"), handlers.UpdateProcurement)
			protected.DELETE("/procurements/:id", middleware.RequireRoles("Admin"), handlers.DeleteProcurement)

			// ---- Reservations ----
			protected.GET("/reservations", handlers.GetReservations)
			protected.GET("/reservations/:id", handlers.GetReservationByID)
			protected.POST("/reservations", handlers.CreateReservation)
			protected.PUT("/reservations/:id", handlers.UpdateReservation)
			protected.DELETE("/reservations/:id", middleware.RequireRoles("Admin"), handlers.DeleteReservation)

			// ---- Locations ----
			protected.GET("/locations", handlers.GetLocations)
			protected.GET("/locations/:id", handlers.GetLocationByID)
			protected.POST("/locations", middleware.RequireRoles("Admin", "Asset Manager"), handlers.CreateLocation)
			protected.PUT("/locations/:id", middleware.RequireRoles("Admin", "Asset Manager"), handlers.UpdateLocation)
			protected.DELETE("/locations/:id", middleware.RequireRoles("Admin"), handlers.DeleteLocation)

			// ---- Financial ----
			financial := protected.Group("/financial")
			{
				financial.GET("/ledgers", handlers.GetLedgers)
				financial.GET("/ledgers/:id", handlers.GetLedgerByID)
				financial.POST("/ledgers", middleware.RequireRoles("Admin", "Finance Manager"), handlers.CreateLedger)
				financial.PUT("/ledgers/:id", middleware.RequireRoles("Admin", "Finance Manager"), handlers.UpdateLedger)
				financial.DELETE("/ledgers/:id", middleware.RequireRoles("Admin"), handlers.DeleteLedger)

				financial.GET("/leases", handlers.GetLeases)
				financial.GET("/leases/:id", handlers.GetLeaseByID)
				financial.POST("/leases", middleware.RequireRoles("Admin", "Finance Manager"), handlers.CreateLease)
				financial.PUT("/leases/:id", middleware.RequireRoles("Admin", "Finance Manager"), handlers.UpdateLease)
				financial.DELETE("/leases/:id", middleware.RequireRoles("Admin"), handlers.DeleteLease)

				financial.GET("/depreciations", handlers.GetDepreciations)
				financial.GET("/depreciations/:id", handlers.GetDepreciationByID)
				financial.POST("/depreciations", middleware.RequireRoles("Admin", "Finance Manager"), handlers.CreateDepreciation)
				financial.PUT("/depreciations/:id", middleware.RequireRoles("Admin", "Finance Manager"), handlers.UpdateDepreciation)
				financial.DELETE("/depreciations/:id", middleware.RequireRoles("Admin"), handlers.DeleteDepreciation)

				financial.GET("/software-licenses", handlers.GetSoftwareLicenses)
				financial.GET("/software-licenses/:id", handlers.GetSoftwareLicenseByID)
				financial.POST("/software-licenses", middleware.RequireRoles("Admin", "Finance Manager"), handlers.CreateSoftwareLicense)
				financial.PUT("/software-licenses/:id", middleware.RequireRoles("Admin", "Finance Manager"), handlers.UpdateSoftwareLicense)
				financial.DELETE("/software-licenses/:id", middleware.RequireRoles("Admin"), handlers.DeleteSoftwareLicense)
			}

			// ---- User Management ----
			protected.GET("/users", handlers.GetUsers)
			protected.GET("/users/:id", handlers.GetUserByID)
			protected.POST("/users", middleware.RequireRoles("Admin"), handlers.CreateUser)
			protected.PUT("/users/:id", middleware.RequireRoles("Admin"), handlers.UpdateUser)
			protected.DELETE("/users/:id", middleware.RequireRoles("Admin"), handlers.DeleteUser)

			protected.GET("/roles", handlers.GetRoles)
			protected.GET("/roles/:id", handlers.GetRoleByID)
			protected.POST("/roles", middleware.RequireRoles("Admin"), handlers.CreateRole)
			protected.PUT("/roles/:id", middleware.RequireRoles("Admin"), handlers.UpdateRole)
			protected.DELETE("/roles/:id", middleware.RequireRoles("Admin"), handlers.DeleteRole)

			// ---- Alerts ----
			protected.GET("/alerts", handlers.GetAlerts)
			protected.GET("/alerts/:id", handlers.GetAlertByID)
			protected.PUT("/alerts/:id", handlers.UpdateAlert)
			protected.DELETE("/alerts/:id", middleware.RequireRoles("Admin"), handlers.DeleteAlert)
			protected.POST("/alerts/mark-all-read", handlers.MarkAllAlertsRead)
			protected.POST("/alerts/generate", middleware.RequireRoles("Admin"), handlers.GenerateAlerts)
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
