package petviewrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	commonrepository "wowcollector.io/internal/repository/repositories/common-repository"
)

type PetViewRepository struct {
	commonrepository.CommonRepository
	collection *mongo.Collection
}

var instance *PetViewRepository

func GetRepository() *PetViewRepository {
	return instance
}

func Init(database *mongo.Database) {
	collection := database.Collection("pet-views")
	instance = &PetViewRepository{
		CommonRepository: commonrepository.CommonRepository{
			Collection: collection,
		},
		collection: collection,
	}
	instance.CreateIndex("isDefault")
}

func (r *PetViewRepository) GetPetViews() ([]*documents.PetViewDocument, error) {
	result, err := r.GetAll()
	if err != nil {
		zap.L().Info("Error fetching pet views" + err.Error())
		return nil, err
	}

	items := make([]*documents.PetViewDocument, len(result))
	for i, record := range result {
		var doc *documents.PetViewDocument
		bsonBytes, _ := bson.Marshal(record)
		bson.Unmarshal(bsonBytes, &doc)
		items[i] = doc
	}
	return items, nil
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

func (r *PetViewRepository) GetPetView(id string) (*documents.PetViewDocument, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		zap.L().Error("Error parsing id to ObjectID: " + id)
		return nil, err
	}
	filter := bson.D{{"_id", objId}}
	var petView *documents.PetViewDocument

	findErr := r.collection.FindOne(context.TODO(), filter).Decode(&petView)
	if findErr != nil {
		if findErr == mongo.ErrNoDocuments {
			zap.L().Info("No pet view found with id: " + id)
			return nil, findErr
		}
		zap.L().Error("Error fetching pet view" + err.Error())
		return nil, findErr
	}

	return petView, nil
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
