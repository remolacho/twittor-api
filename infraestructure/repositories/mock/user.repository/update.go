package user_repository

import (
	"twittor-api/domain/models/user"
)

func (r *UserRepository) Update(u *user.User) (*user.User, error) {
	return u, nil
}
