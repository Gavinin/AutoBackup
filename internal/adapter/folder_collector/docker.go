package folder_collector

import (
	"AutoBuckup/internal/log"
	"AutoBuckup/internal/util"
	"fmt"
	"os"
	"strings"
)

type DockerFolderCollector struct {
	HideFolder bool
}

func NewDockerFolderCollector(hideFolder bool) *DockerFolderCollector {
	return &DockerFolderCollector{
		HideFolder: hideFolder,
	}
}

func (d *DockerFolderCollector) GetFolderList(paths []string) ([]string, error) {
	result := make([]string, 0)
	for _, path := range paths {
		if path == "" {
			continue
		}
		if !util.PathExists(path) {
			log.Logger.Debug("Path does not exist: %s", path)
			continue
		}
		entries, err := os.ReadDir(path)
		if err != nil {
			log.Logger.Error("Error reading directory: %s", path)
			continue
		}
		for _, entry := range entries {
			if !d.HideFolder {
				if strings.HasPrefix(entry.Name(), ".") {
					continue
				}
			}

			if entry.IsDir() {
				result = append(result, fmt.Sprintf("%s/%s", path, entry.Name()))
			}
		}

	}
	return result, nil
}
