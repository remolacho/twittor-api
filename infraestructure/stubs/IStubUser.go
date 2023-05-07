package stubs

import "twittor-api/domain/models/user"

type IStubUser interface {
	User(t string) *user.User
}
