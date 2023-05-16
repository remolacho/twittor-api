package repository_factory_follow

import (
	"twittor-api/domain/models/follow"
	relationMockRepository "twittor-api/infraestructure/repositories/mock/follow.repository"
	relationMongoRepository "twittor-api/infraestructure/repositories/mongoDB/follow.repository"
)

func Build(opts ...string) follow.IFollow {
	if opts == nil {
		return relationMongoRepository.New()
	}

	if opts[0] == "test" {
		return relationMockRepository.New()
	}

	return nil
}
