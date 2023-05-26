package user_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"twittor-api/domain/models/user"
	repositoryFactoryFollow "twittor-api/infraestructure/repositories/factories/repository.factory.follower"
)

func (r *UserRepository) AllFollowed(ID string, page int64, searchTerm string) ([]user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	cursor, err := r.User.Find(ctx, r.searchQuery(searchTerm), r.optionsPage(page))

	if err != nil {
		return []user.User{}, err
	}

	return r.followed(cursor, ID)
}

func (r *UserRepository) followed(cursor *mongo.Cursor, ID string) ([]user.User, error) {
	var users []user.User

	repoFollowers := repositoryFactoryFollow.Build()

	for cursor.Next(context.TODO()) {
		var u user.User

		err := cursor.Decode(&u)

		if err != nil {
			return users, err
		}

		_, errFollower := repoFollowers.FindByUserID(ID, u.ID.Hex())

		if errFollower == nil {
			users = append(users, r.clearData(u))
		}
	}

	return users, nil
}
