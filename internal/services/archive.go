package services

import (
	"AutoBuckup/internal/adapter/archive"
	"AutoBuckup/internal/config"
	"strings"
)

func SelectArchive(cfg *config.Config) archive.IArchiveAdapter {
	switch strings.ToLower(cfg.Archive.Type) {
	case archive.TypeTarGz:
		return archive.NewTagGz()
	default:
		return archive.NewTagGz()

	}
}
