package frame

import (
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
	"log"
)

type (
	ModuleTable map[int]string

	Module struct {
		bundle   *Bundle
		path     string
		code     string
		comments []string
		ast      *ast.Program
		imports  *ModuleTable
		exports  *ModuleTable
	}
)

func (self *Module) Constructor(bundle *Bundle, path string, code string) {
	self.bundle, self.path, self.code = bundle, path, code
	self.imports = new(ModuleTable)
	self.exports = new(ModuleTable)
	ast, err := parser.ParseFile(nil, "", code, 0)
	if err != nil {
		log.Fatal("Error parse!")
	}
	self.ast = ast

	self.analyse()
}

func (self *Module) analyse() {

}
