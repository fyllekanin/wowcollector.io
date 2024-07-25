package seeds

import (
	"encoding/json"

	"go.uber.org/zap"
	"wowcollector.io/entities/documents"
	realmrepository "wowcollector.io/repository/repositories/realm-repository"
	seedrepository "wowcollector.io/repository/repositories/seed-repository"
)

func RealmsSeeder() {
	var seedName = "realms"
	zap.L().Info("Realms seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		zap.L().Info("Realms seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/realms.json")

	var realms []documents.RealmDocument
	err = json.Unmarshal(byteValue, &realms)
	if err != nil {
		zap.L().Fatal("Error parsing realms.json")
	}

	for _, element := range realms {
		realmrepository.GetRepository().CreateRealm(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	zap.L().Info("Realms seeder done")
}
