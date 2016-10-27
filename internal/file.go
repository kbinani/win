package internal

import (
	"os"
)

func FileSize(filePath string) int64 {
	s, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return s.Size()
}

func IsFileExist(filePath string) bool {
	s, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !s.IsDir()
}
