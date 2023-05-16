package repository_factory_relation

import (
	"twittor-api/domain/models/relation"
	relationMockRepository "twittor-api/infraestructure/repositories/mock/relation.repository"
	relationMongoRepository "twittor-api/infraestructure/repositories/mongoDB/relation.repository"
)

func Build(opts ...string) relation.IRelation {
	if opts == nil {
		return relationMongoRepository.New()
	}

	if opts[0] == "test" {
		return relationMockRepository.New()
	}

	return nil
}
