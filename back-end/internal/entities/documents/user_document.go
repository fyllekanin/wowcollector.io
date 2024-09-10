package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserConnections struct {
	BattleTag string `bson:"battleTag" json:"battleTag"`
}

type UserDocument struct {
	ObjectID    primitive.ObjectID `bson:"_id" json:"_id"`
	DisplayName string             `bson:"displayName" json:"displayName"`
	Connections *UserConnections   `bson:"connections" json:"connections"`
}
