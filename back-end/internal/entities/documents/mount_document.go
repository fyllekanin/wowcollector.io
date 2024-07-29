package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type MountDocumentAssets struct {
	Display   string `bson:"display" json:"display"`
	SmallIcon string `bson:"smallIcon" json:"smallIcon"`
	LargeIcon string `bson:"largeIcon"json:"largeIcon"`
}

type MountDocument struct {
	ObjectID       primitive.ObjectID   `bson:"_id" json:"_id"`
	Id             int                  `bson:"id" json:"id"`
	Name           string               `bson:"name" json:"name"`
	Description    string               `bson:"description" json:"description"`
	Source         string               `bson:"source" json:"source"`
	Faction        string               `bson:"faction" json:"faction"`
	IsUnobtainable bool                 `bson:"isUnobtainable" json:"isUnobtainable"`
	Assets         *MountDocumentAssets `bson:"assets" json:"assets"`
}

func (r *MountDocument) IsEqual(other *MountDocument) bool {
	return r.Id == other.Id &&
		r.Name == other.Name &&
		r.Description == other.Description &&
		r.Source == other.Source &&
		r.Faction == other.Faction &&
		r.Assets.Display == other.Assets.Display &&
		r.Assets.SmallIcon == other.Assets.SmallIcon &&
		r.Assets.LargeIcon == other.Assets.LargeIcon
}
