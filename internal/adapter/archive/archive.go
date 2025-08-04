package archive

import "AutoBuckup/internal/config"

const (
	TypeTarGz = "tar.gz"
)

type IArchiveAdapter interface {
	BatchArchive(archive config.Archive, paths []string) ([]string, error)
}
