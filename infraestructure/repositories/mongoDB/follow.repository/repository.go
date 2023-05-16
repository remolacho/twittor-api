package follow_repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"twittor-api/infraestructure/db/mongoDB"
)

type RelationRepository struct {
	Relation *mongo.Collection
}

func New() *RelationRepository {
	database := mongoDB.CurrentSession().DataBase()

	return &RelationRepository{
		database.Collection("followers"),
	}
}
