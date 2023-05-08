package repository_factory_tweet

import (
	"twittor-api/domain/models/tweet"
	tweetMockRepository "twittor-api/infraestructure/repositories/mock/tweet.repository"
	tweetMongoRepository "twittor-api/infraestructure/repositories/mongoDB/tweet.repository"
)

func Build(opts ...string) tweet.ITweet {
	if opts == nil {
		return tweetMongoRepository.New()
	}

	if opts[0] == "test" {
		return tweetMockRepository.New()
	}

	return nil
}
