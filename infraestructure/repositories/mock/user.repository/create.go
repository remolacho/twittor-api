package user_repository

import (
	"twittor-api/domain/models/user"
)

func (r *UserRepository) Create(u *user.User) (*user.User, error) {
	return u, nil
}
