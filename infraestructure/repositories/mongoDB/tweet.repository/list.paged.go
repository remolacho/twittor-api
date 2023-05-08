package tweet_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"twittor-api/domain/models/tweet"
)

func (r *TweetRepository) AllByPagedUser(userID string, page int64) ([]tweet.Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	filter := bson.M{"userid": userID}

	cursor, err := r.Tweet.Find(ctx, filter, r.optionsPage(page))

	if err != nil {
		return []tweet.Tweet{}, err
	}

	return r.tweetsList(cursor)
}

func (r *TweetRepository) optionsPage(page int64) *options.FindOptions {
	var limit int64 = 20
	skipTweets := (page - 1) * limit

	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSort(bson.D{{Key: "createdAt", Value: -1}})
	opts.SetSkip(skipTweets)
	return opts
}

func (r *TweetRepository) tweetsList(cursor *mongo.Cursor) ([]tweet.Tweet, error) {
	var tweets []tweet.Tweet

	for cursor.Next(context.TODO()) {
		var _tweet tweet.Tweet
		err := cursor.Decode(&_tweet)

		if err != nil {
			return tweets, err
		}

		tweets = append(tweets, _tweet)
	}

	return tweets, nil
}
