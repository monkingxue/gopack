package util

import "path"

func ResolvePath(entryPath string, name string) string {
	return path.Dir(entryPath) + "/" + name + ".js"
}
