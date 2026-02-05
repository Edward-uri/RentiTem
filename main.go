package main

import (
	"log"

	"main/src/core/config"
	coredb "main/src/core/db"
)

func main() {
	cfg := config.Load()

	database, err := coredb.New(cfg.DBURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	_ = database
}
