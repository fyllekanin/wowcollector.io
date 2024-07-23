package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"wowcollector.io/apps/scanner/tasks"
	blizzarddata "wowcollector.io/common/data/blizzard-data"
	"wowcollector.io/repository"
)

func main() {
	clientOptions := options.Client().ApplyURI(repository.GetDatabaseUri())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
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
