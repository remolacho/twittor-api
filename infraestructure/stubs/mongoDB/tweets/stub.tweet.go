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

func (s *Stub) messageLimit() *tweet.Tweet {
	t := tweet.New(s.currentUser().ID.Hex())
	t.Message = "Lorem Ipsum es simplemente el texto de relleno de las imprentas y archivos de texto. Lorem Ipsum ha sido el texto de relleno estándar de las industrias desde el año 1500, cuando un impresor (N. del T. persona que se dedica a la imprenta) desconocido usó una galería de textos y los mezcló de tal manera que logró hacer un libro de textos especimen. No sólo sobrevivió 500 años."
	t.CreatedAt = time.Now()
	t.ID = "1"

	return t
}

func (s *Stub) createTweet() *tweet.Tweet {
	t := tweet.New(s.currentUser().ID.Hex())
	t.Message = "My tweet test"
	t.CreatedAt = time.Now()
	t.ID = "1"

	return t
}

func (s *Stub) currentUser() *user.User {
	stub := stubUser.New()
	return stub.User("created")
}
