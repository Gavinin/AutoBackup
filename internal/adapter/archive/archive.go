package archive

import "AutoBuckupG/internal/config"

const (
	TypeTarGz = "tar.gz"
)

type IArchiveAdapter interface {
	BatchArchive(archive config.Archive, paths []string) ([]string, error)
}
