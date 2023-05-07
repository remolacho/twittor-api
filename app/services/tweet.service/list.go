package tweet_service

import (
	"errors"
	"twittor-api/domain/models/tweet"
)

type TweetsListService struct {
	RepositoryTweet tweet.ITweet
}

func NewList(repoTweet tweet.ITweet) *TweetsListService {
	return &TweetsListService{
		repoTweet,
	}
}

func (s TweetsListService) AllByPagedUser(userID string, page int64) ([]tweet.Tweet, error) {
	if len(userID) == 0 {
		return []tweet.Tweet{}, errors.New("the userId is empty")
	}

	return s.RepositoryTweet.AllByPagedUser(userID, page)
}
