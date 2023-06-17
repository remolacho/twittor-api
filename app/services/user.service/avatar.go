package user_service

import (
	"errors"
	"twittor-api/domain/models/upload"
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/shared"
)

type AvatarService struct {
	RepositoryUser user.IUser
	MetaData       *upload.MetaDataFile
}

func NewAvatar(repository user.IUser, metaData ...*upload.MetaDataFile) *AvatarService {
	var meta *upload.MetaDataFile

	if metaData != nil {
		meta = metaData[0]
	}

	return &AvatarService{
		repository,
		meta,
	}
}

func (s *AvatarService) Upload(userID string) (bool, error) {
	currentUser, err := s.RepositoryUser.Find(userID)

	if err != nil {
		return false, errors.New("user not found for upload avatar")
	}

	currentUser.Avatar = s.MetaData.Name
	_, err = s.RepositoryUser.Update(currentUser)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *AvatarService) Get(userID string) (string, bool) {
	rootPath := shared.GetFilesRootPath()
	currentUser, err := s.RepositoryUser.Find(userID)

	if err != nil {
		return "", false
	}

	if len(currentUser.Avatar) == 0 {
		return rootPath + "/avatar.jpeg", true
	}

	return rootPath + "/avatars/" + currentUser.Avatar, true
}
