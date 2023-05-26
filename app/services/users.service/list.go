package users_service

import "twittor-api/domain/models/user"

type FollowedService struct {
	Repository user.IUser
}

func New(repo user.IUser) *FollowedService {
	return &FollowedService{
		repo,
	}
}

func (s *FollowedService) Followed(ID string, page int64, searchTerm string) ([]user.User, error) {
	return s.Repository.AllFollowed(ID, page, searchTerm)
}

func (s *FollowedService) UnFollowed(ID string, page int64, searchTerm string) ([]user.User, error) {
	return s.Repository.AllUnFollowed(ID, page, searchTerm)
}
