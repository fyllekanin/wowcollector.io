package tasks

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	blizzarddata "wowcollector.io/common/data/blizzard-data"
	battlenetEntities "wowcollector.io/entities"
	"wowcollector.io/entities/documents"
	mountrepository "wowcollector.io/repository/repositories/mount-repository"
	battlenethttp "wowcollector.io/services/http/battle-net-http"
	wowheadhttp "wowcollector.io/services/http/wow-head-http"
)

func ScanMounts(region blizzarddata.BattleNetRegion) {
	fmt.Printf("Starting scan of mounts for region %s\n", region)
	repository := mountrepository.GetRepository()

	existingMounts, err := repository.GetMounts()
	if err != nil {
		fmt.Println("Error fetching existing mounts:", err)
		return
	}

	battleNetMountsIndex := battlenethttp.GetInstance().GetMountsIndex(region)
	if battleNetMountsIndex == nil {
		fmt.Println("Error fetching mounts index from battle.net")
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
	fmt.Printf("Finished scan of mounts for region %s\n", region)
}

func runMount(region blizzarddata.BattleNetRegion, mountId int, existingMounts []*documents.MountDocument, repository mountrepository.MountRepository) {
	battleNetMount := battlenethttp.GetInstance().GetMount(region, mountId)
	if battleNetMount == nil {
		fmt.Printf("Error fetching mount with id %d\n", mountId)
		return
	}
	existingMount := getExistingMount(existingMounts, *battleNetMount)
	asset := getCreatureDisplay(battleNetMount)
	tooltip := wowheadhttp.GetInstance().GetMountIcon(mountId)

	if existingMount == nil {
		repository.CreateMount(&documents.MountDocument{
			ObjectID:        primitive.NewObjectID(),
			Id:              battleNetMount.Id,
			Name:            battleNetMount.Name,
			Description:     battleNetMount.Description,
			Source:          getSourceType(battleNetMount.Source),
			Faction:         getFactionType(battleNetMount.Faction),
			CreatureDisplay: asset,
			IsUnobtainable:  battleNetMount.ShouldExcludeIfUncollected,
			Icon:            tooltip.Icon,
		})
		fmt.Printf("Added new mount with id %d\n", mountId)
	} else {
		newMount := &documents.MountDocument{
			ObjectID:        existingMount.ObjectID,
			Id:              battleNetMount.Id,
			Name:            battleNetMount.Name,
			Description:     battleNetMount.Description,
			Source:          getSourceType(battleNetMount.Source),
			Faction:         getFactionType(battleNetMount.Faction),
			CreatureDisplay: asset,
			IsUnobtainable:  battleNetMount.ShouldExcludeIfUncollected,
			Icon:            tooltip.Icon,
		}

		if !newMount.IsEqual(existingMount) {
			repository.UpdateMount(newMount)
			fmt.Printf("Updated mount with id %d\n", newMount.Id)
		}
	}
}

func getSourceType(item *battlenetEntities.BattleNetMountSource) string {
	if item != nil {
		return item.Type
	}
	return ""
}

func getFactionType(item *battlenetEntities.BattleNetFaction) string {
	if item != nil {
		return item.Type
	}
	return ""
}

func getCreatureDisplay(mount *battlenetEntities.BattleNetMount) string {
	if len(mount.CreatureDisplays) > 0 {
		return strings.Replace("https://render.worldofwarcraft.com/us/npcs/zoom/creature-display-{id}.jpg", "{id}", strconv.Itoa(mount.CreatureDisplays[0].Id), 1)
	}
	return ""
}

func getExistingMount(mounts []*documents.MountDocument, mount battlenetEntities.BattleNetMount) *documents.MountDocument {
	for _, element := range mounts {
		if element.Id == mount.Id {
			return element
		}
	}
	return nil
}
