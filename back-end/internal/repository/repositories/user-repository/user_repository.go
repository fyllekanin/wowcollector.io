package userrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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
	instance.CreateIndex("connections.battleTag")
}

func (r *UserRepository) Create(document *documents.UserDocument) error {
	_, err := r.Collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting user document:" + err.Error())
		return err
	}
	return nil
}

func (r *UserRepository) GetByBattleTag(battleTag string) (*documents.UserDocument, error) {
	filter := bson.D{{"connections.battleTag", battleTag}}
	var result *documents.UserDocument

	err := r.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		zap.L().Info("Error finding user document by battleTag:" + err.Error())
		return nil, err
	}
	return result, nil
}
