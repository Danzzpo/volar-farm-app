package controllers

import (
	"net/http"
	"volar-farm-backend/config"
	"volar-farm-backend/models"

	"github.com/gin-gonic/gin"
)

// 1. TAMBAH PASANGAN (Pairing)
func CreatePair(c *gin.Context) {
	var pair models.Pair

	if err := c.ShouldBindJSON(&pair); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Default status Active jika kosong
	if pair.Status == "" {
		pair.Status = "ACTIVE"
	}

	if err := config.DB.Create(&pair).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan pasangan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil menjodohkan!", "data": pair})
}

// 2. LIHAT PASANGAN SAYA
func GetMyPairs(c *gin.Context) {
	userID := c.Query("user_id")

	var pairs []models.Pair

	// Ambil data Pair + Data Burung Jantan & Betinanya (Preload)
	// Urutkan dari yang terbaru (descending)
	err := config.DB.Preload("Male").Preload("Female").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&pairs).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pairs})
}

// 3. UPDATE STATUS (Misal: Cabut Jodoh / Istirahat)
func UpdatePairStatus(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Status string `json:"status"` // ACTIVE, REST, HISTORY
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&models.Pair{}).Where("id = ?", id).Update("status", input.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status berhasil diupdate"})
}