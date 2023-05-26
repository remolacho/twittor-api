package follower_repository

import (
	"errors"
	"twittor-api/domain/models/follower"
	stubFactoryFollower "twittor-api/infraestructure/stubs/factories/factory.followers"
	stubFactoryTweet "twittor-api/infraestructure/stubs/factories/factory.tweets"
)

func (r *FollowRepository) IncludeTweet(userID string, page int64) ([]follower.HasOneTweet, error) {
	var followers []follower.HasOneTweet

	stub := stubFactoryFollower.Build()
	stubTweet := stubFactoryTweet.Build()
	followStub := stub.Follower("created")
	tweetStub := stubTweet.Tweet("createTweet")

	if userID != followStub.UserID {
		return followers, errors.New("the user not follow to tweet")
	}

	followerTweet := follower.NewHaOneTweet()

	followerTweet.ID = followStub.ID
	followerTweet.UserID = followStub.UserID
	followerTweet.FollowUserID = followStub.FollowUserID
	followerTweet.Tweet = *tweetStub

	followers = append(followers, *followerTweet)
	return followers, nil
}
