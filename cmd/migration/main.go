package main

import (
	"fmt"
	"os"

	"github.com/fkihai/payflow/internal/infrastructure/config"
	"github.com/fkihai/payflow/internal/infrastructure/db/postgres"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		return
	}
	conn := postgres.NewPostgresConnection(&cfg.Database)

	db, err := conn.Connect()
	if err != nil {
		fmt.Printf("cannot connect db, %v\n", err)
		return
	}
	defer db.Close()

	migrate := postgres.NewPostgresMigrations(db)

	if len(os.Args) < 3 {
		fmt.Print("use --migrate -up or -down")
		return
	}

	switch os.Args[1] {
	case "migrate":
		if os.Args[2] == "up" {
			migrate.MigrateUp()
			break
		} else if os.Args[2] == "down" {
			migrate.MigrateDown()
			break
		}
	}

}
