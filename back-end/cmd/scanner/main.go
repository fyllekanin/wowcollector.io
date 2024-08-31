package main

import (
	"context"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"wowcollector.io/cmd/scanner/tasks"
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

	c := cron.New()

	time.AfterFunc(1*time.Hour, func() {
		tasks.ScanRealms("eu")
		c.AddFunc("@every 24h", func() {
			tasks.ScanRealms("eu")
		})
	})

	time.AfterFunc(2*time.Hour, func() {
		tasks.ScanMounts("eu")
		c.AddFunc("@every 12h", func() {
			tasks.ScanMounts("eu")
		})
	})

	time.AfterFunc(4*time.Hour, func() {
		tasks.ScanAchievements("eu")
		c.AddFunc("@every 24h", func() {
			tasks.ScanAchievements("eu")
		})
	})

	time.AfterFunc(6*time.Hour, func() {
		tasks.ScanToys("eu")
		c.AddFunc("@every 24h", func() {
			tasks.ScanToys("eu")
		})
	})

	time.AfterFunc(8*time.Hour, func() {
		tasks.ScanPets("eu")
		c.AddFunc("@every 24h", func() {
			tasks.ScanPets("eu")
		})
	})

	c.Start()

	select {}
}
