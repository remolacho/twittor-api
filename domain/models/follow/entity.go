package follow

import "go.mongodb.org/mongo-driver/bson/primitive"

type Follow struct {
	ID           primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	UserID       string             `bson:"userid" json:"userid,omitempty"`
	FollowUserID string             `bson:"followUserId" json:"followUserId,omitempty"`
}

func New() *Follow {
	return &Follow{}
}
