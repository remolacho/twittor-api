package userRepository

import (
	"twittor-api/domain/models/user"
)

func (r *UserRepository) ExistsByEmail(email string) bool {
	if registerUser().Email == email {
		return true
	}

	return false
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	return registerUser(), nil
}
