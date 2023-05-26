package repository_factory_follower

import (
	"twittor-api/domain/models/follower"
	followerMockRepository "twittor-api/infraestructure/repositories/mock/follower.repository"
	followerMongoRepository "twittor-api/infraestructure/repositories/mongoDB/follower.repository"
)

func Build(opts ...string) follower.IFollow {
	if opts == nil {
		return followerMongoRepository.New()
	}

	if opts[0] == "test" {
		return followerMockRepository.New()
	}

	return nil
}
