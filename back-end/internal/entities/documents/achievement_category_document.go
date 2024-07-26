package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type AchievementCategoryDocument struct {
	ObjectID       primitive.ObjectID `bson:"_id" json:"_id"`
	Id             int                `bson:"id" json:"id"`
	Name           string             `bson:"name" json:"name"`
	IsRootCategory bool               `bson:"isRootCategory" json:"isRootCategory"`
	RootCategoryId int                `bson:"rootCategoryId" json:"rootCategoryId"`
	DisplayOrder   int                `bson:"displayOrder" json:"displayOrder"`
}

func (r *AchievementCategoryDocument) IsEqual(other *AchievementCategoryDocument) bool {
	return r.Id == other.Id &&
		r.Name == other.Name &&
		r.IsRootCategory == other.IsRootCategory &&
		r.RootCategoryId == other.RootCategoryId &&
		r.RootCategoryId == other.RootCategoryId &&
		r.DisplayOrder == other.DisplayOrder
}
