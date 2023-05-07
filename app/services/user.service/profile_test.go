package user_service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	mockUserRepository "twittor-api/infraestructure/stubs/users"
)

var testCases = []struct {
	name          string
	ID            string
	expectedError error
	description   string
}{
	{
		name:          "the profile ID error",
		ID:            primitive.NewObjectID().Hex(),
		expectedError: errors.New(""),
		description:   "the user not found",
	},
	{
		name:          "the profile ID",
		ID:            mockUserRepository.GetID().Hex(),
		expectedError: nil,
		description:   "the profile is found",
	},
}

func TestProfile(t *testing.T) {
	mockFactory := repositoryFactoryUser.Build("test")
	profile := NewProfile(mockFactory)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := profile.Get(tc.ID)

			if reflect.TypeOf(tc.expectedError) != reflect.TypeOf(err) {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
