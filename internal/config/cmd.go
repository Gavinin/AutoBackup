package config

import (
	"AutoBuckup/internal/log"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func CmdParam() {

	hasInit := false
	for _, arg := range os.Args[1:] {
		if arg == "--init" {
			hasInit = true
			break
		}
	}
	if hasInit {
		defer os.Exit(0)

		filePath := filepath.Join(dirPath, confFileName)

		// Create dir
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			log.Logger.Fatalf("Create directory fail: %v\n", err)
			os.Exit(1)
		}

		// Create conf if not exist
		cfg := Config{
			AppName:   "MyApp",
			Directory: []string{},
			Cron:      "0 */2 * * *",
			Remote: Remote{
				Protocol: "sftp",
			},
			Archive: Archive{
				Type:         "tar.gz",
				SavePrevious: true,
				NameFormat:   "%Y%m%D%H%M",
				SortByDate:   true,
			},
		}
		data, err := yaml.Marshal(&cfg)
		if err != nil {
			log.Logger.Errorf("序列化YAML失败: %v", err)
		}

		filePathWithFile := fmt.Sprintf("%s/%s", dirPath, confFileName)

		err = os.WriteFile(filePathWithFile, data, 0644)
		if err != nil {
			if os.IsExist(err) {
				log.Logger.Warn("Config file has exist:", filePath)
			} else {
				log.Logger.Fatalf("Create config file fail: %v\n", err)
			}
		}
	}
}
