package petviewrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
)

type PetViewRepository struct {
	collection *mongo.Collection
}

var instance *PetViewRepository

func GetRepository() *PetViewRepository {
	return instance
}

func Init(database *mongo.Database) {
	instance = &PetViewRepository{collection: database.Collection("pet-views")}
	instance.createNameIndex()
	instance.createIsDefaultIndex()
}

func (r *PetViewRepository) GetPetViews() ([]*documents.PetViewDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		zap.L().Info("Error fetching pet view" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var pets []*documents.PetViewDocument
	for result.Next(context.TODO()) {
		var pet *documents.PetViewDocument
		err := result.Decode(&pet)
		if err != nil {
			zap.L().Info("Error decoding pet view" + err.Error())
		}
		pets = append(pets, pet)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching pets" + err.Error())
	}
	return pets, nil
}

func (r *PetViewRepository) CreatePetView(document *documents.PetViewDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting pet view document:" + err.Error())
		return err
	}
	return nil
}

func (r *PetViewRepository) GetDefaultPetView() (*documents.PetViewDocument, error) {
	filter := bson.D{{"isDefault", true}}
	var PetView *documents.PetViewDocument

	err := r.collection.FindOne(context.TODO(), filter).Decode(&PetView)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			zap.L().Info("No default pet view found")
			return nil, nil
		}
		zap.L().Info("Error fetching default pet view" + err.Error())
		return nil, err
	}

	return PetView, nil
}

func (r *PetViewRepository) UpdatePetView(document *documents.PetViewDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		zap.L().Info("Error updating pet view document:" + err.Error())
		return err
	}
	return nil
}

func (r *PetViewRepository) createNameIndex() {
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

func (r *PetViewRepository) createIsDefaultIndex() {
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
