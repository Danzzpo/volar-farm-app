package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Type      string    `json:"type"`      // 'Income' atau 'Expense'
	Category  string    `json:"category"`  // Jual Burung, Pakan, Obat, dll
	Amount    float64   `json:"amount"`
	Date      time.Time `json:"date"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}