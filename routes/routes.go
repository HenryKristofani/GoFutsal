package routes

import (
	"net/http"

	"github.com/HenryKristofani/GoFutsal/config"
	"github.com/HenryKristofani/GoFutsal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// USER CRUD
	r.GET("/api/users", controllers.GetUsers)
	r.GET("/api/users/:id", controllers.GetUserByID)
	r.PUT("/api/users/:id", controllers.UpdateUser)
	r.DELETE("/api/users/:id", controllers.DeleteUser)
	
	// AUTHENTICATION
	r.POST("/api/login", controllers.Login)
	
	// Route test API
	r.GET("/api", controllers.TestEndpoint)

	// 🔹 Route cek koneksi database
	r.GET("/dbcheck", func(c *gin.Context) {
		var version string
		err := config.DB.QueryRow("SELECT version();").Scan(&version)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"postgres_version": version})
	})

	// COURT routes
	r.GET("/api/courts", controllers.GetCourts)
	r.GET("/api/courts/:id", controllers.GetCourtByID)
	r.POST("/api/courts", controllers.CreateCourt)
	r.PUT("/api/courts/:id", controllers.UpdateCourt)
	r.DELETE("/api/courts/:id", controllers.DeleteCourt)

	// USER register (client only)
	r.POST("/api/users/register", controllers.RegisterUser)

	// BOOKING routes
	r.GET("/api/bookings", controllers.GetBookings)
	r.POST("/api/bookings", controllers.CreateBooking)
	r.GET("/api/bookings/:id", controllers.GetBookingByID)
	r.PUT("/api/bookings/:id", controllers.UpdateBooking)
	r.DELETE("/api/bookings/:id", controllers.DeleteBooking)

}
