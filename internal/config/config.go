package config

import (
	"os"
	"strings"
)

const DefaultPath = "./config"
const DefaultConfFileName = "config.yaml"

var (
	dirPath      = DefaultPath
	confFileName = DefaultConfFileName
)

func Init() {
	if os.Getenv("AUTO_BACKUP_PATH") != "" {
		dirPath = os.Getenv("AUTO_BACKUP_PATH")
	}
	if strings.HasSuffix(dirPath, "/") {
		dirPath = dirPath[:len(dirPath)-1]
	}
	if os.Getenv("AUTO_BACKUP_FILE_NAME") != "" {
		confFileName = os.Getenv("AUTO_BACKUP_FILE_NAME")
	}
	CmdParam()

}
