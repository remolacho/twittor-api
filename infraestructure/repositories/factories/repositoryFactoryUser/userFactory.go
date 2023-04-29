package repositoryFactoryUser

import (
	"twittor-api/domain/models/user"
	userMockRepository "twittor-api/infraestructure/repositories/mock/userRepository"
	userMongoRepository "twittor-api/infraestructure/repositories/mongoDB/userRepository"
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
