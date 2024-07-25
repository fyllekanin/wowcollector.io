package realmrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
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
	instance.createIndexes()
}

func (r *RealmRepository) GetRealms() ([]*documents.RealmDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		zap.L().Info("Error fetching realms" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var realms []*documents.RealmDocument
	for result.Next(context.TODO()) {
		var realm *documents.RealmDocument
		err := result.Decode(&realm)
		if err != nil {
			zap.L().Info("Error decoding realm" + err.Error())
		}
		realms = append(realms, realm)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching realms" + err.Error())
	}
	return realms, nil
}

func (r *RealmRepository) CreateRealm(document *documents.RealmDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting realm document:" + err.Error())
		return err
	}
	return nil
}

func (r *RealmRepository) UpdateRealm(document *documents.RealmDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		zap.L().Info("Error updating realm document:" + err.Error())
		return err
	}
	return nil
}

func (r *RealmRepository) createIndexes() {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "region", Value: 1},
			{Key: "slug", Value: 1},
		},
	}

	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		zap.L().Error(err.Error())
	}
}
