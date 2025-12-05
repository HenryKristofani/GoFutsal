package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/HenryKristofani/GoFutsal/config"
	"github.com/HenryKristofani/GoFutsal/routes"

	// 👇 Swagger dependencies
	_ "github.com/HenryKristofani/GoFutsal/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @title GoFutsal API
// @version 1.0
// @description API untuk sistem booking lapangan futsal.
// @termsOfService http://swagger.io/terms/

// @contact.name Developer API Support
// @contact.url https://github.com/HenryKristofani
// @contact.email henry@example.com

// @host localhost:8080
// @BasePath /
func main() {
	// Load .env
	godotenv.Load()

	// Connect ke database
	config.ConnectDB()

	// Run database migrations
	config.CheckAndRunMigrations()

	// Inisialisasi Gin
	r := gin.Default()

	// Setup semua route dari folder routes/
	routes.SetupRoutes(r)

	// ✅ Tambahkan route Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Jalankan server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 GoFutsal server running on port %s\n", port)
	r.Run(":" + port)
}