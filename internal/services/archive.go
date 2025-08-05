package services

import (
	"AutoBuckup/internal/adapter/archive"
	"AutoBuckup/internal/config"
	"AutoBuckup/internal/enum"
	"strings"
)

func SelectArchive(cfg *config.Config) archive.IArchiveAdapter {
	switch strings.ToLower(cfg.Archive.Type) {
	case enum.TypeTarGz:
		return archive.NewTagGz()
	case enum.TypeZipStore:
		return archive.NewZipStore()
	default:
		return archive.NewTagGz()

	}
}
