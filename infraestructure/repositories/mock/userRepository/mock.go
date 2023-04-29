package userRepository

import "twittor-api/domain/models/user"

func GetUser() *user.User {
	return &user.User{
		Email:    "jonathangrh.25@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}
}
