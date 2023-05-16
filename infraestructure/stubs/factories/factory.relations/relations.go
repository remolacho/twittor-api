package factory_relations

import (
	"twittor-api/infraestructure/stubs"
	stubRelationMongo "twittor-api/infraestructure/stubs/mongoDB/relations"
)

func Build(opts ...string) stubs.IStubRelation {
	if opts == nil {
		return stubRelationMongo.New()
	}

	return nil
}
