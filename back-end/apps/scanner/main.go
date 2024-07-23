package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"wowcollector.io/apps/scanner/tasks"
	blizzarddata "wowcollector.io/common/data/blizzard-data"
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
