package util

import (
	"AutoBuckup/internal/enum"
	"fmt"
	"strings"
	"time"
)

func SeparatePath(s string) (string, string) {
	if s[len(s)-1] == '/' {
		s = s[:len(s)-1]
	}
	index := strings.LastIndex(s, "/")
	if index == -1 {
		return "", s
	}

	return s[:index], s[index+1:]
}

func GetExt(typ string) string {
	switch strings.ToLower(typ) {
	case enum.TypeTarGz:
		return ".tar.gz"
	case enum.TypeZipStore:
		return ".zip"
	default:
		return ".tar.gz"
	}
}

func GetFolderName(nameFormat, folderName string) string {
	timeFormatStr := time.Now().Format(NameFormat2DateFormat(nameFormat))
	folderName = fmt.Sprintf("%s%s%s", folderName, enum.FolderSeparateFlag, timeFormatStr)
	return folderName
}
