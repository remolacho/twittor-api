package follow_repository

import (
	"errors"
	"twittor-api/domain/models/follow"
	stubFactoryFollower "twittor-api/infraestructure/stubs/factories/factory.followers"
	stubFactoryTweet "twittor-api/infraestructure/stubs/factories/factory.tweets"
)

func (r *FollowRepository) IncludeTweets(userID string, page int64) ([]follow.Follow, error) {
	var followers []follow.Follow

	stub := stubFactoryFollower.Build()
	stubTweet := stubFactoryTweet.Build()
	followStub := stub.Follow("created")
	tweetStub := stubTweet.Tweet("createTweet")

	if userID != followStub.UserID {
		return followers, errors.New("the user not follow to tweet")
	}

	followStub.Tweet = tweetStub

	followers = append(followers, *followStub)
	return followers, nil
}
