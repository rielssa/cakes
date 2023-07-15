package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config adalah struktur konfigurasi
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
}

// Config adalah variabel global yang menyimpan konfigurasi
var config Config

// ReadConfig membaca konfigurasi dari file config/config.json
func ReadConfig() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal("Failed to read config file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Failed to decode config file:", err)
	}
}
