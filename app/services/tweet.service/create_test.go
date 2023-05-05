package tweet_service

import (
	"errors"
	"reflect"
	"testing"
	"twittor-api/domain/models/tweet"
	repositoryFactoryTweet "twittor-api/infraestructure/repositories/factories/repository.factory.tweet"
	tweetMockRepository "twittor-api/infraestructure/repositories/mock/tweet.repository"
)

var testCasesCreateTweet = []struct {
	name          string
	input         *tweet.Tweet
	errorExpected error
	description   string
}{
	{
		name:          "message is empty",
		input:         tweetMockRepository.StubTweet("message"),
		errorExpected: errors.New(""),
		description:   "It Was expected that tweet has message: ",
	},
	{
		name:          "userId is empty",
		input:         tweetMockRepository.StubTweet("message"),
		errorExpected: errors.New(""),
		description:   "It Was expected that tweet has userID: ",
	},
	{
		name:          "tweet created",
		input:         tweetMockRepository.StubTweet("new"),
		errorExpected: nil,
		description:   "the tweet was created with success: ",
	},
}

func TestCreate(t *testing.T) {
	mockFactory := repositoryFactoryTweet.Build("test")
	service := NewTweet(mockFactory)

	for _, tc := range testCasesCreateTweet {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := service.Create(tc.input)
			if reflect.TypeOf(tc.errorExpected) != reflect.TypeOf(err) {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
