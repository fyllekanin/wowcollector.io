package seeds

import (
	"encoding/json"
	"fmt"
	"log"

	"wowcollector.io/entities/documents"
	realmrepository "wowcollector.io/repository/repositories/realm-repository"
	seedrepository "wowcollector.io/repository/repositories/seed-repository"
)

var seedName string = "realms"

func RealmsSeeder() {
	fmt.Println("Realms seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		fmt.Println("Realms seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/realms.json")

	var realms []documents.RealmDocument
	err = json.Unmarshal(byteValue, &realms)
	if err != nil {
		log.Fatal("Error parsing realms.json")
	}

	for _, element := range realms {
		realmrepository.GetRepository().CreateRealm(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	fmt.Println("Realms seeder done")
}
