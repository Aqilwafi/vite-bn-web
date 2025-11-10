package routes

import (
	"database/sql"
	"net/http"
	"backend/handlers"
)

func SetupRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/api/artikel", handlers.GetArtikel(db))
	mux.HandleFunc("/api/artikel/slug/", handlers.GetArtikelBySlug(db))
	mux.HandleFunc("/api/artikel/id/", handlers.ArtikelById(db))
	mux.HandleFunc("/api/artikel/create/", handlers.CreateArtikel(db))
}
