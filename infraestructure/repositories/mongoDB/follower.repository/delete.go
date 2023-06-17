package follower_repository

import (
	"context"
	"time"
)

func (r *FollowRepository) DestroyAllowed(userID string, followUserID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	f, err := r.FindByUserID(userID, followUserID)

	if err != nil {
		return false, err
	}

	_, err = r.Follower.DeleteOne(ctx, f)

	if err != nil {
		return false, err
	}

	return true, nil
}
