package tweet

import (
	"time"
)

type Tweet struct {
	ID        string    `bson:"_id, omitempty" json:"id"`
	UserID    string    `bson:"userid" json:"userid,omitempty"`
	Message   string    `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt,omitempty"`
}

func New(userID string) *Tweet {
	return &Tweet{
		UserID: userID,
	}
}
