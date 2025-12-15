package controllers

import (
	"net/http"
	"volar-farm-backend/config"
	"volar-farm-backend/models"

	"github.com/gin-gonic/gin"
)

// 1. TAMBAH DATA BURUNG (Create)
func CreateAnimal(c *gin.Context) {
	var animal models.Animal

	// Tangkap data JSON dari Frontend
	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan ke Database
	if err := config.DB.Create(&animal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data hewan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data hewan berhasil disimpan!", "data": animal})
}

// 2. LIHAT DAFTAR BURUNG SAYA (Private - Dashboard)
// Hanya menampilkan burung milik User yang sedang login
func GetMyAnimals(c *gin.Context) {
	userID := c.Query("user_id")

	var animals []models.Animal

	// Tambahkan Preload("Sire") dan Preload("Dam")
	// Artinya: "Tolong ambilkan juga data Bapak dan Ibunya sekalian"
	if err := config.DB.Preload("Sire").Preload("Dam").Where("user_id = ?", userID).Find(&animals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": animals})
}

// 3. AMBIL DATA PUBLIK (Public - Home Page)
// Menampilkan SEMUA burung yang statusnya AVAILABLE (tanpa peduli pemiliknya)
func GetPublicAnimals(c *gin.Context) {
	var animals []models.Animal

	// Ambil semua hewan yang statusnya 'AVAILABLE'
	if err := config.DB.Preload("Sire").Preload("Dam").Where("status = ?", "AVAILABLE").Find(&animals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": animals})
}

// 4. AMBIL STATISTIK FARM (Dashboard)
func GetFarmStats(c *gin.Context) {
	userID := c.Query("user_id")

	var totalBirds int64
	var totalSold int64

	// Hitung Burung yang Masih Ada (AVAILABLE)
	config.DB.Model(&models.Animal{}).Where("user_id = ? AND status = ?", userID, "AVAILABLE").Count(&totalBirds)

	// Hitung Burung yang Terjual (SOLD)
	config.DB.Model(&models.Animal{}).Where("user_id = ? AND status = ?", userID, "SOLD").Count(&totalSold)

	c.JSON(http.StatusOK, gin.H{
		"total_birds": totalBirds,
		"total_sold":  totalSold,
	})
}