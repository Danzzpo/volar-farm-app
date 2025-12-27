package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"volar-farm-backend/config"
	"volar-farm-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SESUAIKAN DENGAN KONEKSI DATABASE ANDA
const DSN = "root:@tcp(127.0.0.1:3306)/volar_farm?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	// 1. KONEKSI DB KHUSUS SEEDER
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal konek DB:", err)
	}
	config.DB = db

	// ============================================================
	// BAGIAN INI YANG DITAMBAHKAN (Agar Tabel Dibuat Otomatis)
	// ============================================================
	fmt.Println("🔄 Memastikan tabel tersedia...")
	err = db.AutoMigrate(&models.Animal{}, &models.Pair{})
	if err != nil {
		log.Fatal("❌ Gagal membuat tabel:", err)
	}
	// ============================================================

	// Seed Randomizer
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🌱 Memulai Seeding 1000 Burung & Poligami...")

	var gen1IDs []uint // ID Buyut
	var gen2IDs []uint // ID Indukan

	visuals := []string{"Green Series", "Blue Series", "Parblue", "Lutino", "Albino", "Biola", "Euwing"}

	// --- TAHAP 1: GENERASI 1 (BUYUT) - 100 Ekor ---
	fmt.Println("... Membuat 100 Buyut (Gen 1)")
	for i := 0; i < 100; i++ {
		gender := "Male"
		if i%2 != 0 {
			gender = "Female"
		}

		bird := models.Animal{
			UserID:     1,                                     // Asumsi User ID 1 (Admin)
			RingNumber: fmt.Sprintf("LB-GEN1-%03d", i+1),      // Mengganti Name
			Species:    "Lovebird",
			Gender:     gender, // Mengganti Sex
			Visual:     visuals[rand.Intn(len(visuals))],
			Status:     "SOLD",
			Notes:      "Generasi Buyut (Founder)",
		}
		db.Create(&bird)
		gen1IDs = append(gen1IDs, bird.ID)
	}

	// --- TAHAP 2: GENERASI 2 (INDUKAN) - 300 Ekor ---
	fmt.Println("... Membuat 300 Indukan (Gen 2 - Anak dari Gen 1)")
	for i := 0; i < 300; i++ {
		gender := "Male"
		if i%2 != 0 {
			gender = "Female"
		}

		// Ambil bapak & ibu acak dari Gen 1
		sireID := gen1IDs[rand.Intn(len(gen1IDs))]
		damID := gen1IDs[rand.Intn(len(gen1IDs))]

		bird := models.Animal{
			UserID:     1,
			RingNumber: fmt.Sprintf("LB-GEN2-%03d", i+1),
			Species:    "Lovebird",
			Gender:     gender,
			Visual:     visuals[rand.Intn(len(visuals))],
			Status:     "AVAILABLE",
			SireID:     &sireID, // Bapak
			DamID:      &damID,  // Ibu
			Notes:      "Generasi Indukan",
		}
		db.Create(&bird)
		gen2IDs = append(gen2IDs, bird.ID)
	}

	// --- TAHAP 3: GENERASI 3 (ANAKAN) - 600 Ekor ---
	fmt.Println("... Membuat 600 Anakan (Gen 3 - Cucu Gen 1)")
	for i := 0; i < 600; i++ {
		gender := "Unknown" // Anakan belum ketahuan kelaminnya
		if rand.Intn(10) > 7 {
			gender = "Male"
		} // Random dikit

		sireID := gen2IDs[rand.Intn(len(gen2IDs))]
		damID := gen2IDs[rand.Intn(len(gen2IDs))]

		bird := models.Animal{
			UserID:     1,
			RingNumber: fmt.Sprintf("LB-GEN3-%03d", i+1),
			Species:    "Lovebird",
			Gender:     gender,
			Visual:     visuals[rand.Intn(len(visuals))],
			Status:     "AVAILABLE",
			SireID:     &sireID,
			DamID:      &damID,
			Notes:      "Anakan hasil ternak sendiri",
		}
		db.Create(&bird)
	}

	// --- TAHAP 4: MEMBUAT PASANGAN & POLIGAMI ---
	fmt.Println("... Membuat Pasangan & Poligami")

	// Cari 1 Jantan dari Gen 2 (Untuk Poligami)
	var alphaMale models.Animal
	db.Where("gender = ? AND ring_number LIKE ?", "Male", "LB-GEN2%").First(&alphaMale)

	// Cari 3 Betina dari Gen 2
	var females []models.Animal
	db.Where("gender = ? AND ring_number LIKE ?", "Female", "LB-GEN2%").Limit(3).Find(&females)

	if alphaMale.ID != 0 && len(females) == 3 {
		// Buat Jantan ini poligami dengan 3 betina
		for i, female := range females {
			pair := models.Pair{
				UserID:   1,
				MaleID:   alphaMale.ID,
				FemaleID: female.ID,
				Cage:     fmt.Sprintf("Kandang Poligami %d", i+1),
				Status:   "Active",
				PairDate: time.Now(), // Mengganti StartDate
				Notes:    "Pasangan Poligami Test",
			}
			db.Create(&pair)
			fmt.Printf("   -> ❤️ Poligami: %s kawin dengan %s\n", alphaMale.RingNumber, female.RingNumber)
		}
	}

	// Buat 17 Pasangan Normal Sisanya (Total target 20)
	var otherMales []models.Animal
	var otherFemales []models.Animal

	// Ambil jantan & betina lain (kecuali yg sudah poligami)
	db.Where("gender = ? AND ring_number LIKE ? AND id != ?", "Male", "LB-GEN2%", alphaMale.ID).Limit(17).Find(&otherMales)
	db.Where("gender = ? AND ring_number LIKE ?", "Female", "LB-GEN2%").Offset(3).Limit(17).Find(&otherFemales)

	for i := 0; i < len(otherMales) && i < len(otherFemales); i++ {
		pair := models.Pair{
			UserID:   1,
			MaleID:   otherMales[i].ID,
			FemaleID: otherFemales[i].ID,
			Cage:     fmt.Sprintf("Kandang %d", i+1),
			Status:   "Active",
			PairDate: time.Now(),
		}
		db.Create(&pair)
	}

	fmt.Println("✅ SELESAI! Database terisi 1000 Burung & 20 Pasangan (Termasuk Poligami).")
}