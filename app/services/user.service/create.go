package user_service

import (
	"errors"
	"twittor-api/domain/models/user"
)

type UserCreateService struct {
	RepositoryUser user.IUser
}

func NewUser(repository user.IUser) *UserCreateService {
	return &UserCreateService{
		repository,
	}
}

func (us *UserCreateService) Create(u *user.User) (*user.User, error) {
	if !u.EmailPresent() {
		return u, errors.New("error to create user, email is empty")
	}

	if !u.PasswordSizeValid(6) {
		return u, errors.New("the password must have a minimum of 6 characters")
	}

	if us.RepositoryUser.ExistsByEmail(u.Email) {
		return u, errors.New("the user already exists")
	}

	return us.RepositoryUser.Create(u)
}
