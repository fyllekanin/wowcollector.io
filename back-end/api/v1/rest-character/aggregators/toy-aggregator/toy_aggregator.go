package toyaggregator

import (
	"slices"

	"wowcollector.io/internal/entities/documents"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	"wowcollector.io/internal/entities/response"
	toyrepository "wowcollector.io/internal/repository/repositories/toy-repository"
)

func GetToysAggregation(collection httpresponses.BattleNetCharacterToyCollection, view documents.ToyViewDocument) []response.ToyCollectionCategory {
	collectedIds := getCollectedToyIds(collection)
	toys, _ := toyrepository.GetRepository().GetToys()
	filteredToys := getFilteredToys(toys, collectedIds)

	return getToyCategories(view, filteredToys, collectedIds)
}

func getToyCategories(view documents.ToyViewDocument, toys map[int]*documents.ToyDocument, collectedIds []int) []response.ToyCollectionCategory {
	var result []response.ToyCollectionCategory
	for _, element := range view.Categories {
		result = append(result, getToyCategory(element, toys, collectedIds))
	}
	if view.IsUnknownIncluded {
		result = append(result, getUnknownCategory(toys, collectedIds))
	}
	return result
}

func getUnknownCategory(toys map[int]*documents.ToyDocument, collectedIds []int) response.ToyCollectionCategory {
	var toysResult []response.ToyCollectionToy
	for _, toy := range toys {
		toysResult = append(toysResult, response.ToyCollectionToy{
			Name:        toy.Name,
			Id:          toy.Id,
			ItemId:      toy.ItemId,
			IsCollected: slices.Contains(collectedIds, toy.Id),
			Assets: &response.ToyCollectionToyAssets{
				LargeIcon: toy.Icon,
			},
		})
	}

	return response.ToyCollectionCategory{
		Name:  "Unknown",
		Order: 99999,
		Toys:  toysResult,
	}
}

func getToyCategory(category documents.ToyViewCategory, toys map[int]*documents.ToyDocument, collectedIds []int) response.ToyCollectionCategory {
	var subCategories []response.ToyCollectionCategory
	for _, category := range category.Categories {
		subCategories = append(subCategories, getToyCategory(category, toys, collectedIds))
	}

	var toysResult []response.ToyCollectionToy
	for _, categoryToy := range category.Toys {
		toy := toys[categoryToy.Id]
		if toy == nil {
			continue
		}
		delete(toys, categoryToy.Id)
		toysResult = append(toysResult, response.ToyCollectionToy{
			Name:        toy.Name,
			Id:          toy.Id,
			ItemId:      toy.ItemId,
			IsCollected: slices.Contains(collectedIds, categoryToy.Id),
			Assets: &response.ToyCollectionToyAssets{
				LargeIcon: toy.Icon,
			},
		})
	}

	return response.ToyCollectionCategory{
		Name:       category.Name,
		Order:      category.Order,
		Categories: subCategories,
		Toys:       toysResult,
	}
}

func getCollectedToyIds(collection httpresponses.BattleNetCharacterToyCollection) []int {
	var result []int
	for _, element := range collection.Toys {
		result = append(result, element.Toys.Id)
	}
	return result
}

func getFilteredToys(toys []*documents.ToyDocument, collectedIds []int) map[int]*documents.ToyDocument {
	var result = make(map[int]*documents.ToyDocument)

	for _, element := range toys {
		if element.IsUnobtainable && !slices.Contains(collectedIds, element.Id) {
			continue
		}
		result[element.Id] = element
	}
	return result
}
