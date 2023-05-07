package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"twittor-api/domain/models/user"
)

func StubUser(t string) *user.User {
	var _user *user.User

	switch t {
	case "email":
		_user = getUserWithOutEmail()
		break
	case "created":
		_user = registerUser()
		break
	case "password":
		_user = getUserSizeLessPassword()
		break
	case "new":
		_user = newUser()
		break
	}

	return _user
}

func GetID() primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex("0000000000000000000000")
	return objectID
}

func newUser() *user.User {
	return &user.User{
		Email:    "jodoe@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}
}

func registerUser() *user.User {
	u := &user.User{
		ID:       GetID(),
		Email:    "jonathangrh.25@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}

	password, _ := u.EncryptPassword()
	u.Password = password
	return u
}

func getUserWithOutEmail() *user.User {
	return &user.User{
		ID:       GetID(),
		Email:    "",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}
}

func getUserSizeLessPassword() *user.User {
	return &user.User{
		ID:       GetID(),
		Email:    "jonathangrh.25@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "12345",
	}
}
