package tweet_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/tweet"
)

func (r *TweetRepository) Create(t *tweet.Tweet) (*tweet.Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var err error

	defer cancel()

	t.ID = primitive.NewObjectID().Hex()
	t.CreatedAt = time.Now()

	_, err = r.Tweet.InsertOne(ctx, t)

	return t, err
}
