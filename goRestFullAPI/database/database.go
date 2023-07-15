package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/earn/goRestFullAPI/config"
)

// OpenDB membuka koneksi ke basis data
func OpenDB() *sql.DB {
	dbConfig := config.GetConfig().Database
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected to database")
	return db
}
