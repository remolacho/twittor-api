package repository_factory_user

import (
	"twittor-api/domain/models/user"
	userMockRepository "twittor-api/infraestructure/repositories/mock/user.repository"
	userMongoRepository "twittor-api/infraestructure/repositories/mongoDB/user.repository"
)

func Build(opts ...string) user.IUser {
	if opts == nil {
		return userMongoRepository.New()
	}

	if opts[0] == "test" {
		return userMockRepository.New()
	}

	return nil
}
