package follower

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"twittor-api/domain/models/tweet"
)

type HasOneTweet struct {
	ID           primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	UserID       string             `bson:"userid" json:"userid,omitempty"`
	FollowUserID string             `bson:"followUserId" json:"followUserId,omitempty"`
	Tweet        tweet.Tweet        `bson:"tweet" json:"tweet,omitempty"`
}

func NewHaOneTweet() *HasOneTweet {
	return &HasOneTweet{}
}
