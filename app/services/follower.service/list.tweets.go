package follower_service

import "twittor-api/domain/models/follower"

type ListTweetsService struct {
	Repository follower.IFollow
}

func NewListTweets(repository follower.IFollow) *ListTweetsService {
	return &ListTweetsService{
		repository,
	}
}

func (s *ListTweetsService) ListTweets(userID string, page int64) ([]follower.HasOneTweet, bool, error) {
	followers, err := s.Repository.IncludeTweet(userID, page)

	if err != nil {
		return followers, false, err
	}

	return followers, true, nil
}
