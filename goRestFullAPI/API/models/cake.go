package models

import (
	"time"
)

// Cake adalah struktur data untuk kue
type Cake struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

// Validate melakukan validasi data kue
func (cake *Cake) Validate() error {
	// Implementasikan validasi sesuai kebutuhan
	return nil
}
