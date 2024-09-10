package authorization

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"wowcollector.io/internal/entities/documents"
)

type AuthorizationTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Authorization struct {
	Id          string                     `json:"id"`
	DisplayName string                     `json:"displayName"`
	Tokens      *AuthorizationTokens       `json:"tokens"`
	Connections *documents.UserConnections `json:"connections"`
}

func GetJwt(userId string, validFor time.Duration) string {
	key := os.Getenv("JWT_SECRET_KEY")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     userId,
		"validUntil": time.Now().Add(validFor).Unix(),
	})
	value, _ := t.SignedString([]byte(key))
	return value
}
