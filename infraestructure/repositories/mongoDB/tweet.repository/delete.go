package tweet_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (r *TweetRepository) Delete(ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	filter := bson.M{"_id": ID}

	_, err := r.Tweet.DeleteOne(ctx, filter)

	if err != nil {
		return false, err
	}

	return true, nil
}
