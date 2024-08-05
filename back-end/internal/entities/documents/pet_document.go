package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type PetDocument struct {
	ObjectID       primitive.ObjectID `json:"_id" bson:"_id"`
	Id             int                `json:"id" bson:"id"`
	Name           string             `json:"name" bson:"name"`
	Source         string             `json:"source" bson:"source"`
	IsUnobtainable bool               `json:"isUnobtainable" bson:"isUnobtainable"`
	Icon           string             `json:"icon" bson:"icon"`
}

func (d *PetDocument) IsEqual(other *PetDocument) bool {
	return d.Id == other.Id &&
		d.Name == other.Name &&
		d.Source == other.Source &&
		d.IsUnobtainable == other.IsUnobtainable &&
		d.Icon == other.Icon
}
