package followers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"twittor-api/domain/models/follow"
	StubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

func (s *Stub) createFollow() *follow.Follow {
	var stubUserCreate = StubFactoryUser.Build()
	user := stubUserCreate.User("created")

	t := follow.New()
	t.ID = s.getID()
	t.UserID = user.ID.Hex() + "1"
	t.FollowUserID = user.ID.Hex()

	return t
}

func (s *Stub) createFollowII() *follow.Follow {
	var stubUserCreate = StubFactoryUser.Build()
	user := stubUserCreate.User("created")

	t := follow.New()
	t.ID = s.getID()
	t.UserID = user.ID.Hex()
	t.FollowUserID = user.ID.Hex() + "1"

	return t
}

func (s *Stub) getID() primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex("0000000000000000000000")
	return objectID
}
