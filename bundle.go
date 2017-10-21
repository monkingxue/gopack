package main

import (
	"io/ioutil"
)

type Bundle struct {
	entryPath string
	destPath  string
}

var LoadModules map[string]*Module = make(map[string]*Module)

func (b *Bundle) isCycle() {
	if(LoadModules[b.entryPath].Dfs()) {
		panic("存在循环调用！")
	}
}

func (b *Bundle) fetchModule() {
	source, err := ioutil.ReadFile(b.entryPath)
	if err != nil {
		panic(err)
	}
	CreateModule(b.entryPath, string(source))
}

func (b *Bundle) deconflict() {
	//得到一个bundle和一堆预处理好的modules，解决重复、冲突与循环依赖
}

func (b *Bundle) generate() {
	//生成代码
}

func (b *Bundle) build() {
	b.fetchModule()
	b.deconflict()
	b.generate()
}
