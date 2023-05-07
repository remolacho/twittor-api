package tweets

import (
	"twittor-api/domain/models/tweet"
)

func StubTweets(t string) []tweet.Tweet {
	var tweets []tweet.Tweet

	switch t {
	case "list":
		tweets = list()
		break
	}

	return tweets
}

func list() []tweet.Tweet {
	var tweets []tweet.Tweet

	for i := 0; i <= 10; i++ {
		tweets = append(tweets, *StubTweet("createTweet"))
	}

	return tweets
}
