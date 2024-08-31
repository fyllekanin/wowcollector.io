package seedrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	commonrepository "wowcollector.io/internal/repository/repositories/common-repository"
)

type SeedRepository struct {
	commonrepository.CommonRepository
	collection *mongo.Collection
}

var instance *SeedRepository

func GetRepository() *SeedRepository {
	return instance
}

func Init(database *mongo.Database) {
	collection := database.Collection("seeds")
	instance = &SeedRepository{
		CommonRepository: commonrepository.CommonRepository{
			Collection: collection,
		},
		collection: collection,
	}
	instance.CreateIndex("name")
}

func (r *SeedRepository) CreateSeed(name string) error {
	_, err := r.collection.InsertOne(context.TODO(), &documents.SeedDocument{
		ObjectID: primitive.NewObjectID(),
		Name:     name,
	})
	if err != nil {
		zap.L().Info("Error inserting seed document:" + err.Error())
		return err
	}
	return nil
}

func (r *SeedRepository) IsExisting(name string) (bool, error) {
	filter := bson.D{{"name", name}}

	var result documents.SeedDocument
	err := r.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		zap.L().Info("Error finding seed document:" + err.Error())
		return false, err
	}
	return true, nil
}
