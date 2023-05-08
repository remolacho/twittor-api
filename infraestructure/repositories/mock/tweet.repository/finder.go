package tweet_repository

import (
	"errors"
	"twittor-api/domain/models/tweet"
	stubFactoryTweet "twittor-api/infraestructure/stubs/factories/factory.tweets"
)

func (r *TweetRepository) FindByUser(ID string, userID string) (*tweet.Tweet, error) {
	stub := stubFactoryTweet.Build()
	t := stub.Tweet("createTweet")

	if t.ID.Hex() != ID || t.UserID != userID {
		return nil, errors.New("the tweet not belong to user")
	}

	return t, nil
}
