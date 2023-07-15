package main

import (
	"log"
	"net/http"

	"github.com/goRestFullAPI/api/handlers"
	"github.com/goRestFullAPI/config"
	"github.com/goRestFullAPI/database"
	"github.com/gorilla/mux"
)

func main() {
	// Membaca konfigurasi dari file config/config.json
	config.ReadConfig()

	// Membuka koneksi ke database
	db := database.OpenDB()
	defer db.Close()

	// Membuat router
	router := mux.NewRouter()

	// Menambahkan routing endpoint
	router.HandleFunc("/cakes", handlers.GetCakesHandler).Methods("GET")
	router.HandleFunc("/cakes/{id}", handlers.GetCakeByIDHandler).Methods("GET")
	router.HandleFunc("/cakes", handlers.AddCakeHandler).Methods("POST")
	router.HandleFunc("/cakes/{id}", handlers.UpdateCakeHandler).Methods("PUT")
	router.HandleFunc("/cakes/{id}", handlers.DeleteCakeHandler).Methods("DELETE")

	// Menjalankan server HTTP
	log.Fatal(http.ListenAndServe(":8080", router))
}
