package user_repository

import (
	"twittor-api/domain/models/user"
	factoryStubFollowers "twittor-api/infraestructure/stubs/factories/factory.followers"
	factoryStubUsers "twittor-api/infraestructure/stubs/factories/factory.users"
)

func (r *UserRepository) AllUnFollowed(ID string, page int64, searchTerm string) ([]user.User, error) {
	var users []user.User

	factoryF := factoryStubFollowers.Build()
	factory := factoryStubUsers.Build()
	cursor := factory.Users()
	follower := factoryF.Follower("createdII")

	for _, c := range cursor {
		if ID == follower.UserID && follower.FollowUserID != c.ID.Hex() {
			users = append(users, c)
		}
	}

	return users, nil
}
