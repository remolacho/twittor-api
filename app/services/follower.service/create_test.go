package follower_service

import (
	"testing"
	repositoryFactoryFollow "twittor-api/infraestructure/repositories/factories/repository.factory.follower"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	StubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

var stubUserCreate = StubFactoryUser.Build()

type Input struct {
	userID         string
	userRelationID string
}

var testCasesRelationCreate = []struct {
	name        string
	input       Input
	expected    bool
	description string
}{
	{
		name: "the follow not found",
		input: Input{
			stubUserCreate.User("created").ID.Hex(),
			"",
		},
		expected:    false,
		description: "It follow was not possibility because user follow id was empty",
	},
	{
		name: "the follow not found",
		input: Input{
			stubUserCreate.User("created").ID.Hex(),
			stubUserCreate.User("created").ID.Hex() + "1",
		},
		expected:    false,
		description: "It follow was not possibility because user not exists",
	},
	{
		name: "the follow not found",
		input: Input{
			stubUserCreate.User("created").ID.Hex(),
			stubUserCreate.User("created").ID.Hex(),
		},
		expected:    false,
		description: "It follow cannot create, follow yourself",
	},
	{
		name: "the follow ready exists",
		input: Input{
			stubUserCreate.User("created").ID.Hex() + "1",
			stubUserCreate.User("created").ID.Hex(),
		},
		expected:    false,
		description: "It follow was not created",
	},
	{
		name: "the follow success",
		input: Input{
			stubUserCreate.User("created").ID.Hex() + "2",
			stubUserCreate.User("created").ID.Hex(),
		},
		expected:    true,
		description: "It follow was created",
	},
}

func TestCreate(t *testing.T) {
	mockFactoryUser := repositoryFactoryUser.Build("test")
	mockFactoryRelation := repositoryFactoryFollow.Build("test")

	service := NewCreate(mockFactoryUser, mockFactoryRelation)

	for _, tc := range testCasesRelationCreate {
		t.Run(tc.name, func(t *testing.T) {
			got, err := service.Create(tc.input.userID, tc.input.userRelationID)
			if got != tc.expected {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
