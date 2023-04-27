package userRepository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/db"
)

type Connection struct {
	Collection *mongo.Collection
}

func NewUserRepository() *Connection {
	database := db.CurrentSession().Client.Database(os.Getenv("DATABASE_NAME"))

	return &Connection{
		database.Collection("users"),
	}
}

func (cnn *Connection) Create(u *user.User) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var err error
	u.Password, err = u.EncryptPassword()

	if err != nil {
		return u, errors.New("the error to encrypt password " + err.Error())
	}

	record := cnn.Collection
	_, errRecord := record.InsertOne(ctx, u)

	return u, errRecord
}

func (cnn *Connection) ExistsByEmail(email string) bool {
	return true
}
