package follower_service

import (
	"testing"
	repositoryFactoryFollow "twittor-api/infraestructure/repositories/factories/repository.factory.follower"
	StubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

var stubUserDelete = StubFactoryUser.Build()

type InputDestroy struct {
	FollowUserId string
	UserID       string
}

var testCasesRelationDelete = []struct {
	name        string
	input       InputDestroy
	expected    bool
	description string
}{
	{
		name: "the follow not found",
		input: InputDestroy{
			"",
			stubUserDelete.User("created").ID.Hex(),
		},
		expected:    false,
		description: "It follow was not possibility because id is empty",
	},
	{
		name: "belongs to the user",
		input: InputDestroy{
			stubUserDelete.User("created").ID.Hex(),
			stubUserDelete.User("other").ID.Hex(),
		},
		expected:    false,
		description: "It is not belongs to user",
	},
	{
		name: "delete follow",
		input: InputDestroy{
			stubUserDelete.User("created").ID.Hex(),
			stubUserDelete.User("created").ID.Hex() + "1",
		},
		expected:    true,
		description: "It was deleted with success",
	},
}

func TestDestroy(t *testing.T) {
	mockFactoryRelation := repositoryFactoryFollow.Build("test")

	service := NewDelete(mockFactoryRelation)

	for _, tc := range testCasesRelationDelete {
		t.Run(tc.name, func(t *testing.T) {
			got, err := service.Destroy(tc.input.UserID, tc.input.FollowUserId)
			if got != tc.expected {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
