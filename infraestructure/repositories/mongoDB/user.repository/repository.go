package user_repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"twittor-api/infraestructure/db/mongoDB"
)

type UserRepository struct {
	User *mongo.Collection
}

func New() *UserRepository {
	database := mongoDB.CurrentSession().DataBase()

	return &UserRepository{
		database.Collection("users"),
	}
}
