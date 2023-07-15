package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/api/models"
	"github.com/earn/goRestFullAPI/api/repositories"
)

// GetCakesHandler menghandle permintaan GET ke endpoint /cakes
func GetCakesHandler(w http.ResponseWriter, r *http.Request) {
	cakes, err := repositories.GetCakes()
	if err != nil {
		log.Println("Failed to get cakes:", err)
		http.Error(w, "Failed to get cakes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cakes)
	if err != nil {
		log.Println("Failed to encode cakes:", err)
		http.Error(w, "Failed to encode cakes", http.StatusInternalServerError)
		return
	}
}

// GetCakeByIDHandler menghandle permintaan GET ke endpoint /cakes/{id}
func GetCakeByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan parameter ID dari URL
	id := r.URL.Query().Get("id")

	cake, err := repositories.GetCakeByID(id)
	if err != nil {
		log.Println("Failed to get cake:", err)
		http.Error(w, "Failed to get cake", http.StatusInternalServerError)
		return
	}

	if cake == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cake)
	if err != nil {
		log.Println("Failed to encode cake:", err)
		http.Error(w, "Failed to encode cake", http.StatusInternalServerError)
		return
	}
}

// AddCakeHandler menghandle permintaan POST ke endpoint /cakes
func AddCakeHandler(w http.ResponseWriter, r *http.Request) {
	var cake models.Cake
	err := json.NewDecoder(r.Body).Decode(&cake)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = cake.Validate()
	if err != nil {
		log.Println("Invalid cake data:", err)
		http.Error(w, "Invalid cake data", http.StatusBadRequest)
		return
	}

	addedCake, err := repositories.AddCake(&cake)
	if err != nil {
		log.Println("Failed to add cake:", err)
		http.Error(w, "Failed to add cake", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(addedCake)
	if err != nil {
		log.Println("Failed to encode added cake:", err)
		http.Error(w, "Failed to encode added cake", http.StatusInternalServerError)
		return
	}
}

// UpdateCakeHandler menghandle permintaan PUT ke endpoint /cakes/{id}
func UpdateCakeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var updatedCake models.Cake
	err := json.NewDecoder(r.Body).Decode(&updatedCake)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = updatedCake.Validate()
	if err != nil {
		log.Println("Invalid cake data:", err)
		http.Error(w, "Invalid cake data", http.StatusBadRequest)
		return
	}

	err = repositories.UpdateCake(&updatedCake)
	if err != nil {
		log.Println("Failed to update cake:", err)
		http.Error(w, "Failed to update cake", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteCakeHandler menghandle permintaan DELETE ke endpoint /cakes/{id}
func DeleteCakeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := repositories.DeleteCake(id)
	if err != nil {
		log.Println("Failed to delete cake:", err)
		http.Error(w, "Failed to delete cake", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
