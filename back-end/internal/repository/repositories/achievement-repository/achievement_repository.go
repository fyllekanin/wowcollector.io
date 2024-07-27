package achievementrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
)

type AchievementRepository struct {
	collection *mongo.Collection
}

var instance *AchievementRepository

func GetRepository() *AchievementRepository {
	return instance
}

func Init(database *mongo.Database) {
	instance = &AchievementRepository{collection: database.Collection("achievements")}
	instance.createIndexes()
}

func (r *AchievementRepository) GetAchievements() ([]*documents.AchievementDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		zap.L().Info("Error fetching achievements" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var achievements []*documents.AchievementDocument
	for result.Next(context.TODO()) {
		var achievement *documents.AchievementDocument
		err := result.Decode(&achievement)
		if err != nil {
			zap.L().Info("Error decoding achievement" + err.Error())
		}
		achievements = append(achievements, achievement)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching achievements" + err.Error())
	}
	return achievements, nil
}

func (r *AchievementRepository) CreateAchievement(document *documents.AchievementDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting achievement document:" + err.Error())
		return err
	}
	return nil
}

func (r *AchievementRepository) UpdateAchievement(document *documents.AchievementDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		zap.L().Info("Error updating achievement document:" + err.Error())
		return err
	}
	return nil
}

func (r *AchievementRepository) createIndexes() {
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
