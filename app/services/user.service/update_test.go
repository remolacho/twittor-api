package user_service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/repositories/factories/repository.factory.user"
	userMockRepository "twittor-api/infraestructure/repositories/mock/user.repository"
)

var testCasesUpdateUser = []struct {
	name          string
	input         *user.User
	ID            string
	errorExpected error
	description   string
}{
	{
		name:          "user not found",
		ID:            primitive.NewObjectID().Hex(),
		input:         userMockRepository.StubUser("created"),
		errorExpected: errors.New(""),
		description:   "The user not found in DB: ",
	},
	{
		name:          "email is empty",
		ID:            userMockRepository.GetID().Hex(),
		input:         userMockRepository.StubUser("email"),
		errorExpected: errors.New(""),
		description:   "It Was expected that user has email: ",
	},
	{
		name:          "update success",
		ID:            userMockRepository.GetID().Hex(),
		input:         userMockRepository.StubUser("created"),
		errorExpected: nil,
		description:   "the user was updated",
	},
}

func TestUpdate(t *testing.T) {
	mockFactory := repository_factory_user.Build("test")
	service := NewUpdateUser(mockFactory)

	for _, tc := range testCasesUpdateUser {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := service.Update(tc.ID, tc.input)
			if reflect.TypeOf(tc.errorExpected) != reflect.TypeOf(err) {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
