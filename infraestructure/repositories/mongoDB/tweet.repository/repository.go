package tweet_repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"twittor-api/infraestructure/db/mongoDB"
)

type TweetRepository struct {
	Tweet *mongo.Collection
}

func New() *TweetRepository {
	database := mongoDB.CurrentSession().DataBase()

	return &TweetRepository{
		database.Collection("tweets"),
	}
}
