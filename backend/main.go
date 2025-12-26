package main

import (
	"volar-farm-backend/config"
	"volar-farm-backend/controllers"
	// "volar-farm-backend/models" // Tidak dipakai di sini, sudah aman dihapus

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Konek Database & Otomatis Migrate
	config.ConnectDatabase()

	r := gin.Default()

	// 2. Setup CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 3. Route API
	api := r.Group("/api")
	{
		// Auth
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Animals (Stok Burung)
		api.POST("/animals", controllers.CreateAnimal)       // Tambah Baru
		api.GET("/animals", controllers.GetMyAnimals)        // Lihat Data
		api.GET("/public/animals", controllers.GetPublicAnimals)
		
		// === TAMBAHAN BARU (UPDATE & DELETE) ===
		api.PUT("/animals/:id", controllers.UpdateAnimal)    // Edit Data
		api.DELETE("/animals/:id", controllers.DeleteAnimal) // Hapus Data

		// Pairs (Jodohan)
		api.POST("/pairs", controllers.CreatePair)
		api.GET("/pairs", controllers.GetMyPairs)
		api.PATCH("/pairs/:id", controllers.UpdatePairStatus)

		// Incubating (Pengeraman)
		api.GET("/pairs/active", controllers.GetActivePairs)
		api.GET("/incubations", controllers.GetIncubations)
		api.POST("/incubations", controllers.CreateIncubation)

		// Stats
		api.GET("/stats", controllers.GetFarmStats)

		// FINANCE (KEUANGAN)
		api.GET("/finance/summary", controllers.GetFinanceSummary)
		api.GET("/finance/transactions", controllers.GetTransactions)
		api.POST("/finance/transactions", controllers.CreateTransaction)
		api.DELETE("/finance/transactions/:id", controllers.DeleteTransaction)
	}

	r.Run(":8080")
}