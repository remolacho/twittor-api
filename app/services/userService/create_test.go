package userService

import (
	"errors"
	"reflect"
	"testing"
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/repositories/factories/repositoryFactoryUser"
	userMockRepository "twittor-api/infraestructure/repositories/mock/userRepository"
)

var testCases = []struct {
	name          string
	input         *user.User
	errorExpected error
	message       string
}{
	{
		name:          "email is empty",
		input:         userMockRepository.GetUserWithOutEmail(),
		errorExpected: errors.New(""),
		message:       "It Was expected that user has email: ",
	},
	{
		name:          "password size",
		input:         userMockRepository.GetUserSizeLessPassword(),
		errorExpected: errors.New(""),
		message:       "It Was expected that user has password had with 6 or more chars: ",
	},
	{
		name:          "user already exists",
		input:         userMockRepository.RegisterUser(),
		errorExpected: errors.New(""),
		message:       "user was expected to not exist: ",
	},
	{
		name:          "user was created",
		input:         userMockRepository.NewUser(),
		errorExpected: nil,
		message:       "user was created!!! ",
	},
}

func TestCreate(t *testing.T) {
	mockFactory := repositoryFactoryUser.Build("test")
	service := NewUser(mockFactory)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := service.Create(tc.input)
			if reflect.TypeOf(tc.errorExpected) != reflect.TypeOf(err) {
				t.Errorf("%s: %v", tc.message, err)
			}
		})
	}
}
