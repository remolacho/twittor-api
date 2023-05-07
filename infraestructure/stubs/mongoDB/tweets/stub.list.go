package tweets

import (
	"twittor-api/domain/models/tweet"
)

func (s *Stub) list() []tweet.Tweet {
	var tweets []tweet.Tweet

	for i := 0; i <= 10; i++ {
		tweets = append(tweets, *s.Tweet("createTweet"))
	}

	return tweets
}
