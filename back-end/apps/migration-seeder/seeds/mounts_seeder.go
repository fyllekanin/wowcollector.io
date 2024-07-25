package seeds

import (
	"encoding/json"

	"go.uber.org/zap"
	"wowcollector.io/entities/documents"
	mountrepository "wowcollector.io/repository/repositories/mount-repository"
	seedrepository "wowcollector.io/repository/repositories/seed-repository"
)

func MountsSeeder() {
	var seedName = "mounts"
	zap.L().Info("Mounts seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		zap.L().Info("Mounts seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/mounts.json")

	var mounts []documents.MountDocument
	err = json.Unmarshal(byteValue, &mounts)
	if err != nil {
		zap.L().Fatal("Error parsing mounts.json")
	}

	for _, element := range mounts {
		mountrepository.GetRepository().CreateMount(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	zap.L().Info("Mounts seeder done")
}
