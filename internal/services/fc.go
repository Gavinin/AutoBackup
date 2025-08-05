package services

import (
	"AutoBuckup/internal/adapter/folder_collector"
	"AutoBuckup/internal/config"
)

func SelectFolderCollector(cfg *config.Config) folder_collector.IFolderCollector {
	if cfg.Docker {
		return folder_collector.NewDockerFolderCollector(cfg.HideFolder)
	}
	return folder_collector.NewNativeFolderCollector(cfg.HideFolder)
}
