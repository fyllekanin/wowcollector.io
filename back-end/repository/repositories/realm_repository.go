package realmrepository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"wowcollector.io/entities/documents"
)

type RealmRepository struct {
	collection *mongo.Collection
}

var instance *RealmRepository

func GetRepository() *RealmRepository {
	return instance
}

func Init(database *mongo.Database) {
	instance = &RealmRepository{collection: database.Collection("realms")}
}

func (r *RealmRepository) CreateRealm(document *documents.RealmDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		fmt.Println("Error inserting document:", err)
		return err
	}
	return nil
}
