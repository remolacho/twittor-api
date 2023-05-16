package factory_followers

import (
	"twittor-api/infraestructure/stubs"
	stubRelationMongo "twittor-api/infraestructure/stubs/mongoDB/followers"
)

func Build(opts ...string) stubs.IStubFollow {
	if opts == nil {
		return stubRelationMongo.New()
	}

	return nil
}
