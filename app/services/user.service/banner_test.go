package user_service

import (
	"testing"
	"twittor-api/domain/models/upload"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	stubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

var stubBanner = stubFactoryUser.Build()

var testCasesUploadBannerUser = []struct {
	name        string
	userID      string
	expected    bool
	description string
}{
	{
		name:        "user id",
		userID:      "",
		expected:    false,
		description: "the user id is empty",
	},
	{
		name:        "upload success",
		userID:      stubBanner.User("created").ID.Hex(),
		expected:    true,
		description: "the banner was uploaded success",
	},
}

func TestUploadBanner(t *testing.T) {
	mockFactory := repositoryFactoryUser.Build("test")
	service := NewBanner(mockFactory, &upload.MetaDataFile{
		Path:      "/upload/banners/test.png",
		Name:      "test.png",
		Extension: ".png",
		Size:      5120,
	})

	for _, tc := range testCasesUploadBannerUser {
		t.Run(tc.name, func(t *testing.T) {
			got, err := service.Upload(tc.userID)

			if got != tc.expected {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}

var testCasesGetBanner = []struct {
	name        string
	userID      string
	expected    bool
	description string
}{
	{
		name:        "user id",
		userID:      "",
		expected:    false,
		description: "the user id is empty",
	},
	{
		name:        "user id",
		userID:      stubBanner.User("created").ID.Hex(),
		expected:    true,
		description: "the banner was obtained",
	},
}

func TestGetBanner(t *testing.T) {
	mockFactory := repositoryFactoryUser.Build("test")
	service := NewBanner(mockFactory)

	for _, tc := range testCasesGetBanner {
		t.Run(tc.name, func(t *testing.T) {
			got, _, err := service.Get(tc.userID)

			if got != tc.expected {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
