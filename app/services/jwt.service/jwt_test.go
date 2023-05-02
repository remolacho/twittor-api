package jwt_service

import (
	"errors"
	"reflect"
	"testing"
	"twittor-api/domain/models/user"
	userMockRepository "twittor-api/infraestructure/repositories/mock/user.repository"
)

var testCasesJwt = []struct {
	name          string
	input         *user.User
	expectedError error
	description   string
}{
	{
		name:          "User nil, Jwt error",
		input:         nil,
		expectedError: errors.New(""),
		description:   "the user is nil",
	},
	{
		name:          "the jwt success",
		input:         userMockRepository.StubUser("created"),
		expectedError: nil,
		description:   "the jwt is success",
	},
}

func TestEncode(t *testing.T) {
	for _, tc := range testCasesJwt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			service := NewJwt(tc.input)
			_, err := service.Encode()

			if reflect.TypeOf(tc.expectedError) != reflect.TypeOf(err) {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
