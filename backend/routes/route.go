package routes

import (
	"database/sql"
	"net/http"
	"backend/handlers"
)

func SetupRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/api/artikel", handlers.GetArtikel(db))
	mux.HandleFunc("/api/artikel/", handlers.GetArtikelBySlug(db))
}
