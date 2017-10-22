package composite

import (
	"io/ioutil"
	"github.com/gopack/util"
)

type Bundle struct {
	entryPath string
	destPath  string
}

var loadModules = make(map[string]*Module)
var loadPaths []string

func (b *Bundle) checkCycle() {
	if HasCycle(loadModules[b.entryPath]) {
		panic("存在循环调用！")
	}
}

func (b *Bundle) fetchModule() {
	loadPaths = append(loadPaths, b.entryPath)

	i, pathCnt := 0, len(loadPaths)

	for ; i != pathCnt; i += 1 {

		source, err := ioutil.ReadFile(loadPaths[i])
		if err != nil {
			panic(err)
		}

		module, imports := CreateModule(loadPaths[i], string(source))
		loadModules [loadPaths[i]] = &module

		for _, path := range imports {
			loadPaths = append(loadPaths, util.ResolvePath(module.Path, path))
		}
		pathCnt = len(loadPaths)
	}
}

func (b *Bundle) deconflict() {
	//得到一个bundle和一堆预处理好的modules，解决重复、冲突与循环依赖
}

func (b *Bundle) generate() string {
	out := ""

	for _, m := range loadModules {
		out += m.Code + "\n"
	}

	const intro = "(function () { 'use strict';\n\n\t"
	const outro = "\n\n})();"

	return intro + out + outro
}

func (b *Bundle) Build(entryPath string, destPath string) string {
	b.entryPath = entryPath
	b.destPath = destPath
	b.fetchModule()
	b.checkCycle()
	b.deconflict()
	return b.generate()
}
