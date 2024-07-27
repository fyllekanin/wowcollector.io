package aggregator

import (
	"fmt"
	"slices"
	"sync"

	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	"wowcollector.io/internal/entities/response"
	achievementcategoryrepository "wowcollector.io/internal/repository/repositories/achievement-category-repository"
	achievementrepository "wowcollector.io/internal/repository/repositories/achievement-repository"
)

func GetAchievementAggregation(character httpresponses.BattleNetCharacter, collection httpresponses.BattleNetCharacterAchievements) []response.AchievementCollectionCategory {
	collectedIds := getCollectedAchievementIds(collection)
	achievementCategories, _ := achievementcategoryrepository.GetRepository().GetAchievementCategories()
	var items []response.AchievementCollectionCategory

	var wg sync.WaitGroup
	ch := make(chan response.AchievementCollectionCategory, len(achievementCategories))

	for _, element := range achievementCategories {
		if !element.IsRootCategory {
			continue
		}
		wg.Add(1)
		go func(element *documents.AchievementCategoryDocument) {
			defer wg.Done()
			ch <- getCategoryResponse(element, achievementCategories, collectedIds)
		}(element)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		items = append(items, result)
	}

	return items
}

func getCategoryResponse(category *documents.AchievementCategoryDocument, categories []*documents.AchievementCategoryDocument, collectedIds []int) response.AchievementCollectionCategory {
	var subCategories []response.AchievementCollectionCategory
	var wg sync.WaitGroup
	ch := make(chan response.AchievementCollectionCategory, len(categories))

	for _, element := range categories {
		if element.RootCategoryId == category.Id {
			wg.Add(1)
			go func(element *documents.AchievementCategoryDocument) {
				defer wg.Done()
				ch <- getCategoryResponse(element, categories, collectedIds)
			}(element)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		subCategories = append(subCategories, result)
	}

	return response.AchievementCollectionCategory{
		Name:         category.Name,
		Achievements: getAchievementsForCategory(category, collectedIds),
		Categories:   subCategories,
		DisplayOrder: category.DisplayOrder,
	}
}

func getAchievementsForCategory(category *documents.AchievementCategoryDocument, collectedIds []int) []response.AchievementCollectionAchievement {
	achievements, err := achievementrepository.GetRepository().GetAchievementsForCategoryId(category.Id)
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error getting achievements for category %d", category.Id))
		return nil
	}

	var items []response.AchievementCollectionAchievement
	var wg sync.WaitGroup
	ch := make(chan response.AchievementCollectionAchievement, len(achievements))

	for _, element := range achievements {
		wg.Add(1)
		go func(element *documents.AchievementDocument) {
			defer wg.Done()
			ch <- response.AchievementCollectionAchievement{
				Id:           element.Id,
				Name:         element.Name,
				Description:  element.Description,
				Points:       element.Points,
				DisplayOrder: element.DisplayOrder,
				Icon:         element.Icon,
				IsCompleted:  slices.Contains(collectedIds, element.Id),
			}
		}(element)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		items = append(items, result)
	}

	return items
}

func getCollectedAchievementIds(collection httpresponses.BattleNetCharacterAchievements) []int {
	var result []int
	for _, element := range collection.Achievements {
		result = append(result, element.Id)
	}
	return result
}
