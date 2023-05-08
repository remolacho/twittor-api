package tweet_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (r *TweetRepository) Delete(ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	objectID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objectID}

	_, err := r.Tweet.DeleteOne(ctx, filter)

	if err != nil {
		return false, err
	}

	return true, nil
}
