package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/earn/goRestFullAPI/api/models"
	"github.com/stretchr/testify/assert"
)

func TestAddCakeHandler(t *testing.T) {
	// Persiapkan request body
	cake := models.Cake{
		Title:       "Lemon cheesecake",
		Description: "A cheesecake made of lemon",
		Rating:      7,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

	body, err := json.Marshal(cake)
	assert.NoError(t, err)

	// Buat HTTP request dengan method POST dan body
	req, err := http.NewRequest("POST", "/cakes", bytes.NewReader(body))
	assert.NoError(t, err)

	// Buat HTTP response recorder untuk menangkap respons dari handler
	recorder := httptest.NewRecorder()

	// Panggil handler AddCakeHandler dengan HTTP request dan response recorder
	AddCakeHandler(recorder, req)

	// Periksa status code yang diharapkan
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Dekode response body
	var addedCake models.Cake
	err = json.NewDecoder(recorder.Body).Decode(&addedCake)
	assert.NoError(t, err)

	// Periksa apakah cake telah ditambahkan dengan benar
	assert.NotZero(t, addedCake.ID)
	assert.Equal(t, cake.Title, addedCake.Title)
	assert.Equal(t, cake.Description, addedCake.Description)
	assert.Equal(t, cake.Rating, addedCake.Rating)
	assert.Equal(t, cake.Image, addedCake.Image)
}

// Implementasikan unit tests untuk endpoint lainnya (GetCakesHandler, GetCakeByIDHandler, UpdateCakeHandler, DeleteCakeHandler)
// ...

func TestMain(m *testing.M) {
	// Setup sebelum menjalankan unit tests
	// ...

	// Menjalankan unit tests
	m.Run()

	// Cleanup setelah menjalankan unit tests
	// ...
}
