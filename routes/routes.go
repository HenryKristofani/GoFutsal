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
	r.GET("/courts", controllers.GetCourts)
	r.GET("/courts/:id", controllers.GetCourtByID)
	r.POST("/courts", controllers.CreateCourt)
	r.PUT("/courts/:id", controllers.UpdateCourt)
	r.DELETE("/courts/:id", controllers.DeleteCourt)

}