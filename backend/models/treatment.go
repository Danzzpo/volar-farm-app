package models

import "time"

type Treatment struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `json:"user_id"`
	Type      string     `json:"type"`      // Vitamin, Obat, Vaksin, Lainnya
	Name      string     `json:"name"`      // Nama obatnya (misal: Super N)
	Date      time.Time  `json:"date"`      // Tanggal pemberian
	NextDate  *time.Time `json:"next_date"` // Jadwal berikutnya (Opsional)
	Note      string     `json:"note"`      // Catatan tambahan
	CreatedAt time.Time  `json:"created_at"`
}