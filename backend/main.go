package main

import (
	"log"
	"net/http"

	"backend/config"
	"backend/routes"
	"github.com/rs/cors"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("DB error:", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	routes.SetupRoutes(mux, db)

	// Tambahkan ini agar preflight (OPTIONS) tidak ditolak
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		http.NotFound(w, r)
	})

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// Konfigurasi CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3333",
			"http://localhost:3000",
			"http://frontend:5173",
			"http://admin:80",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Range"},
        ExposedHeaders:   []string{"Content-Range"}, 
		AllowCredentials: true,
		Debug:            true, // aktifkan untuk lihat log CORS di terminal
	})

	handler := corsHandler.Handler(mux)

	log.Println("Routes sudah disetup, backend jalan di :4000")
	log.Fatal(http.ListenAndServe(":4000", handler))
}
