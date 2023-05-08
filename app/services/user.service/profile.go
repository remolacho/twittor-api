package user_service

import (
	"twittor-api/domain/models/user"
)

type UserProfileService struct {
	RepositoryUser user.IUser
}

func NewProfile(repository user.IUser) *UserProfileService {
	return &UserProfileService{
		repository,
	}
}

func (s *UserProfileService) Get(ID string) (*user.User, error) {
	u, err := s.RepositoryUser.Find(ID)
	u.Password = ""
	return u, err
}
