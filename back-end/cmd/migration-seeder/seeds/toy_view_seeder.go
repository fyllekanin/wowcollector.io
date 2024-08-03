package seeds

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	seedrepository "wowcollector.io/internal/repository/repositories/seed-repository"
	toyviewrepository "wowcollector.io/internal/repository/repositories/toy-view-repository"
)

func ToyViewsSeeder() {
	var seedName = "toy-views"
	zap.L().Info("Toy views seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		zap.L().Info("Toy views seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/seeds/toy-views.json")

	var toyViews []documents.ToyViewDocument
	err = json.Unmarshal(byteValue, &toyViews)
	if err != nil {
		zap.L().Fatal("Error parsing toy-views.json")
	}

	for _, element := range toyViews {
		element.ObjectID = primitive.NewObjectID()
		toyviewrepository.GetRepository().CreateToyView(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	zap.L().Info("Toy views seeder done")
}
