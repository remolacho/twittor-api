package user_service

import (
	"errors"
	"reflect"
	"testing"
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/repositories/factories/repository.factory.user"
	userMockRepository "twittor-api/infraestructure/repositories/mock/user.repository"
)

var testCasesCreateUser = []struct {
	name          string
	input         *user.User
	errorExpected error
	description   string
}{
	{
		name:          "email is empty",
		input:         userMockRepository.StubUser("email"),
		errorExpected: errors.New(""),
		description:   "It Was expected that user has email",
	},
	{
		name:          "password size",
		input:         userMockRepository.StubUser("password"),
		errorExpected: errors.New(""),
		description:   "It Was expected that user has password had with 6 or more chars",
	},
	{
		name:          "user already exists",
		input:         userMockRepository.StubUser("created"),
		errorExpected: errors.New(""),
		description:   "user was expected to not exist",
	},
	{
		name:          "user was created",
		input:         userMockRepository.StubUser("new"),
		errorExpected: nil,
		description:   "user was created!!! ",
	},
}

func TestCreate(t *testing.T) {
	mockFactory := repository_factory_user.Build("test")
	service := NewUser(mockFactory)

	for _, tc := range testCasesCreateUser {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := service.Create(tc.input)
			if reflect.TypeOf(tc.errorExpected) != reflect.TypeOf(err) {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
