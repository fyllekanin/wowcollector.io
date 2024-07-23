package main

import (
	"context"
	"log"
	"os"

	"wowcollector.io/apps/migration-seeder/seeds"
	"wowcollector.io/repository"
)

func main() {
	client := repository.GetDatabaseClient()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	repository.InitRepositories(client.Database(os.Getenv("DATABASE_NAME")))

	runMigrations()
	runSeeds()
}

func runMigrations() {
	// Empty
}

func runSeeds() {
	seeds.RealmsSeeder()
}
