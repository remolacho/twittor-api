package user_repository

import (
	"errors"
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/stubs/users"
)

func (r *UserRepository) ExistsByEmail(email string) bool {
	if users.StubUser("created").Email == email {
		return true
	}

	return false
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	if !r.ExistsByEmail(email) {
		return nil, errors.New("user not found")
	}

	return users.StubUser("created"), nil
}

func (r *UserRepository) Find(ID string) (*user.User, error) {
	u := users.StubUser("created")

	if u.ID.Hex() != ID {
		return nil, errors.New("user not found")
	}

	return u, nil
}
