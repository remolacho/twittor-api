package userRepository

import "twittor-api/domain/models/user"

func NewUser() *user.User {
	return &user.User{
		Email:    "jodoe@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}
}

func RegisterUser() *user.User {
	return &user.User{
		Email:    "jonathangrh.25@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}
}

func GetUserWithOutEmail() *user.User {
	return &user.User{
		Email:    "",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "123456",
	}
}

func GetUserSizeLessPassword() *user.User {
	return &user.User{
		Email:    "jonathangrh.25@gmail.com",
		Name:     "Jonathan",
		LastName: "Rojas",
		Password: "12345",
	}
}
