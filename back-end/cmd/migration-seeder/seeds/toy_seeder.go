package seeds

import (
	"encoding/json"

	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	seedrepository "wowcollector.io/internal/repository/repositories/seed-repository"
	toyrepository "wowcollector.io/internal/repository/repositories/toy-repository"
)

func ToysSeeder() {
	var seedName = "toys"
	zap.L().Info("Toys seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		zap.L().Info("Toys seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/seeds/toys.json")

	var toys []documents.ToyDocument
	err = json.Unmarshal(byteValue, &toys)
	if err != nil {
		zap.L().Fatal("Error parsing toys.json")
	}

	for _, element := range toys {
		toyrepository.GetRepository().CreateToy(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	zap.L().Info("Toys seeder done")
}
