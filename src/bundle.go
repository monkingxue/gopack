package src

import (
	"io/ioutil"
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

func (b *Bundle) FetchModule() {
	source, err := ioutil.ReadFile(b.entryPath)
	if err != nil {
		panic(err)
	}
	CreateModule(b.entryPath, string(source))
}

func (b *Bundle) FetchModule2() {
	loadPaths = append(loadPaths, b.entryPath)
	var i = 0
	var pathCnt = len(loadPaths)

	for ; i != pathCnt; i += 1 {
		if module, exists := LoadModules[loadPaths[i]]; !exists {
			source, err := ioutil.ReadFile(loadPaths[i])
			if err != nil {
				panic(err)
			}
			LoadModules[loadPaths[i]] = CreateModule(loadPaths[i], string(source))
		} else {
			for path, _ := range module.imports {
				//todo 如果有筛选path
				loadPaths = append(loadPaths, path)
			}
		}
		pathCnt = len(loadPaths)
	}

}

func (b *Bundle) deconflict() {
	//得到一个bundle和一堆预处理好的modules，解决重复、冲突与循环依赖
}

func (b *Bundle) generate() string {
	var out string

	//todo 生成代码

	const intro = `(function () { 'use strict';\n\n\t`
	const outro = `\n\n})();`

	return intro + out + outro
}

func (b *Bundle) Build() {
	b.FetchModule()

	b.checkCycle()
	b.deconflict()
	b.generate()
}
