package tweet_repository

import (
	"errors"
	"twittor-api/domain/models/tweet"
	stubTweets "twittor-api/infraestructure/stubs/tweets"
)

func (r *TweetRepository) AllByPagedUser(userID string, page int64) ([]tweet.Tweet, error) {
	var tweets []tweet.Tweet

	for _, t := range stubTweets.StubTweets("list") {
		if t.UserID != userID {
			return tweets, errors.New("user not found in tweet")
		}

		tweets = append(tweets, t)
	}

	return tweets, nil
}
