package battleNetHttp

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"wowcollector.io/common/data"
	battlenetEntities "wowcollector.io/entities/battlenet"
)

type BattleNetToken struct {
	accessToken string `json:"access_token"`
	expiresIn   int64  `json:"expires_in"`
	issuedAt    int64  `json:"issued_at"`
}

func (s *BattleNetToken) IsExpired() bool {
	if s.expiresIn == 0 {
		return true
	}
	expirationTime := s.issuedAt + (s.expiresIn * 1000)
	return expirationTime <= time.Now().UnixMilli()
}

type BattleNetHttpService struct {
	token *BattleNetToken
}

var instance *BattleNetHttpService

func GetInstance() *BattleNetHttpService {
	if instance == nil {
		instance = &BattleNetHttpService{}
	}
	return instance
}

func (s *BattleNetHttpService) GetMountsIndex(region data.BattleNetRegion) *battlenetEntities.BatleNetMountsIndex {
	response, err := s.doRequest("https://"+string(region)+".api.blizzard.com/data/wow/mount/index?namespace=static-"+string(region)+"&locale=en_US", true)
	if err != nil {
		fmt.Println("Error getting mounts index", err)
		return nil
	}

	var result battlenetEntities.BatleNetMountsIndex
	err = response.Decode(&result)
	if err != nil {
		fmt.Println("Error decoding mounts index:", err)
		return nil
	}
	return &result
}

func (s *BattleNetHttpService) doRequest(url string, retry bool) (*json.Decoder, error) {
	response, err := http.Get(url + "&access_token=" + s.getAccessToken())
	if err != nil {
		fmt.Println("Error get request:", err)
		return nil, errors.New("Failed creating request")
	}
	defer response.Body.Close()

	if response.StatusCode == 401 && retry {
		return s.doRequest(url, false)
	}
	if response.StatusCode == 429 && retry {
		fmt.Printf("Error rate limit hit, sleep for 1 second and try again")
		time.Sleep(1 * time.Second)
		return s.doRequest(url, true)
	}
	return json.NewDecoder(response.Body), nil
}

func (s *BattleNetHttpService) getAccessToken() string {
	if s.token == nil || s.token.IsExpired() {
		s.token = resolveToken()
	}
	return s.token.accessToken
}

func resolveToken() *BattleNetToken {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://oauth.battle.net/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating access token request", err)
		return nil
	}

	req.Header.Set("Authorization", getBasicAuthenticationHeader(os.Getenv("BATTLE_NET_CLIENT_ID"), os.Getenv("BATTLE_NET_CLIENT_SECRET")))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request for access token", err)
		return nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d\n", response.StatusCode)
		return nil
	}

	var tokenResponse BattleNetToken
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&tokenResponse)
	if err != nil {
		fmt.Println("Error decoding access token:", err)
		return nil
	}

	return &tokenResponse
}

func getBasicAuthenticationHeader(clientID string, clientSecret string) string {
	auth := clientID + ":" + clientSecret
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + encoded
}
