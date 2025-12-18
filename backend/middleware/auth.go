package middleware

import (
	"net/http"
	"strings"

	"github.com/HenryKristofani/GoFutsal/auth"
	"github.com/gin-gonic/gin"
)

// AuthRequired middleware untuk validasi JWT token
// @Summary JWT Authentication Middleware
// @Description Middleware untuk memvalidasi JWT token di header Authorization
// @Security BearerAuth
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authorization header required",
			})
			c.Abort()
			return
		}

		// Check Bearer format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid authorization format. Use 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token
		claims, err := auth.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid or expired token: " + err.Error(),
			})
			c.Abort()
			return
		}

		// Set user data in context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// AdminRequired middleware untuk memastikan user adalah admin
// @Summary Admin Role Middleware
// @Description Middleware untuk memvalidasi bahwa user memiliki role admin
// @Security BearerAuth
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Admin access required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// CORS middleware untuk menangani cross-origin requests
// @Summary CORS Middleware
// @Description Middleware untuk menangani CORS policy
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
