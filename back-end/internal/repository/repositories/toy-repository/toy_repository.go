package toyrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
)

type ToyRepository struct {
	collection *mongo.Collection
}

var instance *ToyRepository

func GetRepository() *ToyRepository {
	return instance
}

func Init(database *mongo.Database) {
	instance = &ToyRepository{collection: database.Collection("toys")}
	instance.createIndexes()
}

func (r *ToyRepository) GetToys() ([]*documents.ToyDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		zap.L().Info("Error fetching toys" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var toys []*documents.ToyDocument
	for result.Next(context.TODO()) {
		var toy *documents.ToyDocument
		err := result.Decode(&toy)
		if err != nil {
			zap.L().Info("Error decoding toy" + err.Error())
		}
		toys = append(toys, toy)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching toys" + err.Error())
	}
	return toys, nil
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

func (r *ToyRepository) createIndexes() {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "id", Value: 1},
		},
	}

	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		zap.L().Error(err.Error())
	}
}
