package followers

import (
	"twittor-api/domain/models/follow"
	StubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

func (s *Stub) createFollow() *follow.Follow {
	var stubUserCreate = StubFactoryUser.Build()
	user := stubUserCreate.User("created")

	t := follow.New()
	t.UserID = user.ID.Hex() + "1"
	t.FollowUserID = user.ID.Hex()

	return t
}
