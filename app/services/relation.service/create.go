package relation_service

import (
	"errors"
	"twittor-api/domain/models/relation"
	"twittor-api/domain/models/user"
)

type RelationCreateService struct {
	RepositoryRelation relation.IRelation
	RepositoryUser     user.IUser
}

func NewCreate(repoUser user.IUser, repoRelation relation.IRelation) *RelationCreateService {
	return &RelationCreateService{
		repoRelation,
		repoUser,
	}
}

func (s *RelationCreateService) Create(userID string, userRelationID string) (bool, error) {
	t := relation.New()
	t.UserID = userID
	t.UserRelationID = userRelationID

	if flag, err := t.UserRelationIsEmpty(); flag != true {
		return flag, err
	}

	if flag, err := t.UsersEquals(); flag != true {
		return flag, err
	}

	if _, err := s.RepositoryUser.Find(t.UserRelationID); err != nil {
		return false, err
	}

	if s.RepositoryRelation.FindByObject(t) {
		return false, errors.New("the relation is duplicated")
	}

	_, flag, err := s.RepositoryRelation.Create(t)

	return flag, err
}
