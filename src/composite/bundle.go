package composite

import (
	"io/ioutil"
	"github.com/monkingxue/gopack/src/util"
	"path"
	"os"
)

type Bundle struct {
	entryPath string
	destPath  string
	config    util.Config
	out       string
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
	b.checkCycle()
}

func (b *Bundle) generate() string {

	for _, m := range loadModules {
		b.out += m.Code + "\n"
	}

	const intro = "(function () { 'use strict';\n\n\t"
	const outro = "\n\n})();"

	return intro + b.out + outro
}

func (b *Bundle) postProc() {
	if (b.config.Cleanup) {
		os.Remove(b.config.Dest)
	}

	os.Mkdir(b.config.Dest, 0777)

	out, err := os.OpenFile(path.Join(b.config.Dest, b.config.Chunk+".js"), os.O_RDWR|os.O_CREATE, 0777);
	if err != nil {
		panic(err)
	}
	out.WriteString(b.out)
	out.Close()
	if (b.config.Log) {
		println(b.out)
	}
}

func (b *Bundle) Build(p string) {
	b.config.GetConfig(path.Join(p, "gpconfig.json"))
	b.entryPath = b.config.Entry
	b.destPath = b.config.Dest
	b.fetchModule()
	b.deconflict()
	b.generate()
	b.postProc()
}
