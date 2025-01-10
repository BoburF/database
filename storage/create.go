package storage

import (
	"os"
)

func Create(path string, data string) error {
	err := os.WriteFile(path, []byte(data), 0644)
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
