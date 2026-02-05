package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds cross-cutting configuration values.
type Config struct {
	DBURL string
	Port  string
}

// Load builds Config from environment variables.
func Load() Config {
	_ = godotenv.Load(".env")     // best-effort root .env
	_ = godotenv.Load("src/.env") // fallback when running from repo root
	_ = godotenv.Load("../.env")  // fallback when invoked inside src/

	return Config{
		DBURL: envOr("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/rentitems?sslmode=disable"),
		Port:  normalizePort(envOr("HTTP_PORT", "8080")),
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
