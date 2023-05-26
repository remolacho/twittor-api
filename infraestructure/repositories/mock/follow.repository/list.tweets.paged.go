package follow_repository

import (
	"errors"
	"twittor-api/domain/models/follow"
	stubFactoryFollower "twittor-api/infraestructure/stubs/factories/factory.followers"
	stubFactoryTweet "twittor-api/infraestructure/stubs/factories/factory.tweets"
)

func (r *FollowRepository) IncludeTweet(userID string, page int64) ([]follow.HasOneTweet, error) {
	var followers []follow.HasOneTweet

	stub := stubFactoryFollower.Build()
	stubTweet := stubFactoryTweet.Build()
	followStub := stub.Follow("created")
	tweetStub := stubTweet.Tweet("createTweet")

	if userID != followStub.UserID {
		return followers, errors.New("the user not follow to tweet")
	}

	followerTweet := follow.NewHaOneTweet()

	followerTweet.ID = followStub.ID
	followerTweet.UserID = followStub.UserID
	followerTweet.FollowUserID = followStub.FollowUserID
	followerTweet.Tweet = *tweetStub

	followers = append(followers, *followerTweet)
	return followers, nil
}
