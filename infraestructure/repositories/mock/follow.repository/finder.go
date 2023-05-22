package follow_repository

import (
	"errors"
	"twittor-api/domain/models/follow"
	factoryStubFollowers "twittor-api/infraestructure/stubs/factories/factory.followers"
)

func (r *FollowRepository) FindByObject(t *follow.Follow) bool {
	factory := factoryStubFollowers.Build()
	relationObj := factory.Follow("created")

	if (relationObj.UserID == t.UserID) && (relationObj.FollowUserID == t.FollowUserID) {
		return true
	}

	return false
}

func (r *FollowRepository) FindAllowed(ID string, userID string) (*follow.Follow, error) {
	factory := factoryStubFollowers.Build()
	relationObj := factory.Follow("created")

	if userID != relationObj.UserID {
		return nil, errors.New("the user not belongs to the relation")
	}

	return relationObj, nil
}

func (r *FollowRepository) FindByUserID(userID string, followerID string) (*follow.Follow, error) {
	factory := factoryStubFollowers.Build()
	relationObj := factory.Follow("created")

	if userID != relationObj.UserID || followerID != relationObj.FollowUserID {
		return nil, errors.New("the user not belongs to the relation")
	}

	return relationObj, nil
}
