package main

import (
	"context"
	"os"

	"go.uber.org/zap"
	"wowcollector.io/cmd/migration-seeder/seeds"
	"wowcollector.io/internal/repository"
)

func main() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
	client := repository.GetDatabaseClient()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			zap.L().Fatal(err.Error())
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
	seeds.MountViewsSeeder()
	seeds.MountsSeeder()
	seeds.AchievementsSeeder()
	seeds.AchievementCategoriesSeeder()
	seeds.ToysSeeder()
	seeds.ToyViewsSeeder()
	seeds.PetsSeeder()
	seeds.PetViewsSeeder()
}
