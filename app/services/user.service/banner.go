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

func NewBanner(repository user.IUser, metaData *upload.MetaDataFile) *AvatarService {
	return &AvatarService{
		repository,
		metaData,
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
