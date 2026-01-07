package main

import (
	"log"
	"session-24/cmd"
	"session-24/internal/data/repository"
	"session-24/internal/wire"
	"session-24/pkg/database"
	"session-24/pkg/utils"
)

func main() {
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatalf("failed to read file config: %v", err)
	}

	db, err := database.InitDB(config.DB)
	if err != nil {
		log.Fatalf("failed to connect to postgres database: %v", err)
	}

	repo := repository.NewOrderRepository(db)

	app := wire.Wiring(repo)

	cmd.APiserver(app)
}
