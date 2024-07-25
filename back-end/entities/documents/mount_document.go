package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type MountDocument struct {
	ObjectID        primitive.ObjectID `bson:"_id" json:"_id"`
	Id              int                `bson:"id json:"id"`
	Name            string             `bson:"name" json:"name"`
	Description     string             `bson:"description" json:"description"`
	Source          string             `bson:"source" json:"source"`
	Faction         string             `bson:"faction" json:"faction"`
	CreatureDisplay string             `bson:"creatureDisplay" json:"creatureDisplay"`
	IsUnobtainable  bool               `bson:"isUnobtainable" json:"isUnobtainable"`
	Icon            string             `bson:"icon" json:"icon"`
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
