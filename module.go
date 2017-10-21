package main

import (
	"github.com/robertkrimen/otto/ast"
)

type (
	Module struct {
		path    string
		code    string
		ast     *ast.Program
		imports *map[string]*Module
		exports *map[string]*Module
	}
)
