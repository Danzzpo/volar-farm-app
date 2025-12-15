package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Username: root, Password: (kosong), Database: db_volar_farm
	dsn := "root:@tcp(127.0.0.1:3306)/db_volar_farm?charset=utf8mb4&parseTime=True&loc=Local"
	
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal konek ke database! Pastikan XAMPP nyala.")
	}

	fmt.Println("✅ Berhasil terhubung ke Database MySQL!")
	DB = database
}