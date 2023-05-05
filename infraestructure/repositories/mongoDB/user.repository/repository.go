package user_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"twittor-api/infraestructure/db/mongoDB"
)

type UserRepository struct {
	User    *mongo.Collection
	Context context.Context
	Cancel  func()
}

func New() *UserRepository {
	cnn := mongoDB.CurrentSession().DataBase()

	return &UserRepository{
		cnn.Database.Collection("users"),
		cnn.Ctx,
		cnn.Cancel,
	}
}
