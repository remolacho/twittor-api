package user_service

import (
	"errors"
	"twittor-api/domain/models/user"
)

type UserUpdateService struct {
	RepositoryUser user.IUser
}

func NewUpdateUser(repository user.IUser) *UserUpdateService {
	return &UserUpdateService{
		repository,
	}
}

func (s *UserUpdateService) Update(ID string, u *user.User) (*user.User, error) {
	currentUser, err := s.RepositoryUser.Find(ID)

	if err != nil {
		return nil, errors.New("user not found")
	}

	if !u.EmailPresent() {
		return u, errors.New("error to create user, email is empty")
	}

	u.ID = currentUser.ID
	u.Password = currentUser.Password

	return s.RepositoryUser.Update(u)
}
