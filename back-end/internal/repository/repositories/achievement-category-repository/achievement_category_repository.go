package achievementcategoryrepository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
)

type AchievementCategoryRepository struct {
	collection *mongo.Collection
}

var instance *AchievementCategoryRepository

func GetRepository() *AchievementCategoryRepository {
	return instance
}

func Init(database *mongo.Database) {
	instance = &AchievementCategoryRepository{collection: database.Collection("achievement-categories")}
	instance.createIndexes("id")
	instance.createIndexes("isRootCategory")
	instance.createIndexes("rootCategoryId")
}

func (r *AchievementCategoryRepository) GetAchievementRootCategories() ([]*documents.AchievementCategoryDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{{"isRootCategory", true}})
	if err != nil {
		zap.L().Info("Error fetching achievement root categories" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var achievementCategories []*documents.AchievementCategoryDocument
	for result.Next(context.TODO()) {
		var achievementCategory *documents.AchievementCategoryDocument
		err := result.Decode(&achievementCategory)
		if err != nil {
			zap.L().Info("Error decoding achievement root category" + err.Error())
		}
		achievementCategories = append(achievementCategories, achievementCategory)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching achievement root categories" + err.Error())
	}
	return achievementCategories, nil
}

func (r *AchievementCategoryRepository) GetAchievementCategoryWithId(id int) (*documents.AchievementCategoryDocument, error) {
	filter := bson.D{{"id", id}}
	var result *documents.AchievementCategoryDocument

	err := r.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			zap.L().Info(fmt.Sprintf("No category with id %d", id))
			return nil, nil
		}
		zap.L().Error("Error fetching category" + err.Error())
		return nil, err
	}
	return result, nil
}

func (r *AchievementCategoryRepository) GetAchievementCategories() ([]*documents.AchievementCategoryDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		zap.L().Info("Error fetching achievement categories" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var achievementCategories []*documents.AchievementCategoryDocument
	for result.Next(context.TODO()) {
		var achievementCategory *documents.AchievementCategoryDocument
		err := result.Decode(&achievementCategory)
		if err != nil {
			zap.L().Info("Error decoding achievement category" + err.Error())
		}
		achievementCategories = append(achievementCategories, achievementCategory)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching achievement categories" + err.Error())
	}
	return achievementCategories, nil
}

func (r *AchievementCategoryRepository) GetAchievementCategoriesForRootCategory(rootCategoryId int) ([]*documents.AchievementCategoryDocument, error) {
	zap.L().Info("test")
	result, err := r.collection.Find(context.TODO(), bson.D{{"rootCategoryId", rootCategoryId}})
	if err != nil {
		zap.L().Info("Error fetching achievement categories" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var achievementCategories []*documents.AchievementCategoryDocument
	for result.Next(context.TODO()) {
		var achievementCategory *documents.AchievementCategoryDocument
		err := result.Decode(&achievementCategory)
		if err != nil {
			zap.L().Info("Error decoding achievement category" + err.Error())
		}
		achievementCategories = append(achievementCategories, achievementCategory)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching achievement categories" + err.Error())
	}
	return achievementCategories, nil
}

func (r *AchievementCategoryRepository) CreateAchievementCategory(document *documents.AchievementCategoryDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting achievement category document:" + err.Error())
		return err
	}
	return nil
}

func (r *AchievementCategoryRepository) UpdateAchievementCategory(document *documents.AchievementCategoryDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		zap.L().Info("Error updating achievement category document:" + err.Error())
		return err
	}
	return nil
}

func (r *AchievementCategoryRepository) createIndexes(key string) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: key, Value: 1},
		},
	}

	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		zap.L().Error(err.Error())
	}
}
