package repository

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	achievementcategoryrepository "wowcollector.io/internal/repository/repositories/achievement-category-repository"
	achievementrepository "wowcollector.io/internal/repository/repositories/achievement-repository"
	mountleaderboardrepository "wowcollector.io/internal/repository/repositories/mount-leaderboard-repository"
	mountrepository "wowcollector.io/internal/repository/repositories/mount-repository"
	mountviewrepository "wowcollector.io/internal/repository/repositories/mount-view-repository"
	petrepository "wowcollector.io/internal/repository/repositories/pet-repository"
	petviewrepository "wowcollector.io/internal/repository/repositories/pet-view-repository"
	realmrepository "wowcollector.io/internal/repository/repositories/realm-repository"
	seedrepository "wowcollector.io/internal/repository/repositories/seed-repository"
	toyrepository "wowcollector.io/internal/repository/repositories/toy-repository"
	toyviewrepository "wowcollector.io/internal/repository/repositories/toy-view-repository"
)

type RepositoryFactory struct {
	database *mongo.Database
}

func (r *RepositoryFactory) init() {
	realmrepository.Init(r.database)
	mountrepository.Init(r.database)
	seedrepository.Init(r.database)
	mountviewrepository.Init(r.database)
	achievementrepository.Init(r.database)
	achievementcategoryrepository.Init(r.database)
	toyrepository.Init(r.database)
	mountleaderboardrepository.Init(r.database)
	toyviewrepository.Init(r.database)
	petrepository.Init(r.database)
	petviewrepository.Init(r.database)
}

var instance *RepositoryFactory

func InitRepositories(database *mongo.Database) {
	instance = &RepositoryFactory{
		database: database,
	}
	instance.init()
}

func GetRepositoryFactory() *RepositoryFactory {
	return instance
}

func GetDatabaseClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(getDatabaseUri())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		zap.L().Info(err.Error())
	}
	zap.L().Info("Connected to MongoDB")
	return client
}

func getDatabaseUri() string {
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	return fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
}
