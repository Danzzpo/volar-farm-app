package config

import (
	"fmt"
	"log"
	"volar-farm-backend/models" // <--- PENTING: Import folder models agar tabel terbaca

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Konfigurasi koneksi (Username: root, Password: kosong, DB: db_volar_farm)
	// Pastikan database "db_volar_farm" sudah dibuat di phpMyAdmin
	dsn := "root:@tcp(127.0.0.1:3306)/db_volar_farm?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Gagal konek ke database! Pastikan XAMPP nyala dan nama DB benar.")
	}

	fmt.Println("✅ Berhasil terhubung ke Database MySQL!")

	// ==========================================
	// AUTO MIGRATE (MEMBUAT TABEL OTOMATIS)
	// ==========================================
	// GORM akan otomatis membuat atau memperbarui tabel berdasarkan struct di folder models
	err = DB.AutoMigrate(
		&models.User{},        // Tabel Users (Login/Register)
		&models.Animal{},      // Tabel Animals (Stok Burung)
		&models.Pair{},        // Tabel Pairs (Jodohan)
		&models.Incubation{},  // Tabel Incubations (Pengeraman)
		&models.Transaction{}, // Tabel Transactions (Keuangan/Finance)
		&models.Treatment{},	// TRitmen <--- BARU
	)

	if err != nil {
		log.Fatal("❌ Gagal migrasi database: ", err)
	}

	fmt.Println("🚀 Database Migrated Successfully (Tabel Lengkap)!")
}