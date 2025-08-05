package archive

import (
	"AutoBuckup/internal/config"
	"os"
)

type IArchiveAdapter interface {
	BatchArchive(archive config.Archive, paths []string) ([]string, error)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
