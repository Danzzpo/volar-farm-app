package controllers

import (
	"net/http"
	"volar-farm-backend/config"
	"volar-farm-backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ==========================
// 1. BAGIAN REGISTER
// ==========================

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FarmName string `json:"farm_name" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	// Cek input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}

	// Simpan User
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
		FarmName: input.FarmName,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal daftar! Username atau Email mungkin sudah dipakai."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil! Silakan login."})
} 

// ==========================
// 2. BAGIAN LOGIN
// ==========================

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	var user models.User

	// 1. Cek Input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Cari User berdasarkan Email
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email salah atau belum terdaftar!"})
		return
	}

	// 3. Cek Password (Bandingkan password input vs hash di DB)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password salah!"})
		return
	}

	// 4. Sukses! Kirim data user balik ke Frontend
	c.JSON(http.StatusOK, gin.H{
		"message":   "Login Berhasil",
		"user_id":   user.ID,
		"username":  user.Username,
		"farm_name": user.FarmName,
	})
}