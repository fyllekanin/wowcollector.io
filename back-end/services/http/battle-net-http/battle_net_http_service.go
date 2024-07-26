package battlenethttp

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	blizzarddata "wowcollector.io/common/data/blizzard-data"
	entities "wowcollector.io/entities"
)

type BattleNetToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Sub         string `json:"sub"`
	ExpiresIn   int64  `json:"expires_in"`
	issuedAt    int64
}

func (s *BattleNetToken) IsExpired() bool {
	if s.ExpiresIn == 0 {
		return true
	}
	expirationTime := s.issuedAt + (s.ExpiresIn * 1000)
	return expirationTime <= time.Now().UnixMilli()
}

type BattleNetHttpService struct {
	token       *BattleNetToken
	rateLimiter ratelimit.Limiter
}

var instance *BattleNetHttpService

func GetInstance() *BattleNetHttpService {
	if instance == nil {
		instance = &BattleNetHttpService{
			rateLimiter: ratelimit.New(100),
		}
	}
	return instance
}

func (s *BattleNetHttpService) GetCharacter(region blizzarddata.BattleNetRegion, realm string, character string) *entities.BattleNetCharacter {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/profile/wow/character/"+realm+"/"+character+"?namespace=profile-"+string(region)+"&locale=en_US", true)
	if err != nil {
		zap.L().Info("Error getting character:" + err.Error())
		return nil
	}

	var result entities.BattleNetCharacter
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding character:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetCharacterMountCollection(region blizzarddata.BattleNetRegion, realm string, character string) *entities.BattleNetCharacterMountCollection {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/profile/wow/character/"+realm+"/"+character+"/collections/mounts?namespace=profile-"+string(region)+"&locale=en_US", true)
	if err != nil {
		zap.L().Info("Error getting character mount collection:" + err.Error())
		return nil
	}

	var result entities.BattleNetCharacterMountCollection
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding character mount collection:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetMountsIndex(region blizzarddata.BattleNetRegion) *entities.BattleNetMountsIndex {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/mount/index?namespace=static-"+string(region)+"&locale=en_US", true)
	if err != nil {
		zap.L().Info("Error getting mounts index;" + err.Error())
		return nil
	}

	var result entities.BattleNetMountsIndex
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding mounts index:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetMount(region blizzarddata.BattleNetRegion, id int) *entities.BattleNetMount {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/mount/"+strconv.Itoa(id)+"?namespace=static-"+string(region)+"&locale=en_US", true)
	if err != nil {
		zap.L().Info("Error getting mount" + err.Error())
		return nil
	}

	var result entities.BattleNetMount
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding mount:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetRealms(region blizzarddata.BattleNetRegion) *entities.BattleNetRealms {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/realm/index?namespace=dynamic-"+string(region)+"&locale=en_US", true)
	if err != nil {
		zap.L().Info("Error getting realms:" + err.Error())
		return nil
	}

	var result entities.BattleNetRealms
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding realms:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) doRequest(url string, retry bool) ([]byte, error) {
	s.rateLimiter.Take()
	response, err := http.Get(url + "&access_token=" + s.getAccessToken())
	if err != nil {
		zap.L().Info("Error get request:" + err.Error())
		return nil, errors.New("failed creating request")
	}
	defer response.Body.Close()

	if response.StatusCode == 401 && retry {
		return s.doRequest(url, false)
	}
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

func (s *BattleNetHttpService) getAccessToken() string {
	if s.token == nil || s.token.IsExpired() {
		s.token = resolveToken()
		s.token.issuedAt = time.Now().UnixMilli()
	}
	return s.token.AccessToken
}

func resolveToken() *BattleNetToken {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://oauth.battle.net/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		zap.L().Info("Error creating access token request:" + err.Error())
		return nil
	}

	req.Header.Set("Authorization", getBasicAuthenticationHeader(os.Getenv("BATTLE_NET_CLIENT_ID"), os.Getenv("BATTLE_NET_CLIENT_SECRET")))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		zap.L().Info("Error sending request for access token:" + err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		zap.L().Info(fmt.Sprintf("Request failed with status code: %d", response.StatusCode))
		return nil
	}

	var tokenResponse BattleNetToken
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&tokenResponse)
	if err != nil {
		zap.L().Info("Error decoding access token:" + err.Error())
		return nil
	}

	return &tokenResponse
}

func getBasicAuthenticationHeader(clientID string, clientSecret string) string {
	auth := clientID + ":" + clientSecret
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + encoded
}
