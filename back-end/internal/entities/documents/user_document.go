package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDocument struct {
	ObjectID       primitive.ObjectID   `bson:"_id" json:"_id"`
	BattleTag      string               `bson:"battleTag" json:"battleTag"`
}
