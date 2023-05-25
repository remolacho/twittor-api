package follow_service

import "twittor-api/domain/models/follow"

type ListTweetsService struct {
	Repository follow.IFollow
}

func NewListTweets(repository follow.IFollow) *ListTweetsService {
	return &ListTweetsService{
		repository,
	}
}

func (s *ListTweetsService) ListTweets(userID string, page int64) ([]follow.HasOneTweet, bool, error) {
	followers, err := s.Repository.IncludeTweets(userID, page)

	if err != nil {
		return followers, false, err
	}

	return followers, true, nil
}
