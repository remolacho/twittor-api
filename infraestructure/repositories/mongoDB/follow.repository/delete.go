package follow_repository

import (
	"context"
	"time"
)

func (r *FollowRepository) DestroyAllowed(ID string, userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	f, err := r.FindAllowed(ID, userID)

	if err != nil {
		return false, err
	}

	_, err = r.Follow.DeleteOne(ctx, f)

	if err != nil {
		return false, err
	}

	return true, nil
}
