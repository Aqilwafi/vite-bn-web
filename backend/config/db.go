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

		

		fmt.Println("🔗 Using DATABASE_URL from environment")
	} else {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			getEnv("PGHOST", "localhost"),
			getEnv("PGPORT", "5432"),
			getEnv("PGUSER", "baca_user"),
			getEnv("PGPASSWORD", "baca_pass"),
			getEnv("PGDATABASE", "baca_db"),
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
