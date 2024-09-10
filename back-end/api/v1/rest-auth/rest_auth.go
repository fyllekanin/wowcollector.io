package restauth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	errorcodes "wowcollector.io/internal/common/error-codes"
	"wowcollector.io/internal/entities/authorization"
	"wowcollector.io/internal/entities/documents"
	errorresponse "wowcollector.io/internal/entities/response/error"
	userrepository "wowcollector.io/internal/repository/repositories/user-repository"
	battlenethttp "wowcollector.io/internal/services/http/battle-net-http"
	discordhttp "wowcollector.io/internal/services/http/discord-http"
)

func GetRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/battle-net", getBattleNetAuth)
		r.Get("/discord", getDiscordAuth)
	})
}

// @summary Login with discord auth code
// @description Performs login for discord auth code
// @tags Auth
// @produce json
// @param code query string true "Auth code"
// @param redirectUri query string true "Redirect Uri"
// @param scope query string true "Scope"
// @success 200 {object} authorization.Authorization
// @failure 400 {object} errorresponse.ErrorResponse
// @router /api/v1/auth/discord [get]
func getDiscordAuth(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	redirectUri := r.URL.Query().Get("redirectUri")
	scope := r.URL.Query().Get("scope")
	zap.L().Info("Logging in for discord")

	if code == "" || redirectUri == "" || scope == "" {
		zap.L().Error("Error, missing query parameter(s)")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.REQUEST_DATA, "Query parameter(s) missing"))
		return
	}

	auth := discordhttp.GetInstance().GetAuth(
		redirectUri,
		scope,
		code,
	)

	if auth == nil {
		zap.L().Error("Error getting auth from discord")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Error getting auth"))
		return
	}

	userInfo := discordhttp.GetInstance().GetUserInfo(auth.AccessToken)
	if userInfo == nil {
		zap.L().Error("Error getting user info from battle net")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Error getting user info"))
		return
	}

	existing, err := userrepository.GetRepository().GetByDiscordId(userInfo.Id)
	if err != nil {
		existing = &documents.UserDocument{
			ObjectID:    primitive.NewObjectID(),
			DisplayName: userInfo.Username,
			Connections: &documents.UserConnections{
				DiscordId: userInfo.Id,
			},
		}
		userrepository.GetRepository().Create(existing)
	}

	body, err := json.Marshal(&authorization.Authorization{
		Id:          existing.ObjectID.Hex(),
		DisplayName: existing.DisplayName,
		Connections: existing.Connections,
		Tokens: &authorization.AuthorizationTokens{
			AccessToken:  authorization.GetJwt(existing.ObjectID.Hex(), time.Hour*8),
			RefreshToken: authorization.GetJwt(existing.ObjectID.Hex(), time.Hour*48),
		},
	})
	if err != nil {
		zap.L().Error("Failed to stringify response body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Invalid JSON body"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	zap.L().Info("Responded with login response")
}

// @summary Login with battle net auth code
// @description Performs login for battle net auth code
// @tags Auth
// @produce json
// @param code query string true "Auth code"
// @param redirectUri query string true "Redirect Uri"
// @param scope query string true "Scope"
// @success 200 {object} authorization.Authorization
// @failure 400 {object} errorresponse.ErrorResponse
// @router /api/v1/auth/battle-net [get]
func getBattleNetAuth(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	redirectUri := r.URL.Query().Get("redirectUri")
	scope := r.URL.Query().Get("scope")
	zap.L().Info("Logging in for battle net")

	if code == "" || redirectUri == "" || scope == "" {
		zap.L().Error("Error, missing query parameter(s)")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.REQUEST_DATA, "Query parameter(s) missing"))
		return
	}

	auth := battlenethttp.GetInstance().GetAuth(
		redirectUri,
		scope,
		code,
	)

	if auth == nil {
		zap.L().Error("Error getting auth from battle net")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Error getting auth"))
		return
	}

	userInfo := battlenethttp.GetInstance().GetBattleNetUserInfo(auth.AccessToken)
	if userInfo == nil {
		zap.L().Error("Error getting user info from battle net")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Error getting user info"))
		return
	}

	existing, err := userrepository.GetRepository().GetByBattleTag(userInfo.BattleTag)
	if err != nil {
		existing = &documents.UserDocument{
			ObjectID:    primitive.NewObjectID(),
			DisplayName: userInfo.BattleTag,
			Connections: &documents.UserConnections{
				BattleTag: userInfo.BattleTag,
			},
		}
		userrepository.GetRepository().Create(existing)
	}

	body, err := json.Marshal(&authorization.Authorization{
		Id:          existing.ObjectID.Hex(),
		DisplayName: existing.DisplayName,
		Connections: existing.Connections,
		Tokens: &authorization.AuthorizationTokens{
			AccessToken:  authorization.GetJwt(existing.ObjectID.Hex(), time.Hour*8),
			RefreshToken: authorization.GetJwt(existing.ObjectID.Hex(), time.Hour*48),
		},
	})
	if err != nil {
		zap.L().Error("Failed to stringify response body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Invalid JSON body"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	zap.L().Info("Responded with login response")
}
