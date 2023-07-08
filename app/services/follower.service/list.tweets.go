package follower_service

import (
	"time"
	"twittor-api/domain/models/follower"
	"twittor-api/domain/models/user"
)

type ListTweetsService struct {
	RepositoryFollow follower.IFollow
	RepositoryUser   user.IUser
}

type ResponseTweet struct {
	ID        string    `bson:"_id, omitempty" json:"id"`
	Message   string    `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt,omitempty"`
	user.User `bson:"user" json:"user,omitempty"`
}

func NewListTweets(repositoryFollow follower.IFollow, repositoryUser user.IUser) *ListTweetsService {
	return &ListTweetsService{
		repositoryFollow,
		repositoryUser,
	}
}

func (s *ListTweetsService) ListTweets(userID string, page int64) ([]ResponseTweet, bool, error) {
	followers, err := s.RepositoryFollow.IncludeTweet(userID, page)

	if err != nil {
		return []ResponseTweet{}, false, err
	}

	return s.parseTweets(followers)
}

func (s *ListTweetsService) parseTweets(followers []follower.HasOneTweet) ([]ResponseTweet, bool, error) {
	var list []ResponseTweet

	for _, f := range followers {
		currentTweet := ResponseTweet{}

		currentUser := s.getUser(f.FollowUserID)
		currentTweet.ID = f.Tweet.ID
		currentTweet.CreatedAt = f.Tweet.CreatedAt
		currentTweet.Message = f.Tweet.Message
		currentTweet.User.ID = currentUser.ID
		currentTweet.User.Name = currentUser.Name
		currentTweet.User.LastName = currentUser.LastName

		list = append(list, currentTweet)
	}

	return list, true, nil
}

func (s *ListTweetsService) getUser(userID string) *user.User {
	currentUser, err := s.RepositoryUser.Find(userID)

	if err != nil {
		return &user.User{}
	}

	return currentUser
}
