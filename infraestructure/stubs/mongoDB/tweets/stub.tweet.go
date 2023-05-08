package tweets

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor-api/domain/models/tweet"
	"twittor-api/domain/models/user"
	stubUser "twittor-api/infraestructure/stubs/mongoDB/users"
)

func (s *Stub) userNotFound() *tweet.Tweet {
	objectID, _ := primitive.ObjectIDFromHex("99999999999999999999")
	t := tweet.New(objectID.Hex())
	return t
}

func (s *Stub) userID() *tweet.Tweet {
	t := tweet.New("")
	t.Message = "My tweet test"

	return t
}

func (s *Stub) messageEmpty() *tweet.Tweet {
	t := tweet.New(s.currentUser().ID.Hex())
	return t
}

func (s *Stub) createTweet() *tweet.Tweet {
	t := tweet.New(s.currentUser().ID.Hex())
	t.Message = "My tweet test"
	t.CreatedAt = time.Now()

	return t
}

func (s *Stub) currentUser() *user.User {
	stub := stubUser.New()
	return stub.User("created")
}
