package followers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"twittor-api/domain/models/follower"
	StubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

func (s *Stub) createFollower() *follower.Follower {
	var stubUserCreate = StubFactoryUser.Build()
	user := stubUserCreate.User("created")

	t := follower.New()
	t.ID = s.getID()
	t.UserID = user.ID.Hex() + "1"
	t.FollowUserID = user.ID.Hex()

	return t
}

func (s *Stub) createFollowerII() *follower.Follower {
	var stubUserCreate = StubFactoryUser.Build()
	user := stubUserCreate.User("created")

	t := follower.New()
	t.ID = s.getID()
	t.UserID = user.ID.Hex()
	t.FollowUserID = user.ID.Hex() + "1"

	return t
}

func (s *Stub) getID() primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex("5c71f03ccfee587e4212ad90")
	return objectID
}
