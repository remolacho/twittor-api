package repository_factory_follow

import (
	"twittor-api/domain/models/follow"
	followMockRepository "twittor-api/infraestructure/repositories/mock/follow.repository"
	followMongoRepository "twittor-api/infraestructure/repositories/mongoDB/follow.repository"
)

func Build(opts ...string) follow.IFollow {
	if opts == nil {
		return followMongoRepository.New()
	}

	if opts[0] == "test" {
		return followMockRepository.New()
	}

	return nil
}
