package commands

import (
	"fmt"
	"os"
	"path"
	"time"
)

func GenerateTimestampID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()) // Generates a timestamp-based ID
}

func GeneratePath(collection, name string) string {
    currentDir, err := os.Getwd()
    if err != nil{
        return fmt.Sprintf("%v", err)
    }
	return path.Join(currentDir, "data", collection, name)
}
