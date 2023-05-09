package upload_file

import (
	"mime/multipart"
	"path/filepath"
	"twittor-api/domain/models/upload"
)

type Upload struct {
	File    multipart.File
	Handler *multipart.FileHeader
	Token   string
	Root    string
}

func New(file multipart.File, handler *multipart.FileHeader, token string) *Upload {
	return &Upload{
		file,
		handler,
		token,
		getRootPath(),
	}
}

func (s *Upload) Call(folder string) (*upload.MetaDataFile, error) {
	err := s.findOrCreateDir(folder)

	if err != nil {
		return nil, err
	}

	metadata, errData := s.createMetaData(folder)

	if errData != nil {
		return nil, errData
	}

	err = s.openCopyFile(metadata.Path)

	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func getRootPath() string {
	folderPath := "../twittor-api/uploads/"

	pathRoot, err := filepath.Abs(folderPath)

	if err != nil {
		return ""
	}

	return pathRoot
}
