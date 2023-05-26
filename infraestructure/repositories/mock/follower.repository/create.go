package follower_repository

import "twittor-api/domain/models/follower"

func (r *FollowRepository) Create(t *follower.Follower) (*follower.Follower, bool, error) {
	return t, true, nil
}
