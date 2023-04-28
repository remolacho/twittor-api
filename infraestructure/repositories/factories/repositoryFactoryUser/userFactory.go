package repositoryFactoryUser

import (
	"twittor-api/domain/models/user"
	userMongoRepository "twittor-api/infraestructure/repositories/mongoDB/userRepository"
)

func Build(opts ...string) user.IUser {
	if opts == nil {
		return userMongoRepository.New()
	}

	return nil
}
