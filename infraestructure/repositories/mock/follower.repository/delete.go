package follower_repository

func (r *FollowRepository) DestroyAllowed(ID string, userID string) (bool, error) {
	_, err := r.FindAllowed(ID, userID)

	if err != nil {
		return false, err
	}

	return true, nil
}
