package main

import (
	"github.com/robertkrimen/otto/ast"
)

type (
	Module struct {
		path    string
		code    string
		ast     *ast.Program
		imports *map[int]string
		exports *map[int]string
	}
)
