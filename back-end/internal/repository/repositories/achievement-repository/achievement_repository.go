package achievementrepository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	commonrepository "wowcollector.io/internal/repository/repositories/common-repository"
)

type AchievementRepository struct {
	commonrepository.CommonRepository
	collection *mongo.Collection
}

var instance *AchievementRepository

func GetRepository() *AchievementRepository {
	return instance
}

func Init(database *mongo.Database) {
	collection := database.Collection("achievements")
	instance = &AchievementRepository{
		CommonRepository: commonrepository.CommonRepository{
			Collection: collection,
		},
		collection: collection,
	}
	instance.CreateIndex("id")
	instance.CreateIndex("categoryId")
}

func (r *AchievementRepository) GetTotal() (int64, error) {
	result, err := r.collection.EstimatedDocumentCount(context.TODO())
	if err != nil {
		zap.L().Error("Error counting achievement documents")
		return 0, err
	}
	return result, nil
}

func (r *AchievementRepository) GetAchievements() ([]*documents.AchievementDocument, error) {
	result, err := r.GetAll()
	if err != nil {
		zap.L().Info("Error fetching achievements" + err.Error())
		return nil, err
	}

	items := make([]*documents.AchievementDocument, len(result))
	for i, record := range result {
		var doc *documents.AchievementDocument
		bsonBytes, _ := bson.Marshal(record)
		bson.Unmarshal(bsonBytes, &doc)
		items[i] = doc
	}
	return items, nil
}

func (r *AchievementRepository) GetAchievementsForCategoryId(categoryId int) ([]*documents.AchievementDocument, error) {
	filter := bson.D{{"categoryId", categoryId}}
	result, err := r.collection.Find(context.TODO(), filter)

	if err != nil {
		zap.L().Error(fmt.Sprintf("Error fetching achievements for category %d", categoryId))
		return nil, err
	}
	defer result.Close(context.TODO())

	var achievements []*documents.AchievementDocument
	for result.Next(context.TODO()) {
		var achievement *documents.AchievementDocument
		err := result.Decode(&achievement)
		if err != nil {
			zap.L().Error("Error decoding achievement" + err.Error())
		}
		achievements = append(achievements, achievement)
	}

	if err := result.Err(); err != nil {
		zap.L().Error("Error fetching achievements for category" + err.Error())
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
