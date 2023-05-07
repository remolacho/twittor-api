package tweets

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/tweet"
	"twittor-api/domain/models/user"
	userRepository "twittor-api/infraestructure/stubs/users"
)

func StubTweet(t string) *tweet.Tweet {
	var _tweet *tweet.Tweet

	switch t {
	case "messageEmpty":
		_tweet = messageEmpty()
		break
	case "userIdEmpty":
		_tweet = userID()
		break
	case "userNotFound":
		_tweet = userNotFound()
		break
	case "createTweet":
		_tweet = createTweet()
		break
	}

	return _tweet
}

func userNotFound() *tweet.Tweet {
	objectID, _ := primitive.ObjectIDFromHex("99999999999999999999")
	t := tweet.New(objectID.Hex())
	return t
}

func userID() *tweet.Tweet {
	t := tweet.New("")
	t.Message = "My tweet test"

	return t
}

func messageEmpty() *tweet.Tweet {
	t := tweet.New(currentUser().ID.Hex())
	return t
}

func createTweet() *tweet.Tweet {
	t := tweet.New(currentUser().ID.Hex())
	t.Message = "My tweet test"
	t.CreatedAt = time.Now()

	return t
}

func currentUser() *user.User {
	return userRepository.StubUser("created")
}
