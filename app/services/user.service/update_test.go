package user_service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/repositories/factories/repository.factory.user"
	stubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

var stubUpdate = stubFactoryUser.Build()

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
		input:         stubUpdate.User("created"),
		errorExpected: errors.New(""),
		description:   "The user not found in DB",
	},
	{
		name:          "email is empty",
		ID:            stubUpdate.User("email").ID.Hex(),
		input:         stubUpdate.User("email"),
		errorExpected: errors.New(""),
		description:   "It Was expected that user has email",
	},
	{
		name:          "update success",
		ID:            stubUpdate.User("created").ID.Hex(),
		input:         stubUpdate.User("created"),
		errorExpected: nil,
		description:   "the user was updated",
	},
}

func TestUpdate(t *testing.T) {
	mockFactory := repository_factory_user.Build("test")
	service := NewUpdate(mockFactory)

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
