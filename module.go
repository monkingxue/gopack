package main

import (
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
	"github.com/robertkrimen/otto/file"
	"log"
)

type (
	Module struct {
		path    string
		code    string
		ast     *ast.Program
		imports []string
		visited bool
		added   bool
	}

	moduleWalker struct {
		imports []string
		source  string
		shift   file.Idx
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

func CreateModule(path string, code string) Module {
	program, err := parser.ParseFile(nil, "", code, 0)
	if err != nil {
		log.Fatal("Error parse code")
		return Module{}
	}

	w := &moduleWalker{source: code}

	ast.Walk(w, program)

	return Module{path: path, code: code, ast: program, imports: w.imports}
}

func (m *moduleWalker) Enter(n ast.Node) ast.Visitor {

	if cc, ok := n.(*ast.CallExpression); ok && cc != nil {
		if cc.Callee.(*ast.Identifier).Name == "require" {
			m.imports = append(m.imports, cc.ArgumentList[0].(*ast.StringLiteral).Value)
		}
	}

	return m
}

func (m *moduleWalker) Exit(n ast.Node) {
	//
}
