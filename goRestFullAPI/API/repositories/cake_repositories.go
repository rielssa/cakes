package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/goRestFullAPI/api/models"
	"github.com/goRestFullAPI/database"
)

// GetCakes mengambil daftar kue dari basis data
func GetCakes() ([]models.Cake, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM cakes")
	if err != nil {
		log.Println("Failed to get cakes from database:", err)
		return nil, err
	}
	defer rows.Close()

	cakes := []models.Cake{}
	for rows.Next() {
		var cake models.Cake
		err := rows.Scan(
			&cake.ID,
			&cake.Title,
			&cake.Description,
			&cake.Rating,
			&cake.Image,
			&cake.CreatedAt,
			&cake.UpdatedAt,
			&cake.DeletedAt,
		)
		if err != nil {
			log.Println("Failed to scan cake row:", err)
			continue
		}
		cakes = append(cakes, cake)
	}

	return cakes, nil
}

// GetCakeByID mengambil detail kue berdasarkan ID dari basis data
func GetCakeByID(id int) (*models.Cake, error) {
	db := database.GetDB()

	var cake models.Cake
	err := db.QueryRow("SELECT * FROM cakes WHERE id = ?", id).Scan(
		&cake.ID,
		&cake.Title,
		&cake.Description,
		&cake.Rating,
		&cake.Image,
		&cake.CreatedAt,
		&cake.UpdatedAt,
		&cake.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Tidak ada kue dengan ID tersebut
		}
		log.Println("Failed to get cake from database:", err)
		return nil, err
	}

	return &cake, nil
}

// AddCake menambahkan kue ke basis data
func AddCake(cake *models.Cake) (*models.Cake, error) {
	db := database.GetDB()

	cake.CreatedAt = time.Now()
	cake.UpdatedAt = time.Now()

	// Implementasikan logika untuk menambahkan kue ke basis data
	// ...

	return cake, nil
}

// UpdateCake memperbarui kue di basis data
func UpdateCake(cake *models.Cake) error {
	db := database.GetDB()

	cake.UpdatedAt = time.Now()

	// Implementasikan logika untuk memperbarui kue di basis data
	// ...

	return nil
}

// DeleteCake menghapus kue dari basis data
func DeleteCake(id int) error {
	db := database.GetDB()

	// Implementasikan logika untuk menghapus kue dari basis data
	// ...

	return nil
}
