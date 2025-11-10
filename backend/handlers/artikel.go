package handlers

import (
	"backend/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// List & Create Artikel
func Artikel(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			// === List artikel untuk React-Admin ===
			listArtikel, err := models.GetAllArtikel(db)
			if err != nil {
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}

			total := len(listArtikel)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Range", fmt.Sprintf("artikel 0-%d/%d", total-1, total))
			w.Header().Set("Access-Control-Expose-Headers", "Content-Range")
			w.Header().Set("Access-Control-Allow-Origin", "*")

			// React-Admin expects { data: [...] } for list
			json.NewEncoder(w).Encode(map[string]interface{}{"data": listArtikel})

		case http.MethodPost:
			// === Create artikel ===
			var artikel models.Artikel
			if err := json.NewDecoder(r.Body).Decode(&artikel); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			id, err := models.InsertArtikel(db, artikel)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			artikel.ID = id
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")

			json.NewEncoder(w).Encode(map[string]interface{}{"data": artikel})

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

// Get Artikel by Slug (untuk FE user)
func GetArtikelBySlug(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		slug := strings.TrimPrefix(r.URL.Path, "/api/artikel/slug/")
		if slug == "" {
			http.Error(w, "Slug not found", http.StatusBadRequest)
			return
		}

		artikel, err := models.GetArtikelBySlug(db, slug)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Artikel tidak ditemukan", http.StatusNotFound)
				return
			}
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(artikel)
	}
}

// GetOne / Update / Delete Artikel by ID (untuk admin)
func ArtikelById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/api/artikel/id/")
		if id == "" {
			http.Error(w, "Id not found", http.StatusBadRequest)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Range")

		switch r.Method {
		case http.MethodGet:
			artikel, err := models.GetArtikelById(db, id)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Error(w, "Artikel tidak ditemukan", http.StatusNotFound)
					return
				}
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{"data": artikel})

		case http.MethodPut:
			var artikel models.Artikel
			if err := json.NewDecoder(r.Body).Decode(&artikel); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err := models.UpdateArtikelById(db, id, artikel)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{"data": artikel})

		case http.MethodDelete:
			err := models.DeleteArtikelById(db, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// React-Admin expects { data: previousData } for delete
			json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]string{"id": id}})

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
