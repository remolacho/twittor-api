package follow

import "errors"

func (e *Follow) UserRelationIsEmpty() (bool, error) {
	if len(e.FollowUserID) == 0 {
		return false, errors.New("the user id params is empty")
	}

	return true, nil
}

func (e *Follow) UsersEquals() (bool, error) {
	if e.UserID == e.FollowUserID {
		return false, errors.New("the user cannot follow yourself")
	}

	return true, nil
}
