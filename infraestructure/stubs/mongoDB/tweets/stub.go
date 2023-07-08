package tweets

import "twittor-api/domain/models/tweet"

type Stub struct{}

func New() *Stub {
	return &Stub{}
}

func (s *Stub) Tweet(t string) *tweet.Tweet {
	var _tweet *tweet.Tweet

	switch t {
	case "messageEmpty":
		_tweet = s.messageEmpty()
		break
	case "messageLimit":
		_tweet = s.messageLimit()
		break
	case "userIdEmpty":
		_tweet = s.userID()
		break
	case "userNotFound":
		_tweet = s.userNotFound()
		break
	case "createTweet":
		_tweet = s.createTweet()
		break
	}

	return _tweet
}

func (s *Stub) Tweets(t string) []tweet.Tweet {
	var tweets []tweet.Tweet

	switch t {
	case "list":
		tweets = s.list()
		break
	}

	return tweets
}
