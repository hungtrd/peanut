package filemanager

import (
	"peanut/pkg/ary"
	"strings"
)

func CheckExtensionAvailable(ext string, listExt []string) bool {
	ext = strings.ToLower(ext)
	return ary.InArray(ext, listExt)
}
