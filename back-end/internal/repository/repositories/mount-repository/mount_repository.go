package mountrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	commonrepository "wowcollector.io/internal/repository/repositories/common-repository"
)

type MountRepository struct {
	commonrepository.CommonRepository
	collection *mongo.Collection
}

var instance *MountRepository

func GetRepository() *MountRepository {
	return instance
}

func Init(database *mongo.Database) {
	collection := database.Collection("mounts")
	instance = &MountRepository{
		CommonRepository: commonrepository.CommonRepository{
			Collection: collection,
		},
		collection: collection,
	}
	instance.createIndexes()
}

func (r *MountRepository) GetMounts() ([]*documents.MountDocument, error) {
	result, err := r.GetAll()
	if err != nil {
		zap.L().Info("Error fetching mounts" + err.Error())
		return nil, err
	}

	mounts := make([]*documents.MountDocument, len(result))
	for i, record := range result {
		var doc *documents.MountDocument
		bsonBytes, _ := bson.Marshal(record)
		bson.Unmarshal(bsonBytes, &doc)
		mounts[i] = doc
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
			{Key: "id", Value: 1},
		},
	}

	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		zap.L().Error(err.Error())
	}
}
