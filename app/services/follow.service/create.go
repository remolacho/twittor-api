package follow_service

import (
	"errors"
	"twittor-api/domain/models/follow"
	"twittor-api/domain/models/user"
)

type RelationCreateService struct {
	RepositoryRelation follow.IFollow
	RepositoryUser     user.IUser
}

func NewCreate(repoUser user.IUser, repoRelation follow.IFollow) *RelationCreateService {
	return &RelationCreateService{
		repoRelation,
		repoUser,
	}
}

func (s *RelationCreateService) Create(userID string, followUserID string) (bool, error) {
	t := follow.New()
	t.UserID = userID
	t.FollowUserID = followUserID

	if flag, err := t.UserRelationIsEmpty(); flag != true {
		return flag, err
	}

	if flag, err := t.UsersEquals(); flag != true {
		return flag, err
	}

	if _, err := s.RepositoryUser.Find(t.FollowUserID); err != nil {
		return false, err
	}

	if s.RepositoryRelation.FindByObject(t) {
		return false, errors.New("the follow is duplicated")
	}

	_, flag, err := s.RepositoryRelation.Create(t)

	return flag, err
}
