package user_repository

import (
	"errors"
	"twittor-api/domain/models/user"
)

func (r *UserRepository) ExistsByEmail(email string) bool {
	if StubUser("created").Email == email {
		return true
	}

	return false
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	if !r.ExistsByEmail(email) {
		return nil, errors.New("user not found")
	}

	return StubUser("created"), nil
}
