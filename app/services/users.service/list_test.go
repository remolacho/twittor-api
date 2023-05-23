package users_service

import (
	"testing"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	stubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

func TestFollowed(t *testing.T) {
	stubUser := stubFactoryUser.Build()
	user := stubUser.User("created")
	UserID := user.ID.Hex() + "1"
	followerID := user.ID.Hex()

	mockFactory := repositoryFactoryUser.Build("test")
	service := New(mockFactory)

	t.Run("list users followed", func(t *testing.T) {
		users, _ := service.Followed(UserID, 1, "")

		if len(users) == 0 || users[0].ID.Hex() != followerID {
			t.Errorf("the user followed not found")
		}
	})
}

func TestUnFollowed(t *testing.T) {
	stubUser := stubFactoryUser.Build()
	user := stubUser.User("created")
	userID := user.ID.Hex()
	followerID := user.ID.Hex() + "2"

	mockFactory := repositoryFactoryUser.Build("test")
	service := New(mockFactory)

	t.Run("list users unfollowed", func(t *testing.T) {
		users, _ := service.UnFollowed(userID, 1, "")

		if len(users) == 0 || users[0].ID.Hex() == followerID {
			t.Errorf("the user unfollowed not found")
		}
	})
}
