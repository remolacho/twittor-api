package follower_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"twittor-api/domain/models/follower"
)

func (r *FollowRepository) IncludeTweet(userID string, page int64) ([]follower.HasOneTweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var followers []follower.HasOneTweet

	defer cancel()

	cursor, err := r.Follower.Aggregate(ctx, r.buildFilters(userID, page))

	if err != nil {
		return followers, err
	}

	err = cursor.All(ctx, &followers)

	if err != nil {
		return followers, err
	}

	return followers, nil
}

func (r *FollowRepository) buildFilters(userID string, page int64) []bson.M {
	var limit int64 = 20
	skipTweets := (page - 1) * limit

	filters := make([]bson.M, 0)
	filters = append(filters, bson.M{"$match": bson.M{"userid": userID}})
	filters = append(filters, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "followUserId",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})

	filters = append(filters, bson.M{"$unwind": "$tweet"})
	filters = append(filters, bson.M{"$sort": bson.M{"tweets.createdAt": -1}})
	filters = append(filters, bson.M{"$skip": skipTweets})
	filters = append(filters, bson.M{"$limit": limit})

	return filters
}
