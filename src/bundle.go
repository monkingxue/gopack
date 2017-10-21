package src

import (
	"io/ioutil"
	"github.com/gopack/util"
)

type Bundle struct {
	entryPath string
	destPath  string
}

var LoadModules = make(map[string]*Module)
var loadPaths []string

func (b *Bundle) checkCycle() {
	if HasCycle(LoadModules[b.entryPath]) {
		panic("存在循环调用！")
	}
}

func (b *Bundle) fetchModule() {
	loadPaths = append(loadPaths, b.entryPath)

	var i = 0
	var pathCnt = len(loadPaths)

	for ; i != pathCnt; i += 1 {

		source, err := ioutil.ReadFile(loadPaths[i])
		if err != nil {
			panic(err)
		}

		var imports []string
		var newModule = new(Module)
		*(newModule), imports = CreateModule(loadPaths[i], string(source))
		LoadModules [loadPaths[i]] = newModule

		for _, name := range imports {
			loadPaths = append(loadPaths, util.ResolvePath(b.entryPath, name))
		}
		pathCnt = len(loadPaths)
	}
}

func (b *Bundle) deconflict() {
	//得到一个bundle和一堆预处理好的modules，解决重复、冲突与循环依赖
}

func (b *Bundle) generate() string {
	var out = ""

	for _, m := range LoadModules {
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
	//b.checkCycle()
	b.deconflict()
	return b.generate()
}
