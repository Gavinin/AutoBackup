package util

import "strings"

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
