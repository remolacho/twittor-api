package follower_repository

import (
	"errors"
	"twittor-api/domain/models/follower"
	factoryStubFollowers "twittor-api/infraestructure/stubs/factories/factory.followers"
)

func (r *FollowRepository) FindByObject(t *follower.Follower) bool {
	factory := factoryStubFollowers.Build()
	relationObj := factory.Follower("created")

	if (relationObj.UserID == t.UserID) && (relationObj.FollowUserID == t.FollowUserID) {
		return true
	}

	return false
}

func (r *FollowRepository) FindAllowed(ID string, userID string) (*follower.Follower, error) {
	factory := factoryStubFollowers.Build()
	relationObj := factory.Follower("created")

	if userID != relationObj.UserID {
		return nil, errors.New("the user not belongs to the relation")
	}

	return relationObj, nil
}

func (r *FollowRepository) FindByUserID(userID string, followerID string) (*follower.Follower, error) {
	factory := factoryStubFollowers.Build()
	relationObj := factory.Follower("created")

	if userID != relationObj.UserID || followerID != relationObj.FollowUserID {
		return nil, errors.New("the user not belongs to the relation")
	}

	return relationObj, nil
}
