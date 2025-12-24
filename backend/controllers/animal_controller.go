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
func GetMyAnimals(c *gin.Context) {
	userID := c.Query("user_id")

	var animals []models.Animal

	// Mengambil data burung milik user tersebut
	if err := config.DB.Where("user_id = ?", userID).Find(&animals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": animals})
}

// 3. AMBIL DATA PUBLIK (Public - Home Page)
func GetPublicAnimals(c *gin.Context) {
	var animals []models.Animal

	// Ambil semua hewan yang statusnya 'AVAILABLE'
	if err := config.DB.Where("status = ?", "AVAILABLE").Find(&animals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": animals})
}

// 4. AMBIL STATISTIK FARM (Opsional - Jika masih dipakai)
func GetFarmStats(c *gin.Context) {
	userID := c.Query("user_id")

	var totalBirds int64
	var totalSold int64

	config.DB.Model(&models.Animal{}).Where("user_id = ? AND status = ?", userID, "AVAILABLE").Count(&totalBirds)
	config.DB.Model(&models.Animal{}).Where("user_id = ? AND status = ?", userID, "SOLD").Count(&totalSold)

	c.JSON(http.StatusOK, gin.H{
		"total_birds": totalBirds,
		"total_sold":  totalSold,
	})
}

// ==========================================
// TAMBAHAN BARU UNTUK EDIT & HAPUS
// ==========================================

// 5. UPDATE DATA BURUNG (Edit)
func UpdateAnimal(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari URL
	var input models.Animal

	// Validasi input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari data lama berdasarkan ID
	var animal models.Animal
	if err := config.DB.First(&animal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	// Update data di database
	config.DB.Model(&animal).Updates(input)

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate!", "data": animal})
}

// 6. HAPUS DATA BURUNG (Delete)
func DeleteAnimal(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari URL
	
	// Hapus data berdasarkan ID
	if err := config.DB.Delete(&models.Animal{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus!"})
}