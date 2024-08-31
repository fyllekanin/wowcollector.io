package tasks

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	mountrepository "wowcollector.io/internal/repository/repositories/mount-repository"
	battlenethttp "wowcollector.io/internal/services/http/battle-net-http"
	wowheadhttp "wowcollector.io/internal/services/http/wow-head-http"
)

var largeIconUrl = "https://wow.zamimg.com/images/wow/icons/large/{name}.jpg"
var smallIconUrl = "https://wow.zamimg.com/images/wow/icons/small/{name}.jpg"

func ScanMounts(region string) {
	zap.L().Info(fmt.Sprintf("Starting scan of mounts for region %s", region))
	repository := mountrepository.GetRepository()

	existingMounts, err := repository.GetMounts()
	if err != nil {
		zap.L().Error("Error fetching existing mounts:" + err.Error())
		return
	}

	battleNetMountsIndex := battlenethttp.GetInstance().GetMountsIndex(region)
	if battleNetMountsIndex == nil {
		zap.L().Error("Error fetching mounts index from battle.net")
		return
	}

	var wg sync.WaitGroup
	for _, mount := range battleNetMountsIndex.Mounts {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runMount(region, mount.Id, existingMounts, *repository)
		}()
	}

	wg.Wait()
	zap.L().Info(fmt.Sprintf("Finished scan of mounts for region %s", region))
}

func runMount(region string, mountId int, existingMounts []*documents.MountDocument, repository mountrepository.MountRepository) {
	zap.L().Info(fmt.Sprintf("Scanning mount id %d", mountId))
	battleNetMount := battlenethttp.GetInstance().GetMount(region, mountId)
	if battleNetMount == nil {
		zap.L().Error(fmt.Sprintf("Error fetching mount with id %d", mountId))
		return
	}

	existingMount := getExistingMount(existingMounts, *battleNetMount)
	asset := getCreatureDisplay(battleNetMount)
	tooltip := wowheadhttp.GetInstance().GetMountIcon(mountId)
	document := &documents.MountDocument{
		Id:             battleNetMount.Id,
		Name:           battleNetMount.Name,
		Description:    battleNetMount.Description,
		Source:         getSourceType(battleNetMount.Source),
		Faction:        getFactionType(battleNetMount.Faction),
		IsUnobtainable: battleNetMount.ShouldExcludeIfUncollected,
		Assets: &documents.MountDocumentAssets{
			Display:   asset,
			SmallIcon: getSmallIcon(tooltip.Icon),
			LargeIcon: getLargeIcon(tooltip.Icon),
		},
	}

	if existingMount == nil {
		document.ObjectID = primitive.NewObjectID()
		repository.CreateMount(document)
		zap.L().Info(fmt.Sprintf("Added new mount with id %d", mountId))
	} else {
		document.ObjectID = existingMount.ObjectID
		if !document.IsEqual(existingMount) {
			repository.UpdateMount(document)
			zap.L().Info(fmt.Sprintf("Updated mount with id %d", document.Id))
		}
	}
}

func getSourceType(item *httpresponses.BattleNetSource) string {
	if item != nil {
		return item.Type
	}
	return ""
}

func getFactionType(item *httpresponses.BattleNetFaction) string {
	if item != nil {
		return item.Type
	}
	return ""
}

func getCreatureDisplay(mount *httpresponses.BattleNetMount) string {
	if len(mount.CreatureDisplays) > 0 {
		return strings.Replace("https://render.worldofwarcraft.com/us/npcs/zoom/creature-display-{id}.jpg", "{id}", strconv.Itoa(mount.CreatureDisplays[0].Id), 1)
	}
	return ""
}

func getExistingMount(mounts []*documents.MountDocument, mount httpresponses.BattleNetMount) *documents.MountDocument {
	for _, element := range mounts {
		if element.Id == mount.Id {
			return element
		}
	}
	return nil
}

func getSmallIcon(name string) string {
	return strings.Replace(smallIconUrl, "{name}", name, 1)
}

func getLargeIcon(name string) string {
	return strings.Replace(largeIconUrl, "{name}", name, 1)
}
