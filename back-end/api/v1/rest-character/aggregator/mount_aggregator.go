package aggregator

import (
	"slices"

	"wowcollector.io/internal/entities/documents"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	"wowcollector.io/internal/entities/response"
	mountrepository "wowcollector.io/internal/repository/repositories/mount-repository"
)

func GetMountsAggregation(character httpresponses.BattleNetCharacter, collection httpresponses.BattleNetCharacterMountCollection, view documents.MountViewDocument) []response.MountCollectionCategory {
	collectedIds := getCollectedMountIds(collection)
	mounts, _ := mountrepository.GetRepository().GetMounts()
	filteredMounts := getFilteredMounts(character, mounts, collectedIds)

	return getMountCategories(view, filteredMounts, collectedIds)
}

func getMountCategories(view documents.MountViewDocument, mounts map[int]*documents.MountDocument, collectedIds []int) []response.MountCollectionCategory {
	var result []response.MountCollectionCategory
	for _, element := range view.Categories {
		result = append(result, getMountCategory(element, mounts, collectedIds))
	}
	if view.IsUnknownIncluded {
		result = append(result, getUnknownCategory(mounts, collectedIds))
	}
	return result
}

func getUnknownCategory(mounts map[int]*documents.MountDocument, collectedIds []int) response.MountCollectionCategory {
	var mountsResult []response.MountCollectionMount
	for _, mount := range mounts {
		mountsResult = append(mountsResult, response.MountCollectionMount{
			Name:        mount.Name,
			Description: mount.Description,
			Id:          mount.Id,
			IsCollected: slices.Contains(collectedIds, mount.Id),
			Assets: &response.MountCollectionMountAssets{
				Display:   mount.Assets.Display,
				SmallIcon: mount.Assets.SmallIcon,
				LargeIcon: mount.Assets.LargeIcon,
			},
		})
	}

	return response.MountCollectionCategory{
		Name:   "Unknown",
		Order:  99999,
		Mounts: mountsResult,
	}
}

func getMountCategory(category documents.MountViewCategory, mounts map[int]*documents.MountDocument, collectedIds []int) response.MountCollectionCategory {
	var subCategories []response.MountCollectionCategory
	for _, category := range category.Categories {
		subCategories = append(subCategories, getMountCategory(category, mounts, collectedIds))
	}

	var mountsResult []response.MountCollectionMount
	for _, categoryMount := range category.Mounts {
		mount := mounts[categoryMount.Id]
		if mount == nil {
			continue
		}
		delete(mounts, categoryMount.Id)
		mountsResult = append(mountsResult, response.MountCollectionMount{
			Name:        mount.Name,
			Description: mount.Description,
			Id:          mount.Id,
			IsCollected: slices.Contains(collectedIds, categoryMount.Id),
			Assets: &response.MountCollectionMountAssets{
				Display:   mount.Assets.Display,
				SmallIcon: mount.Assets.SmallIcon,
				LargeIcon: mount.Assets.LargeIcon,
			},
		})
	}

	return response.MountCollectionCategory{
		Name:       category.Name,
		Order:      category.Order,
		Categories: subCategories,
		Mounts:     mountsResult,
	}
}

func getCollectedMountIds(collection httpresponses.BattleNetCharacterMountCollection) []int {
	var result []int
	for _, element := range collection.Mounts {
		result = append(result, element.Mount.Id)
	}
	return result
}

func getFilteredMounts(character httpresponses.BattleNetCharacter, mounts []*documents.MountDocument, collectedIds []int) map[int]*documents.MountDocument {
	var result = make(map[int]*documents.MountDocument)

	for _, element := range mounts {
		if element.Faction != "" && element.Faction != character.GetFaction() {
			continue
		}

		if element.IsUnobtainable && !slices.Contains(collectedIds, element.Id) {
			continue
		}
		result[element.Id] = element
	}
	return result
}
