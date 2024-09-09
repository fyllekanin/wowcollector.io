package restauth

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	errorcodes "wowcollector.io/internal/common/error-codes"
	"wowcollector.io/internal/entities/documents"
	"wowcollector.io/internal/entities/response"
	errorresponse "wowcollector.io/internal/entities/response/error"
	userrepository "wowcollector.io/internal/repository/repositories/user-repository"
	battlenethttp "wowcollector.io/internal/services/http/battle-net-http"
)

func GetRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/battle-net", getBattleNetAuth)
	})
}

// @summary Login with battle net auth code
// @description Performs login for battle net auth code
// @tags Auth
// @produce json
// @param code query string true "Auth code"
// @param redirectUri query string true "Redirect Uri"
// @param scope query string true "Scope"
// @success 200 {object} []response.LoginResponse
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

	userInfo := battlenethttp.GetInstance().GetBattleNetUserInfo(auth.AccessCode)
	if userInfo == nil {
		zap.L().Error("Error getting user info from battle net")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Error getting user info"))
		return
	}

	userrepository.GetRepository().Create(&documents.UserDocument{
		ObjectID:  primitive.NewObjectID(),
		BattleTag: userInfo.BattleTag,
	})

	body, err := json.Marshal(&response.LoginResponse{
		AccessToken:  "",
		RefreshToken: "",
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
