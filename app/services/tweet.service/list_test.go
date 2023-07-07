package tweet_service

import (
	"testing"
	repositoryFactoryTweet "twittor-api/infraestructure/repositories/factories/repository.factory.tweet"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	StubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

var stub = StubFactoryUser.Build()

type input struct {
	userID string
	page   int64
}

var testCasesListTweetsValidation = []struct {
	name        string
	input       input
	description string
}{
	{
		name:        "userID is empty",
		input:       input{userID: "", page: 1},
		description: "list empty, userID is empty",
	},
	{
		name:        "user not found empty",
		input:       input{userID: "999999999999999999", page: 1},
		description: "the user no exist",
	},
}

func TestAllByPagedUserError(t *testing.T) {
	mockFactoryTweet := repositoryFactoryTweet.Build("test")
	mockFactoryUser := repositoryFactoryUser.Build("test")
	service := NewList(mockFactoryTweet, mockFactoryUser)

	for _, tc := range testCasesListTweetsValidation {
		t.Run(tc.name, func(t *testing.T) {
			_, err := service.AllByPagedUser(tc.input.userID, tc.input.page)

			if err == nil {
				t.Errorf("%s:", tc.description)
			}
		})
	}
}

func TestAllByPagedUserSuccess(t *testing.T) {
	mockFactoryTweet := repositoryFactoryTweet.Build("test")
	mockFactoryUser := repositoryFactoryUser.Build("test")
	service := NewList(mockFactoryTweet, mockFactoryUser)

	var testCase = struct {
		name        string
		input       input
		description string
	}{
		name:        "list success",
		input:       input{userID: stub.User("created").ID.Hex(), page: 1},
		description: "the list is success ",
	}

	t.Run(testCase.name, func(t *testing.T) {
		_, err := service.AllByPagedUser(testCase.input.userID, testCase.input.page)

		if err != nil {
			t.Errorf("%s: %v", testCase.description, err.Error())
		}
	})
}
