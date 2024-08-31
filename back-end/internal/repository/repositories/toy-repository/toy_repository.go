package toyrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	commonrepository "wowcollector.io/internal/repository/repositories/common-repository"
)

type ToyRepository struct {
	commonrepository.CommonRepository
	collection *mongo.Collection
}

var instance *ToyRepository

func GetRepository() *ToyRepository {
	return instance
}

func Init(database *mongo.Database) {
	collection := database.Collection("toys")
	instance = &ToyRepository{
		CommonRepository: commonrepository.CommonRepository{
			Collection: collection,
		},
		collection: collection,
	}
	instance.CreateIndex("id")
}

func (r *ToyRepository) GetToys() ([]*documents.ToyDocument, error) {
	result, err := r.GetAll()
	if err != nil {
		zap.L().Info("Error fetching toys" + err.Error())
		return nil, err
	}

	items := make([]*documents.ToyDocument, len(result))
	for i, record := range result {
		var doc *documents.ToyDocument
		bsonBytes, _ := bson.Marshal(record)
		bson.Unmarshal(bsonBytes, &doc)
		items[i] = doc
	}
	return items, nil
}

func (r *ToyRepository) CreateToy(document *documents.ToyDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting toy document:" + err.Error())
		return err
	}
	return nil
}

func (r *ToyRepository) UpdateToy(document *documents.ToyDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		zap.L().Info("Error updating toy document:" + err.Error())
		return err
	}
	return nil
}
