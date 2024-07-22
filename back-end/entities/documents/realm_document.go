package documents

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"wowcollector.io/common/data"
)

type RealmDocument struct {
	ObjectID primitive.ObjectID   `bson:"_id"`
	Id       int                  `bson:"id"`
	Name     string               `bson:"name"`
	Slug     string               `bson:"slug"`
	Region   data.BattleNetRegion `bson:"region"`
}

func (r *RealmDocument) IsEqual(other *RealmDocument) bool {
	return r.Id == other.Id &&
		r.Name == other.Name &&
		r.Slug == other.Slug &&
		r.Region == other.Region
}
