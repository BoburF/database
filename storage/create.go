package storage

import (
	"log"
	"os"
)

func Create(path string, data string) error {
    log.Println("Creating...", path, string(data))
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
