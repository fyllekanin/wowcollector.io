package userrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	commonrepository "wowcollector.io/internal/repository/repositories/common-repository"
)

type UserRepository struct {
	commonrepository.CommonRepository
	collection *mongo.Collection
}

var instance *UserRepository

func GetRepository() *UserRepository {
	return instance
}

func Init(database *mongo.Database) {
	collection := database.Collection("users")
	instance = &UserRepository{
		CommonRepository: commonrepository.CommonRepository{
			Collection: collection,
		},
		collection: collection,
	}
	instance.CreateIndex("battleTag")
}

func (r *UserRepository) Create(document *documents.UserDocument) error {
	_, err := r.Collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting user document:" + err.Error())
		return err
	}
	return nil
}
