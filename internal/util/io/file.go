package io

import (
	"os"
	"path/filepath"
)

func GetSubDirectories(path string) (subDirectories []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfos, err := file.Readdir(-1)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			subDirectories = append(subDirectories, filepath.Join(path, fileInfo.Name()))
		}
	}
	return subDirectories, nil
}

func GetFileSize(filePath string) (size int64, err error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {

		return 0, err
	}
	return fileInfo.Size(), nil
}
