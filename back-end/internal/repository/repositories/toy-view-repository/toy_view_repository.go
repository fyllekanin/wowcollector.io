package toyviewrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
)

type ToyViewRepository struct {
	collection *mongo.Collection
}

var instance *ToyViewRepository

func GetRepository() *ToyViewRepository {
	return instance
}

func Init(database *mongo.Database) {
	instance = &ToyViewRepository{collection: database.Collection("toy-views")}
	instance.createNameIndex()
	instance.createIsDefaultIndex()
}

func (r *ToyViewRepository) GetToyViews() ([]*documents.ToyViewDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		zap.L().Info("Error fetching toy view" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var toys []*documents.ToyViewDocument
	for result.Next(context.TODO()) {
		var toy *documents.ToyViewDocument
		err := result.Decode(&toy)
		if err != nil {
			zap.L().Info("Error decoding toy view" + err.Error())
		}
		toys = append(toys, toy)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching toys" + err.Error())
	}
	return toys, nil
}

func (r *ToyViewRepository) CreateToyView(document *documents.ToyViewDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting toy view document:" + err.Error())
		return err
	}
	return nil
}

func (r *ToyViewRepository) GetDefaultToyView() (*documents.ToyViewDocument, error) {
	filter := bson.D{{"isDefault", true}}
	var ToyView *documents.ToyViewDocument

	err := r.collection.FindOne(context.TODO(), filter).Decode(&ToyView)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			zap.L().Info("No default toy view found")
			return nil, nil
		}
		zap.L().Info("Error fetching default toy view" + err.Error())
		return nil, err
	}

	return ToyView, nil
}

func (r *ToyViewRepository) UpdateToyView(document *documents.ToyViewDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		zap.L().Info("Error updating toy view document:" + err.Error())
		return err
	}
	return nil
}

func (r *ToyViewRepository) createNameIndex() {
	nameIndexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "name", Value: 1},
		},
	}

	_, err := r.collection.Indexes().CreateOne(context.TODO(), nameIndexModel)
	if err != nil {
		zap.L().Error(err.Error())
	}
}

func (r *ToyViewRepository) createIsDefaultIndex() {
	isDefaultIndexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "isDefault", Value: 1},
		},
	}

	_, err := r.collection.Indexes().CreateOne(context.TODO(), isDefaultIndexModel)
	if err != nil {
		zap.L().Error(err.Error())
	}
}
