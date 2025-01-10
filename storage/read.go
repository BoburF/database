package storage

import "os"

func Read(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}
