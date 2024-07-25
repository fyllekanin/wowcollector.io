package mountviewrepository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"wowcollector.io/entities/documents"
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
		fmt.Println("Error fetching mount view", err)
		return nil, err
	}
	defer result.Close(context.TODO())
	var mounts []*documents.MountViewDocument
	for result.Next(context.TODO()) {
		var mount *documents.MountViewDocument
		err := result.Decode(&mount)
		if err != nil {
			fmt.Println("Error decoding mount view", err)
		}
		mounts = append(mounts, mount)
	}
	if err := result.Err(); err != nil {
		fmt.Println("Error fetching mounts", err)
	}
	return mounts, nil
}

func (r *MountViewRepository) CreateMountView(document *documents.MountViewDocument) error {
	_, err := r.collection.InsertOne(context.TODO(), document)
	if err != nil {
		fmt.Println("Error inserting mount view document:", err)
		return err
	}
	return nil
}

func (r *MountViewRepository) GetDefaultMountView() (*documents.MountViewDocument, error) {
	filter := bson.D{{"isDefault", true}}
	var mount *documents.MountViewDocument

	err := r.collection.FindOne(context.TODO(), filter).Decode(&mount)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No default mount view found")
			return nil, nil
		}
		fmt.Println("Error fetching default mount view", err)
		return nil, err
	}

	return mount, nil
}

func (r *MountViewRepository) UpdateMountView(document *documents.MountViewDocument) error {
	filter := bson.D{{"_id", document.ObjectID}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", document},
	})
	if err != nil {
		fmt.Println("Error updating mount view document:", err)
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
		log.Fatal(err)
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
		log.Fatal(err)
	}
}
