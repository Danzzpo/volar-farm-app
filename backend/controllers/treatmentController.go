package controllers

import (
	"net/http"
	"volar-farm-backend/config"
	"volar-farm-backend/models"

	"github.com/gin-gonic/gin"
)

// 1. TAMBAH PERAWATAN
func CreateTreatment(c *gin.Context) {
	var input models.Treatment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data perawatan tersimpan!", "data": input})
}

// 2. LIHAT LIST PERAWATAN
func GetTreatments(c *gin.Context) {
	userID := c.Query("user_id")
	var treatments []models.Treatment

	// Urutkan dari yang terbaru (Date DESC)
	if err := config.DB.Where("user_id = ?", userID).Order("date desc").Find(&treatments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatments})
}

// 3. HAPUS PERAWATAN
func DeleteTreatment(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Treatment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Terhapus!"})
}