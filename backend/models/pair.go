package models

import "time"

type Pair struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	
	// Data Pasangan
	MaleID    uint      `json:"male_id"`
	FemaleID  uint      `json:"female_id"`
	
	// Relasi (Supaya bisa ambil nama burungnya)
	Male      *Animal   `json:"male" gorm:"foreignKey:MaleID"`
	Female    *Animal   `json:"female" gorm:"foreignKey:FemaleID"`

	// Info Kandang
	Cage      string    `json:"cage"`       // Nama Kandang / Glodok
	PairDate  time.Time `json:"pair_date"`  // Tanggal Jodoh
	Status    string    `json:"status"`     // ACTIVE (Produksi), REST (Istirahat), HISTORY (Sudah Pisah)
	Notes     string    `json:"notes"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}