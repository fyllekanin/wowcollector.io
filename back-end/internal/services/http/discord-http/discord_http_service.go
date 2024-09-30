package discordhttp

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	httpresponses "wowcollector.io/internal/entities/http-responses"
)

type DiscordHttpService struct {
	rateLimiter ratelimit.Limiter
}

var instance *DiscordHttpService

func GetInstance() *DiscordHttpService {
	if instance == nil {
		instance = &DiscordHttpService{
			rateLimiter: ratelimit.New(100),
		}
	}
	return instance
}

func (s *DiscordHttpService) GetUserInfo(accessToken string) *httpresponses.DiscordUserInfo {
	response, err := s.doRequest("https://discord.com/api/v10/users/@me", true, accessToken)
	if err != nil {
		zap.L().Info("Error getting user @me:" + err.Error())
		return nil
	}

	var result httpresponses.DiscordUserInfo
	err = json.Unmarshal(response, &result)
	if err != nil {
		zap.L().Info("Error decoding user @me:" + err.Error())
		return nil
	}
	return &result
}

func (s *DiscordHttpService) GetAuth(redirectUri string, scope string, code string) *httpresponses.DiscordAuth {
	body := fmt.Sprintf("redirect_uri=%s&scope=%s&code=%s&grant_type=authorization_code", redirectUri, scope, code)

	req, err := http.NewRequest("POST", "https://discord.com/api/v10/oauth2/token", bytes.NewBufferString(body))
	if err != nil {
		zap.L().Info("Error creating auth request:" + err.Error())
		return nil
	}

	req.Header.Set("Authorization", getBasicAuthenticationHeader(os.Getenv("DISCORD_CLIENT_ID"), os.Getenv("DISCORD_SECRET")))
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

	var tokenResponse httpresponses.DiscordAuth
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&tokenResponse)
	if err != nil {
		zap.L().Info("Error decoding auth token:" + err.Error())
		return nil
	}

	return &tokenResponse
}

func (s *DiscordHttpService) doRequest(url string, retry bool, accessToken string) ([]byte, error) {
	s.rateLimiter.Take()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		zap.L().Info("Error creating request:" + err.Error())
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		zap.L().Info("Error sending request:" + err.Error())
		return nil, err
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

func getBasicAuthenticationHeader(clientID string, clientSecret string) string {
	auth := clientID + ":" + clientSecret
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + encoded
}
