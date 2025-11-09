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

	if os.Getenv("DATABASE_URL") != "" {
		dsn = os.Getenv("DATABASE_URL")

		// FIX: ubah prefix agar lib/pq bisa baca
		if strings.HasPrefix(dsn, "postgresql://") {
			dsn = strings.Replace(dsn, "postgresql://", "postgres://", 1)
		}

		fmt.Println("🔗 Using DATABASE_URL from environment")
	} else {
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
