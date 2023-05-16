package relation_service

import (
	"testing"
	repositoryFactoryRelation "twittor-api/infraestructure/repositories/factories/repository.factory.relation"
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
		name: "the relation not found",
		input: Input{
			stubUserCreate.User("created").ID.Hex(),
			"",
		},
		expected:    false,
		description: "It relation was not possibility because user relation id was empty",
	},
	{
		name: "the relation not found",
		input: Input{
			stubUserCreate.User("created").ID.Hex(),
			stubUserCreate.User("created").ID.Hex() + "1",
		},
		expected:    false,
		description: "It relation was not possibility because user not exists",
	},
	{
		name: "the relation not found",
		input: Input{
			stubUserCreate.User("created").ID.Hex(),
			stubUserCreate.User("created").ID.Hex(),
		},
		expected:    false,
		description: "It relation cannot create, follow yourself",
	},
	{
		name: "the relation ready exists",
		input: Input{
			stubUserCreate.User("created").ID.Hex() + "1",
			stubUserCreate.User("created").ID.Hex(),
		},
		expected:    false,
		description: "It relation was not created",
	},
	{
		name: "the relation success",
		input: Input{
			stubUserCreate.User("created").ID.Hex() + "2",
			stubUserCreate.User("created").ID.Hex(),
		},
		expected:    true,
		description: "It relation was created",
	},
}

func TestCreate(t *testing.T) {
	mockFactoryUser := repositoryFactoryUser.Build("test")
	mockFactoryRelation := repositoryFactoryRelation.Build("test")

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
