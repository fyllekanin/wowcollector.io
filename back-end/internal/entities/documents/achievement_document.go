package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type AchievementDocument struct {
	ObjectID      primitive.ObjectID `bson:"_id" json:"_id"`
	Id            int                `bson:"id" json:"id"`
	Name          string             `bson:"name" json:"name"`
	Description   string             `bson:"description" json:"description"`
	Points        int                `bson:"points" json:"points"`
	IsAccountWide bool               `bson:"isAccountWide" json:"isAccountWide"`
	Icon          string             `bson:"icon" json:"icon"`
	DisplayOrder  int                `bson:"displayOrder" json:"displayOrder"`
	CategoryId    int                `bson:"categoryId" json:"categoryId"`
}

func (r *AchievementDocument) IsEqual(other *AchievementDocument) bool {
	return r.Id == other.Id &&
		r.Name == other.Name &&
		r.Description == other.Description &&
		r.Points == other.Points &&
		r.IsAccountWide == other.IsAccountWide &&
		r.Icon == other.Icon &&
		r.DisplayOrder == other.DisplayOrder &&
		r.CategoryId == other.CategoryId
}
