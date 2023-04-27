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

func (us *UserService) Create(u *user.User) (*user.User, error) {
	if !u.EmailPresent() {
		return u, errors.New("error to create user, email is empty")
	}

	if !u.PasswordSizeValid(6) {
		return u, errors.New("the password must have a minimum of 6 characters")
	}

	repository := userRepository.NewUserRepository()

	if repository.ExistsByEmail(u.Email) {
		return u, errors.New("the user already exists")
	}

	return repository.Create(u)
}
