package githubhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	httprequests "wowcollector.io/internal/entities/http-requests"
)

type GithubHttpService struct {
	rateLimiter ratelimit.Limiter
}

var instance *GithubHttpService

func GetInstance() *GithubHttpService {
	if instance == nil {
		instance = &GithubHttpService{
			rateLimiter: ratelimit.New(100),
		}
	}
	return instance
}

func (s *GithubHttpService) CreateIssue(body *httprequests.GithubIssueBody) error {
	s.rateLimiter.Take()

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", body.Owner, body.Repo)

	issueData := map[string]interface{}{
		"title":     body.GetTitle(),
		"body":      body.GetBody(),
		"assigness": []string{},
		"labels":    body.GetLabels(),
	}

	jsonData, err := json.Marshal(issueData)
	if err != nil {
		zap.L().Error("Error creating JSON body for github issue: " + err.Error())
		return errors.New("error creating JSON Body")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		zap.L().Error("Error creating github issue request: " + err.Error())
		return errors.New("error creating request")
	}

	req.Header.Set("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("Error sending github request: " + err.Error())
		return errors.New("error sending request")
	}
	defer resp.Body.Close()

	return nil
}
