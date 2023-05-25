package tweet_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"twittor-api/domain/models/tweet"
)

func (r *TweetRepository) FindByUser(ID string, userID string) (*tweet.Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var t tweet.Tweet

	filter := bson.M{"_id": ID, "userid": userID}
	err := r.Tweet.FindOne(ctx, filter).Decode(&t)

	defer cancel()

	if err != nil {
		return nil, errors.New("the tweet not belong to user " + err.Error())
	}

	return &t, nil
}
