package user_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/user"
)

func (r *UserRepository) Create(u *user.User) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var err error

	u.Password, err = u.EncryptPassword()

	defer cancel()

	if err != nil {
		return u, errors.New("the error to encrypt password " + err.Error())
	}

	u.ID = primitive.NewObjectID()

	_, err = r.User.InsertOne(ctx, u)

	return u, err
}
