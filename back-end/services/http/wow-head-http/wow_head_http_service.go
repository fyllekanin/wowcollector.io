package wowheadhttp

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	entities "wowcollector.io/entities"
)

type WowHeadHttpService struct {
	itemPattern  *regexp.Regexp
	spellPattern *regexp.Regexp
}

var instance *WowHeadHttpService

func GetInstance() *WowHeadHttpService {
	if instance == nil {
		itemPattern, _ := regexp.Compile(`item=(\d+)`)
		spellPattern, _ := regexp.Compile(`spell=(\d+)`)
		instance = &WowHeadHttpService{
			itemPattern:  itemPattern,
			spellPattern: spellPattern,
		}
	}
	return instance
}

func (s *WowHeadHttpService) GetMountIcon(id int) *entities.WowHeadTooltip {
	client := &http.Client{
		Timeout: time.Second * 10,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	response, err := client.Get("https://www.wowhead.com/mount/" + strconv.Itoa(id))
	if err != nil {
		zap.L().Info("Error get request:" + err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 301 {
		locationHeader := response.Header.Get("location")
		tooltip, err := s.getTooltipFromLocation(locationHeader)
		if err == nil {
			return tooltip
		}
	}

	return nil
}

func (s *WowHeadHttpService) GetWowHeadTooltip(tooltipType string, id string) *entities.WowHeadTooltip {
	response, err := s.doRequest("https://nether.wowhead.com/tooltip/"+tooltipType+"/"+id+"?dataEnv=1&locale=0", true)
	if err != nil {
		zap.L().Info("Error getting wowhead tooltip:" + err.Error())
		return nil
	}

	var result entities.WowHeadTooltip
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding wowhead tooltip:" + err.Error())
		return nil
	}
	return &result
}

func (s *WowHeadHttpService) doRequest(url string, retry bool) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		zap.L().Info("Error get request:" + err.Error())
		return nil, errors.New("failed creating request")
	}
	defer response.Body.Close()

	if response.StatusCode == 429 && retry {
		zap.L().Info("Error rate limit hit, sleep for 1 second and try again")
		time.Sleep(1 * time.Second)
		return s.doRequest(url, true)
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		zap.L().Info("Error reading bytes:" + err.Error())
		return nil, nil
	}
	return bodyBytes, nil
}

func (s *WowHeadHttpService) getTooltipFromLocation(locationHeader string) (*entities.WowHeadTooltip, error) {
	if strings.Contains(locationHeader, "item=") {
		match := s.itemPattern.FindStringSubmatch(locationHeader)
		if match != nil {
			return s.GetWowHeadTooltip("item", match[1]), nil
		}
	}
	if strings.Contains(locationHeader, "spell=") {
		match := s.spellPattern.FindStringSubmatch(locationHeader)
		if match != nil {
			return s.GetWowHeadTooltip("spell", match[1]), nil
		}
	}
	return nil, errors.New("Location header not containing item or spell")
}
