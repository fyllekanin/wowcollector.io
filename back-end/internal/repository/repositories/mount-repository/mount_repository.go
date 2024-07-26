package mountrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
)

type MountRepository struct {
	collection *mongo.Collection
}

var instance *MountRepository

func GetRepository() *MountRepository {
	return instance
}

func Init(database *mongo.Database) {
	instance = &MountRepository{collection: database.Collection("mounts")}
	instance.createIndexes()
}

func (r *MountRepository) GetMounts() ([]*documents.MountDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		zap.L().Info("Error fetching mounts" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var mounts []*documents.MountDocument
	for result.Next(context.TODO()) {
		var mount *documents.MountDocument
		err := result.Decode(&mount)
		if err != nil {
			zap.L().Info("Error decoding mount" + err.Error())
		}
		mounts = append(mounts, mount)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching mounts" + err.Error())
	}
	return mounts, nil
}

func (r *MountRepository) CreateMount(document *documents.MountDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting mount document:" + err.Error())
		return err
	}
	return nil
}

func (r *MountRepository) UpdateMount(document *documents.MountDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		zap.L().Info("Error updating mount document:" + err.Error())
		return err
	}
	return nil
}

func (r *MountRepository) createIndexes() {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "name", Value: 1},
		},
	}

	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		zap.L().Error(err.Error())
	}
}
