package user_service

import (
	"testing"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	stubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

var stubProfile = stubFactoryUser.Build()

var testCasesProfileValidations = []struct {
	name        string
	ID          string
	description string
}{
	{
		name:        "the profile ID error",
		ID:          "999999999999999999999999",
		description: "the user not found",
	},
}

func TestProfileError(t *testing.T) {
	mockFactory := repositoryFactoryUser.Build("test")
	profile := NewProfile(mockFactory)

	for _, tc := range testCasesProfileValidations {
		t.Run(tc.name, func(t *testing.T) {
			_, err := profile.Get(tc.ID)

			if err == nil {
				t.Errorf("%s", tc.description)
			}
		})
	}
}

func TestProfileSuccess(t *testing.T) {
	test := struct {
		name        string
		ID          string
		description string
	}{
		name:        "the profile ID",
		ID:          stubProfile.User("created").ID.Hex(),
		description: "the profile is found",
	}

	mockFactory := repositoryFactoryUser.Build("test")
	profile := NewProfile(mockFactory)

	t.Run(test.name, func(t *testing.T) {
		_, err := profile.Get(test.ID)

		if err != nil {
			t.Errorf("%s: %v", test.description, err)
		}
	})
}
