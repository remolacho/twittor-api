package user_service

import (
	"errors"
	"twittor-api/domain/models/user"
)

type UserUpdateService struct {
	RepositoryUser user.IUser
}

func NewUpdate(repository user.IUser) *UserUpdateService {
	return &UserUpdateService{
		repository,
	}
}

func (s *UserUpdateService) Update(ID string, u *user.User) (*user.User, error) {
	currentUser, err := s.RepositoryUser.Find(ID)

	if err != nil {
		return nil, errors.New("user not found")
	}

	u.ID = currentUser.ID
	u.Password = currentUser.Password
	u.Email = currentUser.Email
	u.Avatar = currentUser.Avatar
	u.Banner = currentUser.Banner

	return s.RepositoryUser.Update(u)
}
