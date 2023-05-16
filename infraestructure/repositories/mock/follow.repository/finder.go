package follow_repository

import (
	"twittor-api/domain/models/follow"
	factoryStubFollowers "twittor-api/infraestructure/stubs/factories/factory.followers"
)

func (r *FollowRepository) FindByObject(t *follow.Follow) bool {
	factory := factoryStubFollowers.Build()
	relationObj := factory.Follow("createFollow")

	if (relationObj.UserID == t.UserID) && (relationObj.FollowUserID == t.FollowUserID) {
		return true
	}

	return false
}
