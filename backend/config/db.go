package config

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	var dsn string

	dbURL := os.Getenv("DATABASE_URL")

	if dbURL != "" {
		// Railway biasanya pakai postgresql://
		if strings.HasPrefix(dbURL, "postgresql://") {
			dbURL = strings.Replace(dbURL, "postgresql://", "postgres://", 1)
		}

		// Tambahkan sslmode=require (Railway butuh SSL)
		if !strings.Contains(dbURL, "sslmode=") {
			if strings.Contains(dbURL, "?") {
				dbURL += "&sslmode=require"
			} else {
				dbURL += "?sslmode=require"
			}
		}

		dsn = dbURL
		fmt.Println("🔗 Using DATABASE_URL from Railway")
	} else {
		// fallback lokal
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "5432"),
			getEnv("DB_USER", "baca_user"),
			getEnv("DB_PASSWORD", "baca_pass"),
			getEnv("DB_NAME", "baca_db"),
		)
		fmt.Println("🧩 Using local DB configuration")
	}

	db, err := sql.Open("postgres", strings.TrimSpace(dsn))
	if err != nil {
		return nil, fmt.Errorf("DB open error: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("DB ping error: %v", err)
	}

	fmt.Println("✅ Database connected successfully!")
	return db, nil
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
