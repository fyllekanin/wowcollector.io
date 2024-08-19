package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type PetViewCategoryPet struct {
	Id    int `bson:"id" json:"id"`
	Order int `bson:"order" json:"order"`
}

type PetViewCategory struct {
	Id         string               `bson:"id" json:"id"`
	Name       string               `bson:"name" json:"name"`
	Order      int                  `bson:"order" json:"order"`
	Pets       []PetViewCategoryPet `bson:"pets" json:"pets"`
	Categories []PetViewCategory    `bson:"categories" json:"categories"`
}

type PetViewDocument struct {
	ObjectID          primitive.ObjectID `bson:"_id" json:"_id"`
	Name              string             `bson:"name" json:"name"`
	IsDefault         bool               `bson:"isDefault" json:"isDefault"`
	IsUnknownIncluded bool               `bson:"isUnknownIncluded" json:"isUnknownIncluded"`
	Categories        []PetViewCategory  `json:"categories"`
}
