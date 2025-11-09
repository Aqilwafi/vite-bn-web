package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	var dsn string

	// Jika Railway inject DATABASE_URL, pakai langsung
	if os.Getenv("DATABASE_URL") != "" {
		dsn = os.Getenv("DATABASE_URL")
		fmt.Println("🔗 Using DATABASE_URL from environment")
	} else {
		// Jika lokal, pakai manual setup
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

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
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
