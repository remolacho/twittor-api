package upload_file

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"twittor-api/domain/models/upload"
)

func (s *Upload) createMetaData(folder string) (*upload.MetaDataFile, error) {
	extensionSplit := strings.Split(s.Handler.Filename, ".")
	extension := extensionSplit[1]

	if len(extension) == 0 {
		return nil, errors.New("the extension is empty")
	}

	var name = folder + "-" + s.Token + "." + extension
	var path = filepath.Join(s.Root, folder, name)

	return &upload.MetaDataFile{
		Path:      path,
		Name:      name,
		Extension: extension,
		Size:      s.Handler.Size,
	}, nil
}

func (s *Upload) findOrCreateDir(folder string) error {
	path := filepath.Join(s.Root, folder)

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Upload) openCopyFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	_, err = io.Copy(f, s.File)

	if err != nil {
		return err
	}

	return nil
}
