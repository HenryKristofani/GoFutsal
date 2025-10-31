package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/HenryKristofani/GoFutsal/config"
	"github.com/HenryKristofani/GoFutsal/controllers"
)

func SetupRoutes(r *gin.Engine) {
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

	// BOOKING routes
	r.GET("/api/bookings", controllers.GetBookings)
	r.POST("/api/bookings", controllers.CreateBooking)
	r.GET("/api/bookings/:id", controllers.GetBookingByID)
	r.PUT("/api/bookings/:id", controllers.UpdateBooking)
	r.DELETE("/api/bookings/:id", controllers.DeleteBooking)

}