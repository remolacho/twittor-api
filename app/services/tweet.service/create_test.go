package tweet_service

import (
	"testing"
	"twittor-api/domain/models/tweet"
	repositoryFactoryTweet "twittor-api/infraestructure/repositories/factories/repository.factory.tweet"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	tweetMockRepository "twittor-api/infraestructure/repositories/mock/tweet.repository"
)

var testCasesCreateTweetValidation = []struct {
	name        string
	input       *tweet.Tweet
	description string
}{
	{
		name:        "message is empty",
		input:       tweetMockRepository.StubTweet("messageEmpty"),
		description: "It Was expected that tweet has message",
	},
	{
		name:        "userId is empty",
		input:       tweetMockRepository.StubTweet("userIdEmpty"),
		description: "It Was expected that tweet has userID",
	},
	{
		name:        "user not found",
		input:       tweetMockRepository.StubTweet("userNotFound"),
		description: "The user not found",
	},
}

func TestCreateError(t *testing.T) {
	mockFactoryTweet := repositoryFactoryTweet.Build("test")
	mockFactoryUser := repositoryFactoryUser.Build("test")

	service := NewTweet(mockFactoryTweet, mockFactoryUser)

	for _, tc := range testCasesCreateTweetValidation {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := service.Create(tc.input)
			if err == nil {
				t.Errorf("%s", tc.description)
			}
		})
	}
}

func TestCreateSuccess(t *testing.T) {
	test := struct {
		name        string
		input       *tweet.Tweet
		description string
	}{
		name:        "tweet created",
		input:       tweetMockRepository.StubTweet("createTweet"),
		description: "the tweet was created with success",
	}

	mockFactoryTweet := repositoryFactoryTweet.Build("test")
	mockFactoryUser := repositoryFactoryUser.Build("test")

	service := NewTweet(mockFactoryTweet, mockFactoryUser)

	t.Run(test.name, func(t *testing.T) {
		_, err := service.Create(test.input)
		if err != nil {
			t.Errorf("%s: %v", test.description, err)
		}
	})
}
