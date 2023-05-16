package relation_repository

import (
	"twittor-api/domain/models/relation"
	factoryStubRelation "twittor-api/infraestructure/stubs/factories/factory.relations"
)

func (r *RelationRepository) FindByObject(t *relation.Relation) bool {
	factory := factoryStubRelation.Build()
	relationObj := factory.Relation("createRelation")

	if (relationObj.UserID == t.UserID) && (relationObj.UserRelationID == t.UserRelationID) {
		return true
	}

	return false
}
