package tweet_repository

import (
	"twittor-api/domain/models/tweet"
	"twittor-api/domain/models/user"
	userRepository "twittor-api/infraestructure/repositories/mock/user.repository"
)

func StubTweet(t string) *tweet.Tweet {
	var _tweet *tweet.Tweet

	switch t {
	case "message":
		_tweet = message()
		break
	case "userId":
		_tweet = userID()
		break
	case "created":
		break
	case "new":
		_tweet = newTweet()
		break
	}

	return _tweet
}

func userID() *tweet.Tweet {
	return &tweet.Tweet{
		UserID:  "",
		Message: "My tweet test",
	}
}

func message() *tweet.Tweet {
	return &tweet.Tweet{
		UserID:  currentUser().ID.Hex(),
		Message: "",
	}
}

func newTweet() *tweet.Tweet {
	return &tweet.Tweet{
		UserID:  currentUser().ID.Hex(),
		Message: "My tweet test",
	}
}

func currentUser() *user.User {
	return userRepository.StubUser("created")
}
