package userService

import (
	"errors"
	"reflect"
	"testing"
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/repositories/factories/repositoryFactoryUser"
)

var testCases = []struct {
	name          string
	input         user.User
	errorExpected error
	message       string
}{
	{
		name: "email is empty",
		input: user.User{
			Email:    "",
			Name:     "Jonathan",
			LastName: "Rojas",
			Password: "123456",
		},
		errorExpected: errors.New(""),
		message:       "It Was expected that user has email: ",
	},
	{
		name: "password size",
		input: user.User{
			Email:    "jondoe@gmail.com",
			Name:     "Jonathan",
			LastName: "Rojas",
			Password: "12345",
		},
		errorExpected: errors.New(""),
		message:       "It Was expected that user has password had with 6 or more chars: ",
	},
	{
		name: "user already exists",
		input: user.User{
			Email:    "jonathangrh.25@gmail.com",
			Name:     "Jonathan",
			LastName: "Rojas",
			Password: "123456",
		},
		errorExpected: errors.New(""),
		message:       "user was expected to not exist: ",
	},
	{
		name: "user was created",
		input: user.User{
			Email:    "jhondoe@gmail.com",
			Name:     "Jonathan",
			LastName: "Rojas",
			Password: "123456",
		},
		errorExpected: nil,
		message:       "user was created!!! ",
	},
}

func TestCreate(t *testing.T) {
	mockFactory := repositoryFactoryUser.Build("mock")
	service := NewUser(mockFactory)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := service.Create(&tc.input)
			if reflect.TypeOf(tc.errorExpected) != reflect.TypeOf(err) {
				t.Errorf("%s: %v", tc.message, err)
			}
		})
	}
}
