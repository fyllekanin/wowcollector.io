package tasks

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
	"wowcollector.io/internal/entities/documents"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	realmrepository "wowcollector.io/internal/repository/repositories/realm-repository"
	battlenethttp "wowcollector.io/internal/services/http/battle-net-http"
)

func ScanRealms(region blizzarddata.BattleNetRegion) {
	zap.L().Info(fmt.Sprintf("Starting scan of realms for region %s", region))
	repository := realmrepository.GetRepository()

	existingRealms, err := repository.GetRealms()
	if err != nil {
		zap.L().Error("Error fetching existing realms:" + err.Error())
		return
	}

	battleNetRealms := battlenethttp.GetInstance().GetRealms(region)
	if battleNetRealms == nil {
		zap.L().Error("Error fetching realms from battle.net")
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
			zap.L().Info(fmt.Sprintf("Added new realm with id %d, slug %s for region %s", realm.Id, realm.Slug, region))
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
				zap.L().Info(fmt.Sprintf("Updated realm with id %d, slug %s for region %s", realm.Id, realm.Slug, region))
			}
		}
	}

	zap.L().Info(fmt.Sprintf("Finished scan of realms for region %s", region))
}

func getExistingRealm(realms []*documents.RealmDocument, realm httpresponses.BattleNetRealm, region blizzarddata.BattleNetRegion) *documents.RealmDocument {
	for _, element := range realms {
		if element.Slug == realm.Slug && element.Region == region {
			return element
		}
	}
	return nil
}
