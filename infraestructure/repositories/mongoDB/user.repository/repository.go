package user_repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"twittor-api/infraestructure/db/mongoDB"
)

type UserRepository struct {
	User *mongo.Collection
}

func New() *UserRepository {
	database := mongoDB.CurrentSession().Client.Database(os.Getenv("DATABASE_NAME"))

	return &UserRepository{
		database.Collection("users"),
	}
}
