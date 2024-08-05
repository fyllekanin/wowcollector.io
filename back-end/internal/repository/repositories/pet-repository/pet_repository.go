package petrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
)

type PetRepository struct {
	collection *mongo.Collection
}

var instance *PetRepository

func GetRepository() *PetRepository {
	return instance
}

func Init(database *mongo.Database) {
	instance = &PetRepository{collection: database.Collection("pets")}
	instance.createIndexes()
}

func (r *PetRepository) GetPets() ([]*documents.PetDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		zap.L().Info("Error fetching pets" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var pets []*documents.PetDocument
	for result.Next(context.TODO()) {
		var pet *documents.PetDocument
		err := result.Decode(&pet)
		if err != nil {
			zap.L().Info("Error decoding pet" + err.Error())
		}
		pets = append(pets, pet)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching pets" + err.Error())
	}
	return pets, nil
}

func (r *PetRepository) CreatePet(document *documents.PetDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting pet document:" + err.Error())
		return err
	}
	return nil
}

func (r *PetRepository) UpdatePet(document *documents.PetDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		zap.L().Info("Error updating pet document:" + err.Error())
		return err
	}
	return nil
}

func (r *PetRepository) createIndexes() {
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
