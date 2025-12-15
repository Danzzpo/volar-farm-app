package main

import (
	"volar-farm-backend/config"
	"volar-farm-backend/controllers" // <-- Tambahkan ini
	"volar-farm-backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{}, &models.Animal{}, &models.Pair{})

	r := gin.Default()

	// Setup CORS (Agar Vue bisa akses)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// === ROUTE API ===
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.POST("/animals", controllers.CreateAnimal) // Buat nambah data
		api.GET("/animals", controllers.GetMyAnimals) 
		api.GET("/public/animals", controllers.GetPublicAnimals)
		api.GET("/stats", controllers.GetFarmStats)
		api.POST("/pairs", controllers.CreatePair)
		api.GET("/pairs", controllers.GetMyPairs)
		api.PATCH("/pairs/:id", controllers.UpdatePairStatus) // Buat ubah status nanti // <-- URL Pendaftaran
	}

	r.Run(":8080")
}
