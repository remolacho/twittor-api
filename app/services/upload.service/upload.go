package upload_service

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
	"twittor-api/domain/models/upload"
)

type MetaData struct {
	File    multipart.File
	Handler *multipart.FileHeader
	UserID  string
	Root    string
}

func New(file multipart.File, handler *multipart.FileHeader, userID string) *MetaData {
	return &MetaData{
		file,
		handler,
		userID,
		"../../uploads/",
	}
}

func (s *MetaData) Call(folder string) (*upload.MetaDataFile, error) {
	var ext = strings.Split(s.Handler.Filename, ".")[1]
	var name = folder + "-" + s.UserID + "." + ext
	var path = s.Root + folder + name

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(f, s.File)

	if err != nil {
		return nil, err
	}

	return &upload.MetaDataFile{
		Path:      path,
		Name:      name,
		Extension: ext,
	}, nil
}
