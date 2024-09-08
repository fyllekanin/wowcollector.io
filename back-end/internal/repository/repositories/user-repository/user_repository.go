package userrepository

import (
	"go.mongodb.org/mongo-driver/mongo"
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
