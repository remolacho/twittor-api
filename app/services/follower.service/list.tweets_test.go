package follower_service

import (
	"testing"
	repositoryFactoryFollow "twittor-api/infraestructure/repositories/factories/repository.factory.follower"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	factoryStubUsers "twittor-api/infraestructure/stubs/factories/factory.users"
)

var stubUser = factoryStubUsers.Build()

type InputList struct {
	Page   int64
	UserID string
}

var testCasesList = []struct {
	name        string
	input       InputList
	expected    bool
	description string
}{
	{
		name:        "the list not found",
		description: "the list not found the usr invalid",
		expected:    false,
		input: InputList{
			Page:   1,
			UserID: stubUser.User("created").ID.Hex(),
		},
	},
	{
		name:        "the list  found",
		description: "the list found the usr valid",
		expected:    true,
		input: InputList{
			Page:   1,
			UserID: stubUser.User("created").ID.Hex() + "1",
		},
	},
}

func TestListTweets(t *testing.T) {
	mockFactoryFollower := repositoryFactoryFollow.Build("test")
	mockFactoryUser := repositoryFactoryUser.Build("test")
	service := NewListTweets(mockFactoryFollower, mockFactoryUser)

	for _, tc := range testCasesList {
		t.Run(tc.name, func(t *testing.T) {
			_, got, err := service.ListTweets(tc.input.UserID, tc.input.Page)

			if got != tc.expected {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
