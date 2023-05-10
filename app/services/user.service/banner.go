package user_service

import (
	"errors"
	"twittor-api/domain/models/upload"
	"twittor-api/domain/models/user"
)

type BannerService struct {
	RepositoryUser user.IUser
	MetaData       *upload.MetaDataFile
}

type ResponseBanner struct {
	Banner string
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

func (s *BannerService) Get(userID string) (bool, ResponseBanner, error) {
	currentUser, err := s.RepositoryUser.Find(userID)

	if err != nil {
		return false, ResponseBanner{}, errors.New("user not found")
	}

	response := ResponseBanner{
		Banner: currentUser.Banner,
	}

	return true, response, nil
}
