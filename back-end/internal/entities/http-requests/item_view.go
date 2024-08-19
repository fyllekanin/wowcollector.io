package httprequests

import "wowcollector.io/internal/entities/documents"

type MountViewRequest struct {
	Name              string                         `json:"name"`
	Categories        []*documents.MountViewCategory `json:"categories"`
	IsUnknownIncluded bool                           `json:"isUnknownIncluded"`
}

type PetViewRequest struct {
	Name              string                      `json:"name"`
	Type              string                      `json:"type"`
	Categories        []documents.PetViewCategory `json:"categories"`
	IsUnknownIncluded bool                        `json:"isUnknownIncluded"`
}

type ToyViewRequest struct {
	Name              string                      `json:"name"`
	Type              string                      `json:"type"`
	Categories        []documents.ToyViewCategory `json:"categories"`
	IsUnknownIncluded bool                        `json:"isUnknownIncluded"`
}
