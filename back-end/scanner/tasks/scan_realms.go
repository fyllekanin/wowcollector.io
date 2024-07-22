package tasks

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"wowcollector.io/common/data"
	battlenetEntities "wowcollector.io/entities"
	"wowcollector.io/entities/documents"
	realmrepository "wowcollector.io/repository/repositories"
	battleNetHttp "wowcollector.io/services/http"
)

func ScanRealms(region data.BattleNetRegion) {
	fmt.Printf("Starting scan of realms for region %s\n", region)
	repository := realmrepository.GetRepository()

	existingRealms, err := repository.GetRealms()
	if err != nil {
		fmt.Println("Error fetching existing realms", err)
		return
	}

	battleNetRealms := battleNetHttp.GetInstance().GetRealms(region)
	if battleNetRealms == nil {
		fmt.Println("Error fetching realms from battle.net")
		return
	}

	for _, realm := range battleNetRealms.Realms {
		existingRealm := getExistingRealm(existingRealms, realm, region)
		if existingRealm == nil {
			repository.CreateRealm(&documents.RealmDocument{
				ObjectID: primitive.NewObjectID(),
				Id:       realm.Id,
				Name:     realm.Name,
				Slug:     realm.Slug,
				Region:   region,
			})
			fmt.Printf("Added new realm with id %d, slug %s for region %s\n", realm.Id, realm.Slug, region)
		} else {
			newRealm := &documents.RealmDocument{
				ObjectID: existingRealm.ObjectID,
				Id:       realm.Id,
				Name:     realm.Name,
				Slug:     realm.Slug,
				Region:   region,
			}

			if !newRealm.IsEqual(existingRealm) {
				repository.UpdateRealm(newRealm)
				fmt.Println(newRealm != existingRealm)
				log.Fatal()
				fmt.Printf("Updated realm with id %d, slug %s for region %s\n", realm.Id, realm.Slug, region)
			}
		}
	}

	fmt.Printf("Finished scan of realms for region %s\n", region)
}

func getExistingRealm(realms []*documents.RealmDocument, realm battlenetEntities.BattleNetRealm, region data.BattleNetRegion) *documents.RealmDocument {
	for _, element := range realms {
		if element.Slug == realm.Slug && element.Region == region {
			return element
		}
	}
	return nil
}
