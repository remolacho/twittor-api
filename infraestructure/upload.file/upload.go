package upload_file

import (
	"mime/multipart"
	"twittor-api/domain/models/upload"
	"twittor-api/infraestructure/shared"
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
		shared.GetFilesRootPath(),
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
