package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToyDocument struct {
	ObjectID       primitive.ObjectID `json:"_id" bson:"_id"`
	Id             int                `json:"id" bson:"id"`
	ItemId         int                `json:"itemId" bson:"itemId"`
	Name           string             `json:"name" bson:"name"`
	Source         string             `json:"source" bson:"source"`
	IsUnobtainable bool               `json:"isUnobtainable" bson:"isUnobtainable"`
	Icon           string             `json:"icon" bson:"icon"`
}

func (d *ToyDocument) IsEqual(other *ToyDocument) bool {
	return d.Id == other.Id &&
		d.ItemId == other.ItemId &&
		d.Name == other.Name &&
		d.Source == other.Source &&
		d.IsUnobtainable == other.IsUnobtainable &&
		d.Icon == other.Icon
}
