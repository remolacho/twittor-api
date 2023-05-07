package jwt

import (
	"errors"
	"reflect"
	"testing"
	"twittor-api/domain/models/user"
	userMockRepository "twittor-api/infraestructure/stubs/users"
)

var testCasesEncode = []struct {
	name          string
	input         user.User
	expectedError error
	description   string
}{
	{
		name:          "User nil, Jwt error",
		input:         *userMockRepository.StubUser("email"),
		expectedError: errors.New(""),
		description:   "the user is nil",
	},
	{
		name:          "the jwt success",
		input:         *userMockRepository.StubUser("created"),
		expectedError: nil,
		description:   "the jwt is success",
	},
}

func TestEncode(t *testing.T) {
	service := NewJwt()

	for _, tc := range testCasesEncode {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := service.Encode(tc.input)

			if reflect.TypeOf(tc.expectedError) != reflect.TypeOf(err) {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}

var testCasesDecode = []struct {
	name          string
	user          user.User
	bearer        string
	expectedError error
	description   string
}{
	{
		name:          "the jwt bad format",
		user:          *userMockRepository.StubUser("created"),
		bearer:        "",
		expectedError: errors.New(""),
		description:   "the jwt bad format",
	},
	{
		name:          "the jwt is success decode",
		user:          *userMockRepository.StubUser("created"),
		bearer:        "bearer",
		expectedError: nil,
		description:   "the jwt is success decode",
	},
}

func TestDecode(t *testing.T) {
	service := NewJwt()

	for _, tc := range testCasesDecode {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			token, _ := service.Encode(tc.user)

			var jwt string

			if tc.bearer != "" {
				jwt = tc.bearer + " " + token
			} else {
				jwt = token
			}

			_, err := service.Decode(jwt)

			if reflect.TypeOf(tc.expectedError) != reflect.TypeOf(err) {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
