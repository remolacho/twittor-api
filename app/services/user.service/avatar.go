package user_service

import (
	"errors"
	"twittor-api/domain/models/upload"
	"twittor-api/domain/models/user"
)

type AvatarService struct {
	RepositoryUser user.IUser
	MetaData       *upload.MetaDataFile
}

func NewAvatar(repository user.IUser, metaData *upload.MetaDataFile) *AvatarService {
	return &AvatarService{
		repository,
		metaData,
	}
}

func (s *AvatarService) Upload(userID string) error {
	currentUser, err := s.RepositoryUser.Find(userID)

	if err != nil {
		return errors.New("user not found for upload avatar")
	}

	currentUser.Avatar = s.MetaData.Name

	_, err = s.RepositoryUser.Update(currentUser)
	return err
}
