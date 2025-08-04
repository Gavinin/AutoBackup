package util

import (
	"strings"
)

func NameFormat2DateFormat(nameFormat string) string {
	nameFormat = strings.ReplaceAll(nameFormat, "%Y", "2006")
	nameFormat = strings.ReplaceAll(nameFormat, "%m", "01")
	nameFormat = strings.ReplaceAll(nameFormat, "%D", "02")
	nameFormat = strings.ReplaceAll(nameFormat, "%H", "15")
	nameFormat = strings.ReplaceAll(nameFormat, "%M", "04")
	nameFormat = strings.ReplaceAll(nameFormat, "%S", "05")
	return nameFormat
}
