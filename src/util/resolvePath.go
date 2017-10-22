package util

import (
	"path"
	"strings"
)

func ResolvePath(rootPath, entryPath, importPath string) string {
	if (path.IsAbs(importPath)) {
		return importPath
	} else if (strings.HasPrefix(importPath, ".")) {
		return path.Join(path.Dir(entryPath), importPath) + ".js"
	} else {
		return path.Join(path.Dir(rootPath), "node_modules", importPath) + ".js"
	}
}
