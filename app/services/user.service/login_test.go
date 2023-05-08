package user_service

import (
	"testing"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

type input struct {
	email    string
	password string
}

var testCasesSignInUser = []struct {
	name        string
	input       input
	expected    bool
	description string
}{
	{
		name: "Login not valid, email error",
		input: input{
			email:    "jonathan_error@gmail.com",
			password: "123456",
		},
		expected:    false,
		description: "the email not found",
	},
	{
		name: "Login not valid, password error",
		input: input{
			email:    "jonathangrh.25@gmail.com",
			password: "1234565",
		},
		expected:    false,
		description: "the password not found",
	},
	{
		name: "Login success!!!",
		input: input{
			email:    "jonathangrh.25@gmail.com",
			password: "123456",
		},
		expected:    true,
		description: "the user login success",
	},
}

func TestLogin(t *testing.T) {
	mockFactory := repositoryFactoryUser.Build("test")
	service := NewLogin(mockFactory)

	for _, tc := range testCasesSignInUser {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, got := service.Login(tc.input.email, tc.input.password)
			if tc.expected != got {
				t.Errorf("%s", tc.description)
			}

		})
	}
}
