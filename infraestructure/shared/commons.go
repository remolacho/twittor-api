package shared

import "path/filepath"

func GetFilesRootPath() string {
	folderPath := "../twittor-api/uploads/"

	pathRoot, err := filepath.Abs(folderPath)

	if err != nil {
		return ""
	}

	return pathRoot
}
