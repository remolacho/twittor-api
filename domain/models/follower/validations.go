package follower

import "errors"

func (e *Follower) UserRelationIsEmpty() (bool, error) {
	if len(e.FollowUserID) == 0 {
		return false, errors.New("the user id params is empty")
	}

	return true, nil
}

func (e *Follower) UsersEquals() (bool, error) {
	if e.UserID == e.FollowUserID {
		return false, errors.New("the user cannot follow yourself")
	}

	return true, nil
}
