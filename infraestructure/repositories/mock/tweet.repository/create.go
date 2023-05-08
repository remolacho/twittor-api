package tweet_repository

import "twittor-api/domain/models/tweet"

func (r *TweetRepository) Create(t *tweet.Tweet) (*tweet.Tweet, error) {
	return t, nil
}
