package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type MountDocument struct {
	ObjectID        primitive.ObjectID `bson:"_id"`
	Id              int                `bson:"id"`
	Name            string             `bson:"name"`
	Description     string             `bson:"description"`
	Source          string             `bson:"source"`
	Faction         string             `bson:"faction"`
	CreatureDisplay string             `bson:"creatureDisplay"`
	IsUnobtainable  bool               `bson:"isUnobtainable"`
	Icon            string             `bson:"icon"`
}

func (r *MountDocument) IsEqual(other *MountDocument) bool {
	return r.Id == other.Id &&
		r.Name == other.Name &&
		r.Description == other.Description &&
		r.Source == other.Source &&
		r.Faction == other.Faction &&
		r.CreatureDisplay == other.CreatureDisplay &&
		r.IsUnobtainable == other.IsUnobtainable &&
		r.Icon == other.Icon
}
