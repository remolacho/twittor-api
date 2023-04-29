package userRepository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/user"
)

func (r *UserRepository) Create(u *user.User) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var err error
	u.Password, err = u.EncryptPassword()
	u.ID = primitive.NewObjectID()

	if err != nil {
		return u, errors.New("the error to encrypt password " + err.Error())
	}

	_, errRecord := r.User.InsertOne(ctx, u)

	return u, errRecord
}
