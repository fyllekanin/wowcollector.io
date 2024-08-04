package achievementaggregator

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

func GetAchievementAggregation(character httpresponses.BattleNetCharacter, collection httpresponses.BattleNetCharacterAchievements, rootCategoryId int) response.AchievementCollectionCategory {
	collectedIds := getCollectedAchievementIds(collection)
	category, _ := achievementcategoryrepository.GetRepository().GetAchievementCategoryWithId(rootCategoryId)

	return getCategoryResponse(category, collectedIds)
}

func getCategoryResponse(category *documents.AchievementCategoryDocument, collectedIds []int) response.AchievementCollectionCategory {
	var subCategories []response.AchievementCollectionCategory
	categories, err := achievementcategoryrepository.GetRepository().GetAchievementCategoriesForRootCategory(category.Id)
	if err == nil && len(categories) > 0 {
		var wg sync.WaitGroup
		ch := make(chan response.AchievementCollectionCategory, len(categories))

		for _, element := range categories {
			wg.Add(1)
			go func(element *documents.AchievementCategoryDocument) {
				defer wg.Done()
				ch <- getCategoryResponse(element, collectedIds)
			}(element)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		for result := range ch {
			subCategories = append(subCategories, result)
		}
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
	if len(achievements) == 0 {
		return make([]response.AchievementCollectionAchievement, 0)
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
