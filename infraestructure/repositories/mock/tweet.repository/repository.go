package tweet_repository

import (
	"twittor-api/infraestructure/stubs"
	stubFactoryTweet "twittor-api/infraestructure/stubs/factories/factory.tweets"
)

type TweetRepository struct {
	Stub stubs.IStubTweet
}

func New() *TweetRepository {
	return &TweetRepository{
		stubFactoryTweet.Build(),
	}
}
