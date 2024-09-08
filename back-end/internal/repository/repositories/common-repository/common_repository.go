package commonrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Repository interface {
	GetAll() ([]bson.D, error)
}

type CommonRepository struct {
	Collection *mongo.Collection
}

func (r *CommonRepository) GetAll() ([]bson.D, error) {
	result, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer result.Close(context.TODO())

	var records []bson.D
	if err = result.All(context.TODO(), &records); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *CommonRepository) Create(document *bson.D) error {
	_, err := r.Collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting document:" + err.Error())
		return err
	}
	return nil
}

func (r *CommonRepository) CreateIndex(key string) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: key, Value: 1},
		},
	}

	_, err := r.Collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		zap.L().Error(err.Error())
	}
}

func (r *CommonRepository) CreateCombinedIndex(keys []string) {
	indexKeys := bson.D{}
	for _, key := range keys {
		indexKeys = append(indexKeys, bson.E{Key: key, Value: 1})
	}

	indexModel := mongo.IndexModel{
		Keys: indexKeys,
	}

	_, err := r.Collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		zap.L().Error(err.Error())
	}
}
