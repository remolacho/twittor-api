package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"twittor-api/domain/models/user"
)

func (s *Stub) newUser() *user.User {
	return &user.User{
		Email:    "jodoe@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}
}

func (s *Stub) created() *user.User {
	u := &user.User{
		ID:       s.getID(),
		Email:    "jonathangrh.25@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}

	password, _ := u.EncryptPassword()
	u.Password = password
	return u
}

func (s *Stub) other() *user.User {
	u := &user.User{
		ID:       primitive.NewObjectID(),
		Email:    "jonathangrh.25@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}

	password, _ := u.EncryptPassword()
	u.Password = password
	return u
}

func (s *Stub) email() *user.User {
	return &user.User{
		ID:       s.getID(),
		Email:    "",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}
}

func (s *Stub) password() *user.User {
	return &user.User{
		ID:       s.getID(),
		Email:    "jonathangrh.25@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "12345",
	}
}

func (s *Stub) getID() primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex("0000000000000000000000")
	return objectID
}
