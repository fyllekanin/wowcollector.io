package mountviewrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
)

type MountViewRepository struct {
	collection *mongo.Collection
}

var instance *MountViewRepository

func GetRepository() *MountViewRepository {
	return instance
}

func Init(database *mongo.Database) {
	instance = &MountViewRepository{collection: database.Collection("mount-views")}
	instance.createNameIndex()
	instance.createIsDefaultIndex()
}

func (r *MountViewRepository) GetMountViews() ([]*documents.MountViewDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		zap.L().Info("Error fetching mount view" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var mounts []*documents.MountViewDocument
	for result.Next(context.TODO()) {
		var mount *documents.MountViewDocument
		err := result.Decode(&mount)
		if err != nil {
			zap.L().Info("Error decoding mount view" + err.Error())
		}
		mounts = append(mounts, mount)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching mounts" + err.Error())
	}
	return mounts, nil
}

func (r *MountViewRepository) CreateMountView(document *documents.MountViewDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting mount view document:" + err.Error())
		return err
	}
	return nil
}

func (r *MountViewRepository) GetDefaultMountView() (*documents.MountViewDocument, error) {
	filter := bson.D{{"isDefault", true}}
	var mountView *documents.MountViewDocument

	err := r.collection.FindOne(context.TODO(), filter).Decode(&mountView)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			zap.L().Info("No default mount view found")
			return nil, nil
		}
		zap.L().Error("Error fetching default mount view" + err.Error())
		return nil, err
	}

	return mountView, nil
}

func (r *MountViewRepository) GetMountView(id string) (*documents.MountViewDocument, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		zap.L().Error("Error parsing id to ObjectID: " + id)
		return nil, err
	}
	filter := bson.D{{"_id", objId}}
	var mountView *documents.MountViewDocument

	findErr := r.collection.FindOne(context.TODO(), filter).Decode(&mountView)
	if findErr != nil {
		if findErr == mongo.ErrNoDocuments {
			zap.L().Info("No mount view found with id: " + id)
			return nil, findErr
		}
		zap.L().Error("Error fetching mount view" + err.Error())
		return nil, findErr
	}

	return mountView, nil
}

func (r *MountViewRepository) UpdateMountView(document *documents.MountViewDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		zap.L().Info("Error updating mount view document:" + err.Error())
		return err
	}
	return nil
}

func (r *MountViewRepository) createNameIndex() {
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

func (r *MountViewRepository) createIsDefaultIndex() {
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
