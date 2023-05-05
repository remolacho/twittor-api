package user_service

import (
	"twittor-api/domain/models/user"
)

type ProfileService struct {
	RepositoryUser user.IUser
}

func NewProfile(repository user.IUser) *ProfileService {
	return &ProfileService{
		repository,
	}
}

func (s *ProfileService) Get(ID string) (*user.User, error) {
	u, err := s.RepositoryUser.Find(ID)
	u.Password = ""
	return u, err
}
