package tweet_repository

import "twittor-api/domain/models/tweet"

func (r *TweetRepository) AllByPagedUser(userID string, page int64) ([]tweet.Tweet, error) {
	return []tweet.Tweet{}, nil
}
