package tasks

import (
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
	"wowcollector.io/internal/entities/documents"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	petrepository "wowcollector.io/internal/repository/repositories/pet-repository"
	battlenethttp "wowcollector.io/internal/services/http/battle-net-http"
)

func ScanPets(region blizzarddata.BattleNetRegion) {
	zap.L().Info(fmt.Sprintf("Starting scan of pets for region %s", region))
	repository := petrepository.GetRepository()

	existingPets, err := repository.GetPets()
	if err != nil {
		zap.L().Error("Error fetching existing pets:" + err.Error())
		return
	}

	battleNetPetsIndex := battlenethttp.GetInstance().GetPetsIndex(region)
	if battleNetPetsIndex == nil {
		zap.L().Error("Error fetching pets index from battle.net")
		return
	}

	var wg sync.WaitGroup
	for _, pet := range battleNetPetsIndex.Pets {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runPet(region, pet.Id, existingPets, *repository)
		}()
	}

	wg.Wait()
	zap.L().Info(fmt.Sprintf("Finished scan of pets for region %s", region))
}

func runPet(region blizzarddata.BattleNetRegion, petId int, existingPets []*documents.PetDocument, repository petrepository.PetRepository) {
	zap.L().Info(fmt.Sprintf("Scanning pet id %d", petId))
	battleNetPet := battlenethttp.GetInstance().GetPet(region, petId)
	if battleNetPet == nil {
		zap.L().Error(fmt.Sprintf("Error fetching pet with id %d", petId))
		return
	}

	existingPet := getExistingPet(existingPets, *battleNetPet)

	document := &documents.PetDocument{
		Id:             battleNetPet.Id,
		Name:           battleNetPet.Name,
		IsUnobtainable: battleNetPet.ShouldExcludeIfUncollected,
		Icon:           battleNetPet.Icon,
		NpcId:          battleNetPet.Creature.Id,
	}

	if battleNetPet.Source != nil {
		document.Source = battleNetPet.Source.Name
	}

	if existingPet == nil {
		document.ObjectID = primitive.NewObjectID()
		repository.CreatePet(document)
		zap.L().Info(fmt.Sprintf("Added new pet with id %d", petId))
	} else {
		document.ObjectID = existingPet.ObjectID
		if !document.IsEqual(existingPet) {
			repository.UpdatePet(document)
			zap.L().Info(fmt.Sprintf("Updated pet with id %d", petId))
		}
	}
}

func getExistingPet(pets []*documents.PetDocument, pet httpresponses.BattleNetPet) *documents.PetDocument {
	for _, element := range pets {
		if element.Id == pet.Id {
			return element
		}
	}
	return nil
}
