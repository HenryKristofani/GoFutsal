package routes

import (
	"net/http"

	"github.com/HenryKristofani/GoFutsal/config"
	"github.com/HenryKristofani/GoFutsal/controllers"
	"github.com/HenryKristofani/GoFutsal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Add CORS middleware
	r.Use(middleware.CORS())

	// Public API routes
	api := r.Group("/api")
	{
		// AUTHENTICATION (Public routes)
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/refresh", controllers.RefreshToken)
		}

		// PUBLIC ROUTES
		api.POST("/users/register", controllers.RegisterUser)

		// Public court info (can be viewed without auth)
		api.GET("/courts", controllers.GetCourts)
		api.GET("/courts/:id", controllers.GetCourtByID)
	}

	// Protected API routes (requires JWT)
	protected := api.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		// USER PROFILE
		protected.GET("/profile", controllers.GetProfile)

		// USER CRUD (protected)
		protected.GET("/users", controllers.GetUsers)
		protected.GET("/users/:id", controllers.GetUserByID)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)

		// BOOKING routes (user can manage their bookings)
		protected.GET("/bookings", controllers.GetBookings)
		protected.POST("/bookings", controllers.CreateBooking)
		protected.GET("/bookings/:id", controllers.GetBookingByID)
		protected.PUT("/bookings/:id", controllers.UpdateBooking)
		protected.DELETE("/bookings/:id", controllers.DeleteBooking)
	}

	// Admin routes (requires JWT + admin role)
	admin := protected.Group("/admin")
	admin.Use(middleware.AdminRequired())
	{
		// COURT MANAGEMENT (admin only)
		admin.POST("/courts", controllers.CreateCourt)
		admin.PUT("/courts/:id", controllers.UpdateCourt)
		admin.DELETE("/courts/:id", controllers.DeleteCourt)
	}

	// Health check and test endpoints (public)
	r.GET("/api", controllers.TestEndpoint)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "GoFutsal API is running",
			"version": "1.0.0",
		})
	})

	// Database check endpoint (public)
	r.GET("/dbcheck", func(c *gin.Context) {
		var version string
		err := config.DB.QueryRow("SELECT version();").Scan(&version)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"postgres_version": version})
	})
}
