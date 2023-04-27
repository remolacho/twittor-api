package userRepository

import (
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/db"
)

type Connection struct {
	Data *db.Mongo
}

func NewUserRepository() *Connection {
	return &Connection{
		db.CurrentSession(),
	}
}

func (cnn *Connection) Create(u *user.User) (error, *user.User) {
	return nil, u
}

func (cnn *Connection) ExistsByEmail(email string) bool {
	return true
}
