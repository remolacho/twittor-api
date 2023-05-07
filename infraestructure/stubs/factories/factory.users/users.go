package factory_users

import (
	"twittor-api/infraestructure/stubs"
	stubUserMongo "twittor-api/infraestructure/stubs/mongoDB/users"
)

func Build(opts ...string) stubs.IStubUser {
	if opts == nil {
		return stubUserMongo.New()
	}

	return nil
}
