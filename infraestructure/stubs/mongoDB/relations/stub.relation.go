package relations

import (
	"twittor-api/domain/models/relation"
	StubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

func (s *Stub) createRelation() *relation.Relation {
	var stubUserCreate = StubFactoryUser.Build()
	user := stubUserCreate.User("created")

	t := relation.New()
	t.UserID = user.ID.Hex() + "1"
	t.UserRelationID = user.ID.Hex()

	return t
}
