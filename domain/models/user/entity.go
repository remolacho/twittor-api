package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name      string             `bson:"name, omitempty" json:"name,omitempty"`
	LastName  string             `bson:"lastname, omitempty" json:"lastname,omitempty"`
	Birthday  time.Time          `bson:"birthday, omitempty" json:"birthday,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	SideWeb   string             `bson:"side_web" json:"sideWeb,omitempty"`
}

func New() *User {
	return &User{}
}
