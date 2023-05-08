package tweet_service

import (
	"errors"
	"twittor-api/domain/models/tweet"
)

type TweetDeleteService struct {
	RepositoryTweet tweet.ITweet
}

func NewDelete(repoTweeter tweet.ITweet) *TweetDeleteService {
	return &TweetDeleteService{
		repoTweeter,
	}
}

func (s *TweetDeleteService) Delete(tweetID string, userID string) (bool, error) {
	if len(tweetID) == 0 {
		return false, errors.New("the ID is empty")
	}

	_, err := s.RepositoryTweet.FindByUser(tweetID, userID)

	if err != nil {
		return false, err
	}

	return s.RepositoryTweet.Delete(tweetID)
}
