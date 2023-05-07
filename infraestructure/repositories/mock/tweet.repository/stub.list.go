package tweet_repository

import "twittor-api/domain/models/tweet"

func StubTweets(t string) []tweet.Tweet {
	var tweets []tweet.Tweet

	switch t {
	case "list":
		tweets = list5()
		break
	}

	return tweets
}

func list5() []tweet.Tweet {
	var tweets []tweet.Tweet

	for i := 0; i <= 10; i++ {
		tweets = append(tweets, *StubTweet("createTweet"))
	}

	return tweets
}
