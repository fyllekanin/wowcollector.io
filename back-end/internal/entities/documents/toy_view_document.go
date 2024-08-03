package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToyViewCategoryToy struct {
	Id int `bson:"id" json:"id"`
}

type ToyViewCategory struct {
	Name       string               `bson:"name" json:"name"`
	Order      int                  `bson:"order" json:"order"`
	Toys       []ToyViewCategoryToy `bson:"toys" json:"toys"`
	Categories []ToyViewCategory    `bson:"categories" json:"categories"`
}

type ToyViewDocument struct {
	ObjectID          primitive.ObjectID `bson:"_id" json:"_id"`
	Name              string             `bson:"name" json:"name"`
	IsDefault         bool               `bson:"isDefault" json:"isDefault"`
	IsUnknownIncluded bool               `bson:"isUnknownIncluded" json:"isUnknownIncluded"`
	Categories        []ToyViewCategory  `json:"categories"`
}
