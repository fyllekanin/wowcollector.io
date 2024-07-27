package tasks

import (
	"fmt"
	"slices"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
	"wowcollector.io/internal/entities/documents"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	achievementcategoryrepository "wowcollector.io/internal/repository/repositories/achievement-category-repository"
	achievementrepository "wowcollector.io/internal/repository/repositories/achievement-repository"
	battlenethttp "wowcollector.io/internal/services/http/battle-net-http"
)

func ScanAchievements(region blizzarddata.BattleNetRegion) {
	zap.L().Info(fmt.Sprintf("Starting scan of achievements for region %s\n", region))
	index := battlenethttp.GetInstance().GetAchievementCategoryIndex(region)
	allCategories := slices.Concat(index.Categories, index.RootCategories)

	existingAchievementCategories, err := achievementcategoryrepository.GetRepository().GetAchievementCategories()
	if err != nil {
		zap.L().Error("Error fetching existing achievement categories")
		return
	}
	existingAchievements, err := achievementrepository.GetRepository().GetAchievements()
	if err != nil {
		zap.L().Error("Error fetching existing achievements")
		return
	}

	var wg sync.WaitGroup
	for _, category := range allCategories {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runCategory(region, category, existingAchievementCategories, existingAchievements)
		}()
	}
	wg.Wait()

	zap.L().Info(fmt.Sprintf("Finished scan of achievements for region %s\n", region))
}

func runCategory(region blizzarddata.BattleNetRegion, category *httpresponses.BattleNetAchievementCategory,
	existingCategories []*documents.AchievementCategoryDocument, existingAchievements []*documents.AchievementDocument) {

	zap.L().Info(fmt.Sprintf("Scanning achievement category %d", category.Id))
	achievementCategory := battlenethttp.GetInstance().GetAchievementCategory(region, category.Id)
	if achievementCategory == nil {
		zap.L().Info(fmt.Sprintf("Could not get achievement category %d", category.Id))
		return
	}

	existingCategory := getExistingCategory(existingCategories, category.Id)
	document := &documents.AchievementCategoryDocument{
		Id:             category.Id,
		Name:           category.Name,
		IsRootCategory: achievementCategory.ParentCategory == nil,
		RootCategoryId: getRootCategoryId(*achievementCategory),
		DisplayOrder:   achievementCategory.DisplayOrder,
	}

	if existingCategory == nil {
		document.ObjectID = primitive.NewObjectID()
		achievementcategoryrepository.GetRepository().CreateAchievementCategory(document)
		zap.L().Info(fmt.Sprintf("Added new achievement category %d", category.Id))
	} else {
		if !document.IsEqual(existingCategory) {
			document.ObjectID = existingCategory.ObjectID
			achievementcategoryrepository.GetRepository().UpdateAchievementCategory(document)
			zap.L().Info(fmt.Sprintf("Updated achievement category %d", category.Id))
		}
	}

	var wg sync.WaitGroup
	if achievementCategory.Achievements != nil {
		for _, element := range achievementCategory.Achievements {
			go runAchievement(region, category, element, existingAchievements)
		}
		for _, element := range achievementCategory.Achievements {
			wg.Add(1)
			go func() {
				defer wg.Done()
				runAchievement(region, category, element, existingAchievements)
			}()
		}
	}

	wg.Wait()
}

func runAchievement(region blizzarddata.BattleNetRegion, category *httpresponses.BattleNetAchievementCategory,
	achievementIndex *httpresponses.BattleNetAchievementIndexItem, existingAchievements []*documents.AchievementDocument) {

	zap.L().Info(fmt.Sprintf("Scanning achievement %d", achievementIndex.Id))
	achievement := battlenethttp.GetInstance().GetAchievement(region, achievementIndex.Id)
	if achievement == nil {
		zap.L().Info(fmt.Sprintf("Could not get achievement %d", achievementIndex.Id))
		return
	}
	achievementMedia := battlenethttp.GetInstance().GetAchievementMedia(region, achievement.Id)
	if achievementMedia == nil {
		zap.L().Info(fmt.Sprintf("Error fetching media for achievement %d", achievement.Id))
		return
	}
	existingAchievement := getExistingAchievement(existingAchievements, achievement.Id)
	document := &documents.AchievementDocument{
		Id:            achievement.Id,
		Name:          achievement.Name,
		Description:   achievement.Description,
		Points:        achievement.Points,
		IsAccountWide: achievement.IsAccountWide,
		Icon:          achievementMedia.GetIconAsset(),
		DisplayOrder:  achievement.DisplayOrder,
		CategoryId:    category.Id,
	}

	if existingAchievement == nil {
		document.ObjectID = primitive.NewObjectID()
		achievementrepository.GetRepository().CreateAchievement(document)
		zap.L().Info(fmt.Sprintf("Added new achievement %d", achievement.Id))
	} else {
		if !document.IsEqual(existingAchievement) {
			document.ObjectID = existingAchievement.ObjectID
			achievementrepository.GetRepository().UpdateAchievement(document)
			zap.L().Info(fmt.Sprintf("Updated achievement %d", achievement.Id))
		}
	}
}

func getExistingCategory(categories []*documents.AchievementCategoryDocument, id int) *documents.AchievementCategoryDocument {
	for _, element := range categories {
		if element.Id == id {
			return element
		}
	}
	return nil
}

func getExistingAchievement(achievements []*documents.AchievementDocument, id int) *documents.AchievementDocument {
	for _, element := range achievements {
		if element.Id == id {
			return element
		}
	}
	return nil
}

func getRootCategoryId(item httpresponses.BattleNetAchievementIndex) int {
	if item.ParentCategory == nil {
		return -1
	} else {
		return item.ParentCategory.Id
	}
}
