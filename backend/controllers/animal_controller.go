package controllers

import (
	"math"
	"net/http"
	"strconv" // Penting untuk konversi halaman
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
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.Query("search")
	status := c.Query("status")
	gender := c.Query("gender") // <--- 1. TANGKAP GENDER

	// Konversi string ke int
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	var animals []models.Animal
	var total int64

	// Query Dasar
	query := config.DB.Model(&models.Animal{}).Where("user_id = ?", userID)

	// Filter Search
	if search != "" {
		query = query.Where("(ring_number LIKE ? OR species LIKE ? OR visual LIKE ?)", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Filter Status
	if status != "" && status != "All" {
		query = query.Where("status = ?", status)
	}

	// Filter Gender (M/F)
	if gender != "" { // <--- 2. LOGIKA FILTER GENDER
		query = query.Where("gender = ?", gender)
	}

	// === BAGIAN PENTING YANG TIDAK BOLEH HILANG ===
	
	// 3. Hitung Total Data (Untuk Pagination)
	query.Count(&total)

	// 4. Ambil Data (Eksekusi Query)
	// Wajib pakai Preload agar data Bapak/Ibu terbawa
	if err := query.Limit(limit).Offset(offset).Preload("Sire").Preload("Dam").Order("created_at DESC").Find(&animals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	// ==============================================

	c.JSON(http.StatusOK, gin.H{
		"data":  animals,
		"total": total,
		"page":  page,
		"limit": limit,
		"pages": int(math.Ceil(float64(total) / float64(limit))),
	})
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

// 4. AMBIL STATISTIK FARM (LENGKAP)
func GetFarmStats(c *gin.Context) {
	userID := c.Query("user_id")

	var totalBirds int64
	var totalSold int64
	var activePairs int64
	var incubatingEggs int64

	// 1. Hitung Stok Burung
	config.DB.Model(&models.Animal{}).Where("user_id = ? AND status = ?", userID, "AVAILABLE").Count(&totalBirds)
	
	// 2. Hitung Terjual
	config.DB.Model(&models.Animal{}).Where("user_id = ? AND status = ?", userID, "SOLD").Count(&totalSold)

	// 3. Hitung Pasangan
	config.DB.Model(&models.Pair{}).Where("user_id = ?", userID).Count(&activePairs)

	// 4. Hitung Pengeraman
	config.DB.Model(&models.Incubation{}).Where("user_id = ?", userID).Count(&incubatingEggs)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"total_birds":     totalBirds,
			"total_sold":      totalSold,
			"active_pairs":    activePairs,    
			"incubating_eggs": incubatingEggs, 
		},
	})
}

// 5. UPDATE DATA BURUNG (Edit)
func UpdateAnimal(c *gin.Context) {
	id := c.Param("id") 
	var input models.Animal

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var animal models.Animal
	if err := config.DB.First(&animal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	config.DB.Model(&animal).Updates(input)

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate!", "data": animal})
}

// 6. HAPUS DATA BURUNG (Delete)
func DeleteAnimal(c *gin.Context) {
	id := c.Param("id") 
	
	if err := config.DB.Delete(&models.Animal{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus!"})
}