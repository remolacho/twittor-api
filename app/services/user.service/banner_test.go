package user_service

import (
	"testing"
	"twittor-api/domain/models/upload"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	stubFactoryUser "twittor-api/infraestructure/stubs/factories/factory.users"
)

var stubBanner = stubFactoryUser.Build()

var testCasesBannerUser = []struct {
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
	service := NewAvatar(mockFactory, &upload.MetaDataFile{
		Path:      "/upload/banners/test.png",
		Name:      "test.png",
		Extension: ".png",
		Size:      5120,
	})

	for _, tc := range testCasesBannerUser {
		t.Run(tc.name, func(t *testing.T) {
			got, err := service.Upload(tc.userID)

			if got != tc.expected {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
