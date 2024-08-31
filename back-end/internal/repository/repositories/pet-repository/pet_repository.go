package petrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	commonrepository "wowcollector.io/internal/repository/repositories/common-repository"
)

type PetRepository struct {
	commonrepository.CommonRepository
	collection *mongo.Collection
}

var instance *PetRepository

func GetRepository() *PetRepository {
	return instance
}

func Init(database *mongo.Database) {
	collection := database.Collection("pets")
	instance = &PetRepository{
		CommonRepository: commonrepository.CommonRepository{
			Collection: collection,
		},
		collection: collection,
	}
	instance.CreateIndex("id")
}

func (r *PetRepository) GetPets() ([]*documents.PetDocument, error) {
	result, err := r.GetAll()
	if err != nil {
		zap.L().Info("Error fetching pets" + err.Error())
		return nil, err
	}

	items := make([]*documents.PetDocument, len(result))
	for i, record := range result {
		var doc *documents.PetDocument
		bsonBytes, _ := bson.Marshal(record)
		bson.Unmarshal(bsonBytes, &doc)
		items[i] = doc
	}
	return items, nil
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
