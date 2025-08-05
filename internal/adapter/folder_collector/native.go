package folder_collector

import (
	"AutoBuckup/internal/log"
	"AutoBuckup/internal/util"
)

type NativeFolderCollector struct {
	HideFolder bool
}

func NewNativeFolderCollector(hideFolder bool) *NativeFolderCollector {
	return &NativeFolderCollector{
		HideFolder: hideFolder,
	}
}

func (n *NativeFolderCollector) GetFolderList(paths []string) ([]string, error) {
	result := make([]string, 0)
	for _, path := range paths {
		if path == "" {
			continue
		}
		if !util.PathExists(path) {
			log.Logger.Debug("Path does not exist: %s", path)
			continue
		}
		result = append(result, path)

	}
	return result, nil
}
