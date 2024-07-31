package mountleaderboardrepository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
	"wowcollector.io/internal/entities/documents"
)

type MountLeaderboardRepository struct {
	collection *mongo.Collection
}

var instance *MountLeaderboardRepository

func GetRepository() *MountLeaderboardRepository {
	return instance
}

func Init(database *mongo.Database) {
	instance = &MountLeaderboardRepository{collection: database.Collection("mounts-leaderboard")}
	instance.createIndexes()
}

func (r *MountLeaderboardRepository) GetLeaderboardEntries() ([]*documents.LeaderboardDocument, error) {
	result, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		zap.L().Info("Error fetching mount leaderboard" + err.Error())
		return nil, err
	}
	defer result.Close(context.TODO())
	var mounts []*documents.LeaderboardDocument
	for result.Next(context.TODO()) {
		var mount *documents.LeaderboardDocument
		err := result.Decode(&mount)
		if err != nil {
			zap.L().Info("Error decoding mount leaderboard entry" + err.Error())
		}
		mounts = append(mounts, mount)
	}
	if err := result.Err(); err != nil {
		zap.L().Info("Error fetching mount leaderboard" + err.Error())
	}
	return mounts, nil
}

func (r *MountLeaderboardRepository) CreateLeaderboardEntry(document *documents.LeaderboardDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		zap.L().Info("Error inserting mount leaderboard document:" + err.Error())
		return err
	}
	return nil
}

func (r *MountLeaderboardRepository) UpdateLeaderboardEntry(document *documents.LeaderboardDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		zap.L().Info("Error updating mount leaderboard document:" + err.Error())
		return err
	}
	return nil
}

func (r *MountLeaderboardRepository) GetLeaderBoardEntry(character string, realm string, region blizzarddata.BattleNetRegion) (*documents.LeaderboardDocument, error) {
	filter := bson.D{
		{"character", character},
		{"realm", realm},
		{"region", region},
	}

	var leaderboardDocument documents.LeaderboardDocument
	err := r.collection.FindOne(context.TODO(), filter).Decode(&leaderboardDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Handle the case where no document was found
			return nil, fmt.Errorf("no leaderboard entry found for character: %s, realm: %s, region: %v", character, realm, region)
		}
		// Handle other potential errors
		return nil, err
	}

	return &leaderboardDocument, nil
}

func (r *MountLeaderboardRepository) createIndexes() {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "character", Value: 1},
			{Key: "realm", Value: 1},
			{Key: "region", Value: 1},
		},
	}

	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		zap.L().Error(err.Error())
	}
}
