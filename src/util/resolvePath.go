package util

import (
	"path"
	"strings"
)

func ResolvePath(entryPath string, importPath string) string {
	if (path.IsAbs(importPath)) {
		return importPath
	} else if (strings.HasPrefix(importPath, ".")) {
		return path.Join(path.Dir(entryPath), importPath) + ".js"
	} else {
		return path.Join(path.Dir(entryPath), "node_modules", importPath) + ".js"
	}
}
