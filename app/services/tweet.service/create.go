package tweet_service

import (
	"errors"
	"twittor-api/domain/models/tweet"
	"twittor-api/domain/models/user"
)

type TweetCreateService struct {
	RepositoryTweet tweet.ITweet
	RepositoryUser  user.IUser
}

func NewTweet(repoTweeter tweet.ITweet, repoUser user.IUser) *TweetCreateService {
	return &TweetCreateService{
		repoTweeter,
		repoUser,
	}
}

func (s *TweetCreateService) Create(t *tweet.Tweet) (*tweet.Tweet, error) {
	if !t.MessagePresent() {
		return t, errors.New("the tweet should has message")
	}

	if !t.UserIdPresent() {
		return t, errors.New("the tweet should has userID")
	}

	_, err := s.RepositoryUser.Find(t.UserID)

	if err != nil {
		return t, err
	}

	return s.RepositoryTweet.Create(t)
}
