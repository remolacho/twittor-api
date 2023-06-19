package user_repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"twittor-api/domain/models/user"
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

// private

func (r *UserRepository) optionsPage(page int64) *options.FindOptions {
	var limit int64 = 20
	skipUsers := (page - 1) * limit

	opts := options.Find()
	opts.SetSkip(skipUsers)
	opts.SetLimit(limit)

	return opts
}

func (r *UserRepository) searchQuery(searchTerm string) bson.M {
	return bson.M{
		"name": bson.M{
			"$regex": `(?i)` + searchTerm,
		},
	}
}

func (r *UserRepository) clearData(u user.User) user.User {
	u.Password = ""
	u.Banner = ""
	u.SideWeb = ""
	u.Email = ""
	u.Location = ""

	return u
}
