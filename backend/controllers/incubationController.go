package controllers

import (
	"net/http"
	"volar-farm-backend/config"
	"volar-farm-backend/models"

	"github.com/gin-gonic/gin"
)

// Struct untuk respons JSON yang rapi
type IncubationResponse struct {
	ID         uint   `json:"id"`
	PairID     uint   `json:"pair_id"`
	PairVisual string `json:"pair_visual"`
	Cage       string `json:"cage"`
	StartDate  string `json:"start_date"`
	EggCount   int    `json:"egg_count"`
	Status     string `json:"status"`
}

type ActivePairResponse struct {
	ID     uint   `json:"id"`
	Cage   string `json:"cage"`
	Visual string `json:"visual"`
}

// 1. GET: Ambil Data Incubating
func GetIncubations(c *gin.Context) {
	userID := c.Query("user_id")
	var results []IncubationResponse

	// Menggunakan Raw SQL via GORM untuk Join yang kompleks
	query := `
		SELECT 
			i.id, i.pair_id, i.start_date, i.egg_count, i.status,
			p.cage,
			CONCAT(m.visual, ' x ', f.visual) as pair_visual
		FROM incubations i
		JOIN pairs p ON i.pair_id = p.id
		LEFT JOIN animals m ON p.male_id = m.id
		LEFT JOIN animals f ON p.female_id = f.id
		WHERE i.user_id = ?
		ORDER BY i.id DESC
	`

	// Perbaikan: Scan hasil query ke struct response
	if err := config.DB.Raw(query, userID).Scan(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}

// 2. GET: Ambil Pasangan Aktif (Dropdown)
func GetActivePairs(c *gin.Context) {
	userID := c.Query("user_id")
	var results []ActivePairResponse

	query := `
		SELECT 
			p.id, p.cage,
			CONCAT(m.visual, ' x ', f.visual) as visual
		FROM pairs p
		LEFT JOIN animals m ON p.male_id = m.id
		LEFT JOIN animals f ON p.female_id = f.id
		WHERE p.user_id = ? AND p.status = 'Active'
	`

	if err := config.DB.Raw(query, userID).Scan(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}

// 3. POST: Simpan Data Baru
func CreateIncubation(c *gin.Context) {
	var input models.Incubation

	// Bind JSON dari request ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbaikan: Gunakan GORM Create (Lebih simpel & aman daripada Exec raw SQL)
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil input telur!", "data": input})
}