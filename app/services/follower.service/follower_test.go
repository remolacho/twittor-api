package follower_service

import (
	"testing"
	repositoryFactoryFollow "twittor-api/infraestructure/repositories/factories/repository.factory.follower"
	stubFactoryFollower "twittor-api/infraestructure/stubs/factories/factory.followers"
)

var stubFollowFind = stubFactoryFollower.Build()

type InputFind struct {
	FollowerID string
	UserID     string
}

var testCasesFindFollower = []struct {
	name        string
	input       InputFind
	expected    bool
	description string
}{
	{
		name:        "the user id empty",
		description: "It is not valid by the userID is empty",
		input: InputFind{
			UserID:     "",
			FollowerID: stubFollowFind.Follower("created").FollowUserID,
		},
		expected: false,
	},
	{
		name:        "the follower id empty",
		description: "It is not valid by the followerID is empty",
		input: InputFind{
			UserID:     stubFollowFind.Follower("created").FollowUserID,
			FollowerID: "",
		},
		expected: false,
	},
	{
		name:        "the follower not found userID",
		description: "It is not valid by the userID",
		input: InputFind{
			UserID:     stubFollowFind.Follower("created").FollowUserID,
			FollowerID: stubFollowFind.Follower("created").FollowUserID,
		},
		expected: false,
	},
	{
		name:        "the follower not found followerID",
		description: "It is not valid by the followerID",
		input: InputFind{
			UserID:     stubFollowFind.Follower("created").UserID,
			FollowerID: stubFollowFind.Follower("created").UserID,
		},
		expected: false,
	},
	{
		name:        "the follower success",
		description: "It success",
		input: InputFind{
			UserID:     stubFollowFind.Follower("created").UserID,
			FollowerID: stubFollowFind.Follower("created").FollowUserID,
		},
		expected: true,
	},
}

func TestFindFollower(t *testing.T) {
	mockFactoryRelation := repositoryFactoryFollow.Build("test")

	service := NewFollower(mockFactoryRelation)

	for _, tc := range testCasesFindFollower {
		t.Run(tc.name, func(t *testing.T) {
			response, err := service.Followed(tc.input.UserID, tc.input.FollowerID)
			if response.Followed != tc.expected {
				t.Errorf("%s: %v", tc.description, err)
			}
		})
	}
}
