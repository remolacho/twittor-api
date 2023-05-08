package tweet_service

import (
	"testing"
	repositoryFactoryTweet "twittor-api/infraestructure/repositories/factories/repository.factory.tweet"
	StubFactoryTweet "twittor-api/infraestructure/stubs/factories/factory.tweets"
)

var stubDelete = StubFactoryTweet.Build()

type inputDelete struct {
	tweetID string
	userID  string
}

var testCasesDeleteTweetValidation = []struct {
	name        string
	expected    bool
	input       inputDelete
	description string
}{
	{
		name:        "the user isn't owner",
		expected:    false,
		input:       inputDelete{tweetID: stubDelete.Tweet("createTweet").ID.Hex(), userID: "999999999999999999"},
		description: "this tweet not belong to user",
	},
	{
		name:        "the id tweet not found",
		expected:    false,
		input:       inputDelete{tweetID: "999999999999999999", userID: stubDelete.Tweet("createTweet").UserID},
		description: "this tweet not found by id",
	},
	{
		name:        "the id tweet is empty",
		expected:    false,
		input:       inputDelete{tweetID: "", userID: stubDelete.Tweet("createTweet").UserID},
		description: "this tweet not found by id empty",
	},
}

func TestDeleteError(t *testing.T) {
	mockFactoryTweet := repositoryFactoryTweet.Build("test")

	service := NewDelete(mockFactoryTweet)

	for _, tc := range testCasesDeleteTweetValidation {
		t.Run(tc.name, func(t *testing.T) {

			got, _ := service.Delete(tc.input.tweetID, tc.input.userID)

			if got != tc.expected {
				t.Errorf("%s", tc.description)
			}
		})
	}
}

func TestDeleteSuccess(t *testing.T) {
	testCase := struct {
		name        string
		expected    bool
		input       inputDelete
		description string
	}{
		name:        "delete success",
		expected:    true,
		input:       inputDelete{tweetID: stubDelete.Tweet("createTweet").ID.Hex(), userID: stubDelete.Tweet("createTweet").UserID},
		description: "this tweet was deleted success",
	}

	mockFactoryTweet := repositoryFactoryTweet.Build("test")

	service := NewDelete(mockFactoryTweet)

	t.Run(testCase.name, func(t *testing.T) {
		got, _ := service.Delete(testCase.input.tweetID, testCase.input.userID)

		if got != testCase.expected {
			t.Errorf("%s", testCase.description)
		}
	})
}
