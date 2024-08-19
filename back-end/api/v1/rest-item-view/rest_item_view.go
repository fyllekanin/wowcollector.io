package restitemview

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	errorcodes "wowcollector.io/internal/common/error-codes"
	"wowcollector.io/internal/entities/documents"
	httprequests "wowcollector.io/internal/entities/http-requests"
	errorresponse "wowcollector.io/internal/entities/response/error"
	mountviewrepository "wowcollector.io/internal/repository/repositories/mount-view-repository"
	petviewrepository "wowcollector.io/internal/repository/repositories/pet-view-repository"
	toyviewrepository "wowcollector.io/internal/repository/repositories/toy-view-repository"
)

func GetRoutes(r chi.Router) {
	r.Route("/item-view", func(r chi.Router) {
		r.Post("/mount", postMountView)
		r.Post("/pet", postPetView)
		r.Post("/toy", postToyView)
	})
}

// @summary Create mount view
// @description Create an mount view
// @tags ItemView
// @accept json
// @Param view body httprequests.MountViewRequest true "View JSON"
// @produce json
// @success 201 {string} string
// @failure 400 {object} errorresponse.ErrorResponse
// @router /api/v1/item-view/mount [post]
func postMountView(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Creating an mount view")

	var view httprequests.MountViewRequest

	if err := json.NewDecoder(r.Body).Decode(&view); err != nil {
		zap.L().Error("Error parsing request body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Invalid JSON body"))
		return
	}

	document := documents.MountViewDocument{
		ObjectID:          primitive.NewObjectID(),
		Name:              view.Name,
		IsDefault:         false,
		IsUnknownIncluded: view.IsUnknownIncluded,
		Categories:        view.Categories,
	}

	err := mountviewrepository.GetRepository().CreateMountView(&document)
	if err != nil {
		zap.L().Error("Error creating mount view")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Creating mount view"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(document.ObjectID.Hex()))
}

// @summary Create toy view
// @description Create an toy view
// @tags ItemView
// @accept json
// @Param view body httprequests.ToyViewRequest true "View JSON"
// @produce json
// @success 201 {string} string
// @failure 400 {object} errorresponse.ErrorResponse
// @router /api/v1/item-view/toy [post]
func postToyView(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Creating an toy view")

	var view httprequests.ToyViewRequest

	if err := json.NewDecoder(r.Body).Decode(&view); err != nil {
		zap.L().Error("Error parsing request body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Invalid JSON body"))
		return
	}

	document := documents.ToyViewDocument{
		ObjectID:          primitive.NewObjectID(),
		Name:              view.Name,
		IsDefault:         false,
		IsUnknownIncluded: view.IsUnknownIncluded,
		Categories:        view.Categories,
	}

	err := toyviewrepository.GetRepository().CreateToyView(&document)
	if err != nil {
		zap.L().Error("Error creating toy view")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Creating toy view"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(document.ObjectID.Hex()))
}

// @summary Create pet view
// @description Create an pet view
// @tags ItemView
// @accept json
// @Param view body httprequests.PetViewRequest true "View JSON"
// @produce json
// @success 201 {string} string
// @failure 400 {object} errorresponse.ErrorResponse
// @router /api/v1/item-view/pet [post]
func postPetView(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Creating an pet view")

	var view httprequests.PetViewRequest

	if err := json.NewDecoder(r.Body).Decode(&view); err != nil {
		zap.L().Error("Error parsing request body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Invalid JSON body"))
		return
	}

	document := documents.PetViewDocument{
		ObjectID:          primitive.NewObjectID(),
		Name:              view.Name,
		IsDefault:         false,
		IsUnknownIncluded: view.IsUnknownIncluded,
		Categories:        view.Categories,
	}

	err := petviewrepository.GetRepository().CreatePetView(&document)
	if err != nil {
		zap.L().Error("Error creating pet view")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Creating pet view"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(document.ObjectID.Hex()))
}
