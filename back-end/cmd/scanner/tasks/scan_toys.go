package tasks

import (
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
	"wowcollector.io/internal/entities/documents"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	toyrepository "wowcollector.io/internal/repository/repositories/toy-repository"
	battlenethttp "wowcollector.io/internal/services/http/battle-net-http"
)

func ScanToys(region blizzarddata.BattleNetRegion) {
	zap.L().Info(fmt.Sprintf("Starting scan of toys for region %s", region))
	repository := toyrepository.GetRepository()

	existingToys, err := repository.GetToys()
	if err != nil {
		zap.L().Error("Error fetching existing toys:" + err.Error())
		return
	}

	battleNetToysIndex := battlenethttp.GetInstance().GetToysIndex(region)
	if battleNetToysIndex == nil {
		zap.L().Error("Error fetching toys index from battle.net")
		return
	}

	var wg sync.WaitGroup
	for _, toy := range battleNetToysIndex.Toys {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runToy(region, toy.Id, existingToys, *repository)
		}()
	}

	wg.Wait()
	zap.L().Info(fmt.Sprintf("Finished scan of toys for region %s", region))
}

func runToy(region blizzarddata.BattleNetRegion, toyId int, existingToys []*documents.ToyDocument, repository toyrepository.ToyRepository) {
	zap.L().Info(fmt.Sprintf("Scanning toy id %d", toyId))
	battleNetToy := battlenethttp.GetInstance().GetToy(region, toyId)
	if battleNetToy == nil {
		zap.L().Error(fmt.Sprintf("Error fetching toy with id %d", toyId))
		return
	}

	existingToy := getExistingToy(existingToys, *battleNetToy)
	media := battlenethttp.GetInstance().GetMedia(region, "item", battleNetToy.Media.Id)

	document := &documents.ToyDocument{
		Id:             battleNetToy.Id,
		ItemId:         battleNetToy.Item.Id,
		Name:           battleNetToy.Item.Name,
		Source:         battleNetToy.Source.Name,
		IsUnobtainable: battleNetToy.ShouldExcludeIfUncollected,
		Icon:           media.GetIconAsset(),
	}

	if existingToy == nil {
		document.ObjectID = primitive.NewObjectID()
		repository.CreateToy(document)
		zap.L().Info(fmt.Sprintf("Added new toy with id %d", toyId))
	} else {
		document.ObjectID = existingToy.ObjectID
		if !document.IsEqual(existingToy) {
			repository.UpdateToy(document)
			zap.L().Info(fmt.Sprintf("Updated toy with id %d", toyId))
		}
	}
}

func getExistingToy(toys []*documents.ToyDocument, toy httpresponses.BattleNetToy) *documents.ToyDocument {
	for _, element := range toys {
		if element.Id == toy.Id {
			return element
		}
	}
	return nil
}
