package utils

import (
	"path"
	"path/filepath"
	"strings"
)

func GetRootPath(module string) string {
	root, _ := filepath.Abs("")

	return path.Join(root[:strings.Index(root, module)], module)
}
