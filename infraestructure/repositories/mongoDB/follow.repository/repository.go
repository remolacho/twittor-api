package follow_repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"twittor-api/infraestructure/db/mongoDB"
)

type FollowRepository struct {
	Follow *mongo.Collection
}

func New() *FollowRepository {
	database := mongoDB.CurrentSession().DataBase()

	return &FollowRepository{
		database.Collection("followers"),
	}
}
