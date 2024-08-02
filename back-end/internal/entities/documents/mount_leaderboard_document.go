package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type LeaderboardDocument struct {
	ObjectID  primitive.ObjectID `bson:"_id" json:"_id"`
	Character string             `bson:"character" json:"character"`
	Realm     string             `bson:"realm" json:"realm"`
	Region    string             `bson:"region" json:"region"`
	Count     int                `bson:"count" json:"count"`
}
