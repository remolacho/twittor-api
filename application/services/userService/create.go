package userService

import (
	"errors"
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/repositories/userRepository"
)

type UserService struct{}

func NewUser() *UserService {
	return &UserService{}
}

func (us *UserService) Create(u *user.User) (error, *user.User) {
	if !u.EmailPresent() {
		return errors.New("error to create user, email is empty"), u
	}

	if !u.PasswordSizeValid(6) {
		return errors.New("the password must have a minimum of 6 characters"), u
	}

	repository := userRepository.NewUserRepository()

	if repository.ExistsByEmail(u.Email) {
		return errors.New("the user already exists"), u
	}

	return repository.Create(u)
}
