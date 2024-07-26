package main

import (
	"context"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"wowcollector.io/cmd/scanner/tasks"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
	"wowcollector.io/internal/repository"
)

func main() {
	logger, _ := zap.NewProduction()
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
		tasks.ScanRealms(blizzarddata.REGION_EU)
		c.AddFunc("@every 24h", func() {
			tasks.ScanRealms(blizzarddata.REGION_EU)
		})
	})

	time.AfterFunc(2*time.Hour, func() {
		tasks.ScanMounts(blizzarddata.REGION_EU)
		c.AddFunc("@every 12h", func() {
			tasks.ScanMounts(blizzarddata.REGION_EU)
		})
	})

	c.Start()

	select {}
}
