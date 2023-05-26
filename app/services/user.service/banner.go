package user_service

import (
	"errors"
	"os"
	"twittor-api/domain/models/upload"
	"twittor-api/domain/models/user"
)

type BannerService struct {
	RepositoryUser user.IUser
	MetaData       *upload.MetaDataFile
}

type ResponseBanner struct {
	Banner    string `json:"banner"`
	Alternate string `json:"alternate"`
	HasBanner bool   `json:"hasBanner"`
}

func NewBanner(repository user.IUser, metaData ...*upload.MetaDataFile) *BannerService {
	var meta *upload.MetaDataFile

	if metaData != nil {
		meta = metaData[0]
	}

	return &BannerService{
		repository,
		meta,
	}
}

func (s *BannerService) Upload(userID string) (bool, error) {
	currentUser, err := s.RepositoryUser.Find(userID)

	if err != nil {
		return false, errors.New("user not found for upload banner")
	}

	currentUser.Banner = s.MetaData.Name
	_, err = s.RepositoryUser.Update(currentUser)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *BannerService) Get(userID string) (ResponseBanner, error) {
	currentUser, err := s.RepositoryUser.Find(userID)

	var response = ResponseBanner{
		Alternate: os.Getenv("BANNER_ALT"),
		HasBanner: false,
	}

	if err != nil {
		return response, errors.New("user not found")
	}

	response.Banner = currentUser.Banner
	response.HasBanner = true

	return response, nil
}
