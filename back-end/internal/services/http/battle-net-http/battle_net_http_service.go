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
	httpresponses "wowcollector.io/internal/entities/http-responses"
)

type BattleNetToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Sub         string `json:"sub"`
	ExpiresIn   int64  `json:"expires_in"`
	issuedAt    int64
}

func (s *BattleNetToken) isExpired() bool {
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

func (s *BattleNetHttpService) Ping() {
	_, err := s.doRequest("https://eu.api.blizzard.com/data/wow/region/index?namespace=dynamic-eu&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Fatal("Failed ping (region index):" + err.Error())
	}
}

func (s *BattleNetHttpService) GetCharacter(region string, realm string, character string) *httpresponses.BattleNetCharacter {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/profile/wow/character/"+realm+"/"+character+"?namespace=profile-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting character:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetCharacter
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding character:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetCharacterMedia(region string, realm string, character string) *httpresponses.BattleNetMedia {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/profile/wow/character/"+realm+"/"+character+"/character-media?namespace=profile-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting character media:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetMedia
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding character media:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetAchievementCategoryIndex(region string) *httpresponses.BattleNetAchievementCategoryIndex {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/achievement-category/index?namespace=static-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting achievement category index:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetAchievementCategoryIndex
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding achievement category index:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetAchievementCategory(region string, categoryId int) *httpresponses.BattleNetAchievementIndex {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/achievement-category/"+strconv.Itoa(categoryId)+"?namespace=static-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting achievement category:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetAchievementIndex
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding achievement category:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetAchievement(region string, id int) *httpresponses.BattleNetAchievement {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/achievement/"+strconv.Itoa(id)+"?namespace=static-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting achievement:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetAchievement
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding achievement:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetMedia(region string, kind string, id int) *httpresponses.BattleNetMedia {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/media/"+kind+"/"+strconv.Itoa(id)+"?namespace=static-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting media:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetMedia
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding media:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetCharacterAchievementCollection(region string, realm string, character string) *httpresponses.BattleNetCharacterAchievements {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/profile/wow/character/"+realm+"/"+character+"/achievements?namespace=profile-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting character achievement collection:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetCharacterAchievements
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding character achievement collection:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetCharacterMountCollection(region string, realm string, character string) *httpresponses.BattleNetCharacterMountCollection {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/profile/wow/character/"+realm+"/"+character+"/collections/mounts?namespace=profile-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting character mount collection:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetCharacterMountCollection
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding character mount collection:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetCharacterToyCollection(region string, realm string, character string) *httpresponses.BattleNetCharacterToyCollection {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/profile/wow/character/"+realm+"/"+character+"/collections/toys?namespace=profile-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting character toy collection:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetCharacterToyCollection
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding character toy collection:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetCharacterPetCollection(region string, realm string, character string) *httpresponses.BattleNetCharacterPetCollection {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/profile/wow/character/"+realm+"/"+character+"/collections/pets?namespace=profile-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting character pet collection:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetCharacterPetCollection
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding character pet collection:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetMountsIndex(region string) *httpresponses.BattleNetMountsIndex {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/mount/index?namespace=static-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting mounts index;" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetMountsIndex
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding mounts index:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetPetsIndex(region string) *httpresponses.BattleNetPetsIndex {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/pet/index?namespace=static-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting pets index;" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetPetsIndex
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding pets index:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetPet(region string, id int) *httpresponses.BattleNetPet {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/pet/"+strconv.Itoa(id)+"?namespace=static-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting pet" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetPet
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding pet:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetMount(region string, id int) *httpresponses.BattleNetMount {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/mount/"+strconv.Itoa(id)+"?namespace=static-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting mount" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetMount
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding mount:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetToysIndex(region string) *httpresponses.BattleNetToysIndex {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/toy/index?namespace=static-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting toys index;" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetToysIndex
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding toys index:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetToy(region string, id int) *httpresponses.BattleNetToy {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/toy/"+strconv.Itoa(id)+"?namespace=static-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting toy;" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetToy
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding toy:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetRealms(region string) *httpresponses.BattleNetRealms {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/realm/index?namespace=dynamic-"+string(region)+"&locale=en_US", true, s.getAccessToken())
	if err != nil {
		zap.L().Info("Error getting realms:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetRealms
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding realms:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetBattleNetUserInfo(accessToken string) *httpresponses.BattleNetUserInfo {
	response, err := s.doRequest("https://oauth.battle.net/oauth/userinfo", true, accessToken)
	if err != nil {
		zap.L().Info("Error getting battle net user info:" + err.Error())
		return nil
	}

	var result httpresponses.BattleNetUserInfo
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding battle net user info:" + err.Error())
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) GetAuth(redirectUri string, scope string, code string) *httpresponses.BattleNetAuth {
	body := fmt.Sprintf("redirectUri=%s&scope=%s&code=%s&grant_type=authorization_code", redirectUri, scope, code)

	req, err := http.NewRequest("POST", "https://oauth.battle.net/token", bytes.NewBufferString(body))
	if err != nil {
		zap.L().Info("Error creating auth request:" + err.Error())
		return nil
	}

	req.Header.Set("Authorization", getBasicAuthenticationHeader(os.Getenv("BATTLE_NET_CLIENT_ID"), os.Getenv("BATTLE_NET_CLIENT_SECRET")))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		zap.L().Info("Error sending request for auth token:" + err.Error())
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		zap.L().Info(fmt.Sprintf("Request failed with status code: %d", response.StatusCode))
		return nil
	}

	var tokenResponse httpresponses.BattleNetAuth
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&tokenResponse)
	if err != nil {
		zap.L().Info("Error decoding auth token:" + err.Error())
		return nil
	}

	return &tokenResponse
}

func (s *BattleNetHttpService) doRequest(url string, retry bool, accessToken string) ([]byte, error) {
	s.rateLimiter.Take()
	response, err := http.Get(url + "&access_token=" + accessToken)
	if err != nil {
		zap.L().Info("Error get request:" + err.Error())
		return nil, errors.New("failed creating request")
	}
	defer response.Body.Close()

	if response.StatusCode == 401 && retry {
		return s.doRequest(url, false, accessToken)
	}
	if response.StatusCode == 429 && retry {
		zap.L().Info("Error rate limit hit, sleep for 1 second and try again")
		time.Sleep(1 * time.Second)
		return s.doRequest(url, true, accessToken)
	}
	if response.StatusCode == 404 {
		zap.L().Info(fmt.Sprintf("Error not found result for url: %s", url))
		return nil, errors.New("not found error")
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		zap.L().Info("Error reading bytes:" + err.Error())
		return nil, nil
	}
	return bodyBytes, nil
}

func (s *BattleNetHttpService) getAccessToken() string {
	if s.token == nil || s.token.isExpired() {
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
