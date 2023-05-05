package user_repository

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"twittor-api/domain/models/user"
)

func (r *UserRepository) Create(u *user.User) (*user.User, error) {
	var err error

	u.Password, err = u.EncryptPassword()

	if err != nil {
		return u, errors.New("the error to encrypt password " + err.Error())
	}

	u.ID = primitive.NewObjectID()

	_, err = r.User.InsertOne(r.Context, u)

	defer r.Cancel()

	return u, err
}
