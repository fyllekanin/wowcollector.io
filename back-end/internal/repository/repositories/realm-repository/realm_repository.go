package realmrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	commonrepository "wowcollector.io/internal/repository/repositories/common-repository"
)

type RealmRepository struct {
	commonrepository.CommonRepository
	collection *mongo.Collection
}

var instance *RealmRepository

func GetRepository() *RealmRepository {
	return instance
}

func Init(database *mongo.Database) {
	collection := database.Collection("realms")
	instance = &RealmRepository{
		CommonRepository: commonrepository.CommonRepository{
			Collection: collection,
		},
		collection: collection,
	}
	instance.CreateIndex("region")
}

func (r *RealmRepository) GetRealms() ([]*documents.RealmDocument, error) {
	result, err := r.GetAll()
	if err != nil {
		zap.L().Info("Error fetching realms" + err.Error())
		return nil, err
	}

	items := make([]*documents.RealmDocument, len(result))
	for i, record := range result {
		var doc *documents.RealmDocument
		bsonBytes, _ := bson.Marshal(record)
		bson.Unmarshal(bsonBytes, &doc)
		items[i] = doc
	}
	return items, nil
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
