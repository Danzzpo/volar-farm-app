package models

import "time"

type Animal struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	
	// --- Identitas Utama ---
	RingNumber string   `json:"ring_number"`
	Species    string   `json:"species"`
	Gender     string   `json:"gender"`
	Visual     string   `json:"visual"`
	
	// --- Silsilah (Pedigree) ---
	// Pilihan 1: Indukan ada di database
	SireID     *uint    `json:"sire_id"`
	DamID      *uint    `json:"dam_id"`
	Sire       *Animal  `json:"sire" gorm:"foreignKey:SireID"` 
	Dam        *Animal  `json:"dam" gorm:"foreignKey:DamID"`

	// Pilihan 2: Indukan dari luar (Input Manual)
	SireOther  string   `json:"sire_other"` 
	DamOther   string   `json:"dam_other"`

	// --- Status ---
	Status     string   `json:"status"` // AVAILABLE, SOLD
	Notes      string   `json:"notes"`
	
	// Tanggal Telur & Menetas DIHAPUS (Pindah ke fitur Breeding nanti)

	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
