package main

import (
	"fmt"
	"gopos/internal/config"
	"gopos/internal/router"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	fmt.Println("test")
	// Inisialisasi koneksi database
	db := config.ConnectDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get raw DB from GORM:", err)
	}
	defer sqlDB.Close()

	// Inisialisasi router Gin
	r := gin.Default()

	// Middleware CORS
	r.Use(cors.Default())
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Load semua route dari router.go
	router.LoadRoutes(r, db)

	// Jalankan server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
