package repository

import (
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
