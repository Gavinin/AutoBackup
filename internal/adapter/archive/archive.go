package archive

import (
	"AutoBuckup/internal/config"
)

type IArchiveAdapter interface {
	BatchArchive(archive config.Archive, paths []string) ([]string, error)
}
