package follow_service

import (
	"errors"
	"twittor-api/domain/models/follow"
)

type FollowDeleteService struct {
	RepositoryRelation follow.IFollow
}

func NewDelete(repoFollow follow.IFollow) *FollowDeleteService {
	return &FollowDeleteService{
		repoFollow,
	}
}

func (s *FollowDeleteService) Destroy(followID string, userID string) (bool, error) {
	if len(followID) == 0 {
		return false, errors.New("the follow was not deleted")
	}

	return s.RepositoryRelation.DestroyAllowed(followID, userID)
}
