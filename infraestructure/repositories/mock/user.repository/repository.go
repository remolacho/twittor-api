package user_repository

import (
	"twittor-api/infraestructure/stubs"
	stubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

type UserRepository struct {
	Stub stubs.IStubUser
}

func New() *UserRepository {
	return &UserRepository{
		stubFactoryUser.Build(),
	}
}
