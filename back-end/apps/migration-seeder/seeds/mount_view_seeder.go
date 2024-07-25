package seeds

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"wowcollector.io/entities/documents"
	mountviewrepository "wowcollector.io/repository/repositories/mount-view-repository"
	seedrepository "wowcollector.io/repository/repositories/seed-repository"
)

func MountViewsSeeder() {
	var seedName = "mount-views"
	zap.L().Info("Mount views seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		zap.L().Info("Mount views seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/mount-views.json")

	var mountViews []documents.MountViewDocument
	err = json.Unmarshal(byteValue, &mountViews)
	if err != nil {
		zap.L().Fatal("Error parsing mount-views.json")
	}

	for _, element := range mountViews {
		element.ObjectID = primitive.NewObjectID()
		mountviewrepository.GetRepository().CreateMountView(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	zap.L().Info("Mount views seeder done")
}
