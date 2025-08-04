package services

import (
	"AutoBuckupG/internal/adapter/archive"
	"AutoBuckupG/internal/config"
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
