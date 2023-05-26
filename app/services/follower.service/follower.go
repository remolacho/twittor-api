package follower_service

import (
	"errors"
	"twittor-api/domain/models/follower"
)

type FindFollowerService struct {
	RepositoryRelation follower.IFollow
}

type ResponseFollower struct {
	Followed bool
}

func NewFollower(repository follower.IFollow) *FindFollowerService {
	return &FindFollowerService{
		repository,
	}
}

func (s *FindFollowerService) Followed(userID string, followerID string) (ResponseFollower, error) {
	response := ResponseFollower{
		Followed: false,
	}

	if len(userID) == 0 || len(followerID) == 0 {
		return response, errors.New("the userID or followerID is empties")
	}

	_, err := s.RepositoryRelation.FindByUserID(userID, followerID)

	if err != nil {
		return response, err
	}

	response.Followed = true
	return response, nil
}
