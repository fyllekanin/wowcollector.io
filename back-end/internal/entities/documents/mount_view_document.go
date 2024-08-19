package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type MountViewCategoryMount struct {
	Id    int `bson:"id" json:"id"`
	Order int `bson:"order" json:"order"`
}

type MountViewCategory struct {
	Id         string                   `bson:"id" json:"id"`
	Name       string                   `bson:"name" json:"name"`
	Order      int                      `bson:"order" json:"order"`
	Mounts     []MountViewCategoryMount `bson:"mounts" json:"mounts"`
	Categories []*MountViewCategory     `bson:"categories" json:"categories"`
}

type MountViewDocument struct {
	ObjectID          primitive.ObjectID   `bson:"_id" json:"_id"`
	Name              string               `bson:"name" json:"name"`
	IsDefault         bool                 `bson:"isDefault" json:"isDefault"`
	IsUnknownIncluded bool                 `bson:"isUnknownIncluded" json:"isUnknownIncluded"`
	Categories        []*MountViewCategory `json:"categories"`
}
