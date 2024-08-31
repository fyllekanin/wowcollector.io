package mountleaderboardrepository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
	"wowcollector.io/internal/entities/documents"
	commonrepository "wowcollector.io/internal/repository/repositories/common-repository"
)

type MountLeaderboardRepository struct {
	commonrepository.CommonRepository
	collection *mongo.Collection
}

var instance *MountLeaderboardRepository

func GetRepository() *MountLeaderboardRepository {
	return instance
}

func Init(database *mongo.Database) {
	collection := database.Collection("mounts-leaderboard")
	instance = &MountLeaderboardRepository{
		CommonRepository: commonrepository.CommonRepository{
			Collection: collection,
		},
		collection: collection,
	}
	instance.CreateCombinedIndex([]string{"character", "realm", "region"})
}

func (r *MountLeaderboardRepository) GetLeaderboardEntries() ([]*documents.LeaderboardDocument, error) {
	result, err := r.GetAll()
	if err != nil {
		zap.L().Info("Error fetching mount leaderboard entries" + err.Error())
		return nil, err
	}

	items := make([]*documents.LeaderboardDocument, len(result))
	for i, record := range result {
		var doc *documents.LeaderboardDocument
		bsonBytes, _ := bson.Marshal(record)
		bson.Unmarshal(bsonBytes, &doc)
		items[i] = doc
	}
	return items, nil
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
