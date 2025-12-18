package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/HenryKristofani/GoFutsal/config"
	"github.com/HenryKristofani/GoFutsal/routes"

	// 👇 Swagger dependencies
	_ "github.com/HenryKristofani/GoFutsal/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GoFutsal API
// @version 1.0
// @description API untuk sistem booking lapangan futsal dengan JWT Authentication.
// @termsOfService http://swagger.io/terms/

// @contact.name Developer API Support
// @contact.url https://github.com/HenryKristofani
// @contact.email henry@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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

	// Setup graceful shutdown
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Jalankan server di goroutine
	go func() {
		fmt.Printf("🚀 GoFutsal server running on port %s\n", port)
		fmt.Println("📝 Tekan Ctrl+C untuk shutdown gracefully")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("❌ Server gagal start: %s\n", err)
		}
	}()

	// Tunggu signal interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("🛑 Shutdown server...")

	// Graceful shutdown dengan timeout 5 detik
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("❌ Server forced shutdown: %s\n", err)
	}

	fmt.Println("✅ Server berhasil shutdown")
}
