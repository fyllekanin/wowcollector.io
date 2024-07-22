package repository

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	realmrepository "wowcollector.io/repository/repositories"
)

type RepositoryFactory struct {
	database *mongo.Database
}

func (r *RepositoryFactory) init() {
	realmrepository.Init(r.database)
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

func GetDatabaseUri() string {
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	return fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
}
