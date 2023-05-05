package tweet

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tweet struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	UserID    string             `bson:"userid" json:"userid,omitempty"`
	Message   string             `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt,omitempty"`
}
