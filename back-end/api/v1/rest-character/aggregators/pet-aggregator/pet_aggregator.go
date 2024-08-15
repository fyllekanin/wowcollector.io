package petaggregator

import (
	"slices"

	"wowcollector.io/internal/entities/documents"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	"wowcollector.io/internal/entities/response"
	petrepository "wowcollector.io/internal/repository/repositories/pet-repository"
)

func GetPetsAggregation(collection httpresponses.BattleNetCharacterPetCollection, view documents.PetViewDocument) []response.PetCollectionCategory {
	collectedIds := getCollectedPetIds(collection)
	pets, _ := petrepository.GetRepository().GetPets()
	filteredPets := getFilteredPets(pets, collectedIds)

	return getPetCategories(view, filteredPets, collectedIds)
}

func getPetCategories(view documents.PetViewDocument, pets map[int]*documents.PetDocument, collectedIds []int) []response.PetCollectionCategory {
	var result []response.PetCollectionCategory
	for _, element := range view.Categories {
		result = append(result, getPetCategory(element, pets, collectedIds))
	}
	if view.IsUnknownIncluded {
		result = append(result, getUnknownCategory(pets, collectedIds))
	}
	return result
}

func getUnknownCategory(pets map[int]*documents.PetDocument, collectedIds []int) response.PetCollectionCategory {
	var petsResult []response.PetCollectionPet
	for _, pet := range pets {
		petsResult = append(petsResult, response.PetCollectionPet{
			Name:        pet.Name,
			Id:          pet.Id,
			IsCollected: slices.Contains(collectedIds, pet.Id),
			Assets: &response.PetAssets{
				LargeIcon: pet.Icon,
			},
		})
	}

	return response.PetCollectionCategory{
		Name:  "Unknown",
		Order: 99999,
		Pets:  petsResult,
	}
}

func getPetCategory(category documents.PetViewCategory, pets map[int]*documents.PetDocument, collectedIds []int) response.PetCollectionCategory {
	var subCategories []response.PetCollectionCategory
	for _, category := range category.Categories {
		subCategories = append(subCategories, getPetCategory(category, pets, collectedIds))
	}

	var petsResult []response.PetCollectionPet
	for _, categoryPet := range category.Pets {
		pet := pets[categoryPet.Id]
		if pet == nil {
			continue
		}
		delete(pets, categoryPet.Id)
		petsResult = append(petsResult, response.PetCollectionPet{
			Name:        pet.Name,
			Id:          pet.Id,
			IsCollected: slices.Contains(collectedIds, categoryPet.Id),
			Assets: &response.PetAssets{
				LargeIcon: pet.Icon,
			},
		})
	}

	return response.PetCollectionCategory{
		Name:       category.Name,
		Order:      category.Order,
		Categories: subCategories,
		Pets:       petsResult,
	}
}

func getCollectedPetIds(collection httpresponses.BattleNetCharacterPetCollection) []int {
	var result []int
	for _, element := range collection.Pets {
		result = append(result, element.Id)
	}
	return result
}

func getFilteredPets(pets []*documents.PetDocument, collectedIds []int) map[int]*documents.PetDocument {
	var result = make(map[int]*documents.PetDocument)

	for _, element := range pets {
		if element.IsUnobtainable && !slices.Contains(collectedIds, element.Id) {
			continue
		}
		result[element.Id] = element
	}
	return result
}
