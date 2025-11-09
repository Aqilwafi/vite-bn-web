package main

import (
	"log"
	"net/http"
	"os"

	"backend/config"
	"backend/routes"
	"github.com/rs/cors"
)

func main() {
	// Koneksi Database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("DB error:", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	routes.SetupRoutes(mux, db)

	// File statis
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// Ambil PORT dari environment (Railway)
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000" // fallback lokal
	}

	// CORS setup
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",          // FE lokal
			"http://localhost:3333",          // Admin lokal
			"https://bnfe-production.up.railway.app", // FE di Railway
			"https://admin-panel.up.railway.app", // Admin di Railway
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Range"},
		ExposedHeaders:   []string{"Content-Range"},
		AllowCredentials: true,
		Debug:            false, // set true kalau mau debugging log CORS
	})

	handler := corsHandler.Handler(mux)

	log.Printf("✅ Backend jalan di port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
