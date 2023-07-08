package tweet_service

import (
	"errors"
	"time"
	"twittor-api/domain/models/tweet"
	"twittor-api/domain/models/user"
)

type TweetsListService struct {
	RepositoryTweet tweet.ITweet
	RepositoryUser  user.IUser
}

type ResponseTweet struct {
	ID        string    `bson:"_id, omitempty" json:"id"`
	Message   string    `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt,omitempty"`
	user.User `bson:"user" json:"user,omitempty"`
}

func NewList(repoTweet tweet.ITweet, repoUser user.IUser) *TweetsListService {
	return &TweetsListService{
		repoTweet,
		repoUser,
	}
}

func (s TweetsListService) AllByPagedUser(userID string, page int64) ([]ResponseTweet, error) {
	if len(userID) == 0 {
		return []ResponseTweet{}, errors.New("the userId is empty")
	}

	currentUser, err := s.RepositoryUser.Find(userID)

	if err != nil {
		return []ResponseTweet{}, err
	}

	tweets, errTweet := s.RepositoryTweet.AllByPagedUser(userID, page)

	if errTweet != nil {
		return []ResponseTweet{}, errTweet
	}

	return s.listTweets(tweets, currentUser)
}

func (s TweetsListService) listTweets(tweets []tweet.Tweet, currentUser *user.User) ([]ResponseTweet, error) {
	var list []ResponseTweet

	for _, t := range tweets {
		currentTweet := ResponseTweet{}

		currentTweet.ID = t.ID
		currentTweet.CreatedAt = t.CreatedAt
		currentTweet.Message = t.Message
		currentTweet.User.ID = currentUser.ID
		currentTweet.User.Name = currentUser.Name
		currentTweet.User.LastName = currentUser.LastName

		list = append(list, currentTweet)
	}

	return list, nil
}
