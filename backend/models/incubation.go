package models

// Incubation merepresentasikan data pengeraman telur
type Incubation struct {
	ID         int    `json:"id"`
	PairID     int    `json:"pair_id"`
	
	// Field Tambahan untuk Tampilan (Hasil Join)
	PairVisual string `json:"pair_visual"` // Gabungan Nama Jantan x Betina
	Cage       string `json:"cage"`        // Kode Kandang
	
	StartDate  string `json:"start_date"`
	EggCount   int    `json:"egg_count"`
	Notes      string `json:"notes"`
	UserID     int    `json:"user_id"`
}