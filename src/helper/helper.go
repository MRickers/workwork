package helper

import (
	"strings"
)

func StripExeName(exePath string) string {
	path := strings.Replace(exePath, "\\", "/", -1)
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash == -1 {
		return ""
	}
	return path[:lastSlash+1]
}
