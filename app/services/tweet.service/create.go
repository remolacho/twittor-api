package tweet_service

import (
	"errors"
	"twittor-api/domain/models/tweet"
)

type TweetCreateService struct {
	RepositoryTweet tweet.ITweet
}

func NewTweet(repository tweet.ITweet) *TweetCreateService {
	return &TweetCreateService{
		repository,
	}
}

func (s *TweetCreateService) Create(t *tweet.Tweet) (*tweet.Tweet, error) {
	if !t.MessagePresent() {
		return t, errors.New("the tweet should has message")
	}

	if !t.UserIdPresent() {
		return t, errors.New("the tweet should has userID")
	}

	return s.RepositoryTweet.Create(t)
}
