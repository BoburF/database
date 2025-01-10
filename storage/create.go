package storage

import (
	"os"
	"path/filepath"
)

func Create(path string, data string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

func CreateLink(path, linkPath string) error {
	err := os.Symlink(path, linkPath)
	if err != nil {
		return err
	}
	return nil
}
