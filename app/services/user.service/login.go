package user_service

import (
	"twittor-api/domain/models/user"
)

type UserLoginService struct {
	RepositoryUser user.IUser
}

func NewLogin(repository user.IUser) *UserLoginService {
	return &UserLoginService{
		repository,
	}
}

func (s *UserLoginService) Login(email, password string) (*user.User, bool) {
	u, err := s.RepositoryUser.FindByEmail(email)

	if err != nil {
		return nil, false
	}

	err = u.DecodePassword(password)

	if err != nil {
		return nil, false
	}

	return u, true
}
