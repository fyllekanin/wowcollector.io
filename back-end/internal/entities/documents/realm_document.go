package documents

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
)

type RealmDocument struct {
	ObjectID primitive.ObjectID           `bson:"_id" json:"_id"`
	Id       int                          `bson:"id" json:"id"`
	Name     string                       `bson:"name" json:"name"`
	Slug     string                       `bson:"slug" json:"slug"`
	Region   blizzarddata.BattleNetRegion `bson:"region" json:"region"`
}

func (r *RealmDocument) IsEqual(other *RealmDocument) bool {
	return r.Id == other.Id &&
		r.Name == other.Name &&
		r.Slug == other.Slug &&
		r.Region == other.Region
}
