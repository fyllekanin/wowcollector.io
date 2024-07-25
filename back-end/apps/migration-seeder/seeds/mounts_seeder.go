package seeds

import (
	"encoding/json"
	"fmt"
	"log"

	"wowcollector.io/entities/documents"
	mountrepository "wowcollector.io/repository/repositories/mount-repository"
	seedrepository "wowcollector.io/repository/repositories/seed-repository"
)

func MountsSeeder() {
	var seedName = "mounts"
	fmt.Println("Mounts seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		fmt.Println("Mounts seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/mounts.json")

	var mounts []documents.MountDocument
	err = json.Unmarshal(byteValue, &mounts)
	if err != nil {
		log.Fatal("Error parsing mounts.json")
	}

	for _, element := range mounts {
		mountrepository.GetRepository().CreateMount(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	fmt.Println("Mounts seeder done")
}
