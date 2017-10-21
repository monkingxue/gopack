package main

import (
	"github.com/robertkrimen/otto/ast"
)

type (
	Module struct {
		path    string
		code    string
		ast     *ast.Program
		imports map[string]*Module
		exports map[string]*Module
		visited bool
		added   bool
	}
)

func (m *Module) Dfs() bool {
	if (m.visited) {
		return true
	}
	m.visited = true
	for _, v := range m.imports {
		if (v.Dfs()) {
			return true
		}
	}
	m.visited = false
	return false
}

func CreateModule(path string, source string) *Module {
	return &Module{}
}
