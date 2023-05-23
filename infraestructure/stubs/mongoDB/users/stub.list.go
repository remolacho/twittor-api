package users

import (
	"twittor-api/domain/models/user"
)

func (s *Stub) list() []user.User {
	var users []user.User
	users = append(users, *s.User("created"))
	return users
}
