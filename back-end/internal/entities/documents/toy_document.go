package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToyDocument struct {
	ObjectID       primitive.ObjectID `bson:"_id"`
	Id             int                `bson:"id"`
	ItemId         int                `bson:"itemId"`
	Name           string             `bson:"name"`
	Source         string             `bson:"source"`
	IsUnobtainable bool               `bson:"isUnobtainable"`
	Icon           string             `bson:"icon"`
}

func (d *ToyDocument) IsEqual(other *ToyDocument) bool {
	return d.Id == other.Id &&
		d.ItemId == other.ItemId &&
		d.Name == other.Name &&
		d.Source == other.Source &&
		d.IsUnobtainable == other.IsUnobtainable &&
		d.Icon == other.Icon
}
