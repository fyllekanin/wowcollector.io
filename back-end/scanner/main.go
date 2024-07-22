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
	"wowcollector.io/common/data"
	"wowcollector.io/repository"
	"wowcollector.io/scanner/tasks"
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
		tasks.ScanRealms(data.REGION_EU)
		c.AddFunc("@every 2h", func() {
			tasks.ScanRealms(data.REGION_EU)
		})
	})

	c.Start()

	select {}
}
