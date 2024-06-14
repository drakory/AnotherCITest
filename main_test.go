package main

import (
	"bookapi/entity"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAddition(t *testing.T){
	a := 3
	b := 2
	result := Addition(a,b)
	if (result != 5){
		t.Error("Invalid")
	}
}

func TestInsertBook(t *testing.T) {
	router := setupRouter()

	// Créez une requête POST
	w := httptest.NewRecorder()
	book := entity.Book{
		Title:       "Test Book",
		Year:        "2023",
		Description: "A book used for testing",
		BookCover:   "https://example.com/cover.jpg",
		UserID:      1, // Assurez-vous que cet utilisateur existe dans la base de données pour le test
	}
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/api/v1/book/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Envoyez la requête
	router.ServeHTTP(w, req)

	// Vérifiez la réponse
	assert.Equal(t, http.StatusOK, w.Code)
	// Vous pouvez aussi vérifier le corps de la réponse ici
}

/*func TestGetAllBooks(t *testing.T) {
	router := setupRouter()

	// Créez une requête GET
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/book/", nil)

	// Envoyez la requête
	router.ServeHTTP(w, req)

	// Vérifiez la réponse
	assert.Equal(t, http.StatusOK, w.Code)
	// Vous pouvez aussi vérifier le corps de la réponse ici
}

func TestGetBook(t *testing.T) {
	router := setupRouter()

	// Créez une requête GET
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/book/1", nil)

	// Envoyez la requête
	router.ServeHTTP(w, req)

	// Vérifiez la réponse
	assert.Equal(t, http.StatusOK, w.Code)
	// Vous pouvez aussi vérifier le corps de la réponse ici
}*/