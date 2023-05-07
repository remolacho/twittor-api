package tweet_service

import "twittor-api/domain/models/tweet"

type TweetsListService struct {
	RepositoryTweet tweet.ITweet
}

func NewList(repoTweet tweet.ITweet) *TweetsListService {
	return &TweetsListService{
		repoTweet,
	}
}

func (s TweetsListService) AllByPagedUser(userID string, page int64) ([]tweet.Tweet, error) {
	return s.RepositoryTweet.AllByPagedUser(userID, page)
}
