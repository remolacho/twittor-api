package tweet_service

import (
	"testing"
	repositoryFactoryTweet "twittor-api/infraestructure/repositories/factories/repository.factory.tweet"
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
		name:        "list empty",
		input:       input{userID: "999999999999999999", page: 1},
		description: "the list is empty, the user not found tweets",
	},
	{
		name:        "userID is empty",
		input:       input{userID: "", page: 1},
		description: "list empty, userID is empty",
	},
}

func TestAllByPagedUserError(t *testing.T) {
	mockFactoryTweet := repositoryFactoryTweet.Build("test")

	service := NewList(mockFactoryTweet)

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

	service := NewList(mockFactoryTweet)

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
