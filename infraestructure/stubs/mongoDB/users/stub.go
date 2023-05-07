package users

import (
	"twittor-api/domain/models/user"
)

type Stub struct{}

func New() *Stub {
	return &Stub{}
}

func (s *Stub) User(t string) *user.User {
	var _user *user.User

	switch t {
	case "email":
		_user = s.email()
		break
	case "created":
		_user = s.created()
		break
	case "password":
		_user = s.password()
		break
	case "new":
		_user = s.newUser()
		break
	}

	return _user
}
