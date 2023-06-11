package user_service

import (
	"os"
	"twittor-api/domain/models/user"
)

type Banner struct {
	BannerUrl    string `json:"bannerUrl"`
	AlternateUrl string `json:"alternateUrl"`
}

type Avatar struct {
	AvatarUrl    string `json:"avatarUrl"`
	AlternateUrl string `json:"alternateUrl"`
}

type ResponseProfile struct {
	*user.User `json:"user"`
	Avatar     `json:"avatar"`
	Banner     `json:"banner"`
}

type UserProfileService struct {
	RepositoryUser user.IUser
}

func NewProfile(repository user.IUser) *UserProfileService {
	return &UserProfileService{
		repository,
	}
}

func (s *UserProfileService) Get(ID string) (*ResponseProfile, error) {
	u, err := s.RepositoryUser.Find(ID)

	if err != nil {
		return nil, err
	}

	u.Password = ""

	response := &ResponseProfile{
		User: u,
		Avatar: Avatar{
			AvatarUrl:    u.Avatar,
			AlternateUrl: os.Getenv("AVATAR_ALT"),
		},
		Banner: Banner{
			BannerUrl:    u.Banner,
			AlternateUrl: os.Getenv("BANNER_ALT"),
		},
	}

	u.Password = ""
	u.Banner = ""
	u.Avatar = ""

	return response, nil
}
