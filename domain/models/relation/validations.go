package relation

import "errors"

func (e *Relation) UserRelationIsEmpty() (bool, error) {
	if len(e.UserRelationID) == 0 {
		return false, errors.New("the user id params is empty")
	}

	return true, nil
}

func (e *Relation) UsersEquals() (bool, error) {
	if e.UserID == e.UserRelationID {
		return false, errors.New("the user cannot follow yourself")
	}

	return true, nil
}
