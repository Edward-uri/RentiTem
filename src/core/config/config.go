package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL     string
	Port      string
	UploadDir string
}

func Load() Config {
	_ = godotenv.Load(".env")
	_ = godotenv.Load("src/.env")
	_ = godotenv.Load("../.env")

	return Config{
		DBURL:     envOr("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/rentitems?sslmode=disable"),
		Port:      normalizePort(envOr("HTTP_PORT", "8080")),
		UploadDir: envOr("UPLOAD_DIR", "uploads"),
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func normalizePort(p string) string {
	if strings.HasPrefix(p, ":") {
		return p[1:]
	}
	return p
}
