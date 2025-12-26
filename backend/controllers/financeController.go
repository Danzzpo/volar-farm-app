package controllers

import (
	"net/http"
	"volar-farm-backend/config"
	"volar-farm-backend/models"

	"github.com/gin-gonic/gin"
)

type FinanceSummary struct {
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	Balance      float64 `json:"balance"`
}

// 1. RINGKASAN SALDO
func GetFinanceSummary(c *gin.Context) {
	userID := c.Query("user_id")
	var summary FinanceSummary

	config.DB.Model(&models.Transaction{}).Where("user_id = ? AND type = ?", userID, "Income").Select("COALESCE(SUM(amount), 0)").Scan(&summary.TotalIncome)
	config.DB.Model(&models.Transaction{}).Where("user_id = ? AND type = ?", userID, "Expense").Select("COALESCE(SUM(amount), 0)").Scan(&summary.TotalExpense)

	summary.Balance = summary.TotalIncome - summary.TotalExpense
	c.JSON(http.StatusOK, gin.H{"data": summary})
}

// 2. LIST TRANSAKSI
func GetTransactions(c *gin.Context) {
	userID := c.Query("user_id")
	var transactions []models.Transaction
	if err := config.DB.Where("user_id = ?", userID).Order("date desc").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

// 3. TAMBAH TRANSAKSI
func CreateTransaction(c *gin.Context) {
	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil!", "data": input})
}

// 4. HAPUS TRANSAKSI
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Transaction{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Terhapus!"})
}