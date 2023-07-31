package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAlbums(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response []album
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	expectedAlbums := albums
	if len(response) != len(expectedAlbums) {
		t.Errorf("Expected %d albums, but got %d", len(expectedAlbums), len(response))
	}

	for i, album := range response {
		if album != expectedAlbums[i] {
			t.Errorf("Expected album %v, but got %v", expectedAlbums[i], album)
		}
	}
}

func TestGetAlbumByID(t *testing.T) {
	router := setupRouter()
	id := "1"
	req, _ := http.NewRequest("GET", "/albums/"+id, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response album
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	expectedAlbum := albums[0]
	if response != expectedAlbum {
		t.Errorf("Expected album %v, but got %v", expectedAlbum, response)
	}
}

func TestPostAlbums(t *testing.T) {
	router := setupRouter()

	newAlbum := album{ID: "4", Title: "New Album", Artist: "New Artist", Price: 25.99}
	jsonData, _ := json.Marshal(newAlbum)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

	var response album
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	if response != newAlbum {
		t.Errorf("Expected album %v, but got %v", newAlbum, response)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	return router
}
