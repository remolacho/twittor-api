package factory_tweets

import (
	"twittor-api/infraestructure/stubs"
	stubTweetMongo "twittor-api/infraestructure/stubs/mongoDB/tweets"
)

func Build(opts ...string) stubs.IStubTweet {
	if opts == nil {
		return stubTweetMongo.New()
	}

	return nil
}
