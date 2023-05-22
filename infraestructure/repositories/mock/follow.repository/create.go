package follow_repository

import "twittor-api/domain/models/follow"

func (r *FollowRepository) Create(t *follow.Follow) (*follow.Follow, bool, error) {
	return t, true, nil
}
