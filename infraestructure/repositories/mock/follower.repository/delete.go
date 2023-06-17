package follower_repository

func (r *FollowRepository) DestroyAllowed(userID string, followUserID string) (bool, error) {
	_, err := r.FindByUserID(userID, followUserID)

	if err != nil {
		return false, err
	}

	return true, nil
}
