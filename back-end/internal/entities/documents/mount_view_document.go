package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type MountViewCategoryMount struct {
	Id int `json:"id"`
}

type MountViewCategory struct {
	Name       string                   `json:"name"`
	Order      int                      `json:"Order"`
	Mounts     []MountViewCategoryMount `json:"mounts"`
	Categories []MountViewCategory      `json:"categories"`
}

type MountViewDocument struct {
	ObjectID          primitive.ObjectID  `bson:"_id" json:"_id"`
	Name              string              `bson:"name" json:"name"`
	IsDefault         bool                `bson:"isDefault" json:"isDefault"`
	IsUnknownIncluded bool                `bson:"isUnknownIncluded" json:"isUnknownIncluded"`
	Categories        []MountViewCategory `json:"categories"`
}
