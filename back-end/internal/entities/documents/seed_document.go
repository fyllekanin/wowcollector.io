package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type SeedDocument struct {
	ObjectID primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
}
