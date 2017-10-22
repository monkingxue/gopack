package composite

import (
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
	"github.com/robertkrimen/otto/file"
	"github.com/monkingxue/gopack/src/util"
)

type (
	Module struct {
		Path    string
		Code    string
		Ast     *ast.Program
		imports []string
		visited bool
	}

	moduleWalker struct {
		imports []string
		source  string
		shift   file.Idx
		delCol  [][]file.Idx
	}
)

func HasCycle(m *Module) bool {
	if m.visited {
		return true
	}
	m.visited = true
	for _, path := range m.imports {
		if HasCycle(loadModules[util.ResolvePath(m.Path, path)]) {
			return true
		}
	}
	m.visited = false
	return false
}

func CreateModule(path string, code string) (Module, []string) {
	if loadModules[path] != nil {
		return *loadModules[path], []string{}
	}

	program, err := parser.ParseFile(nil, "", code, 0)
	if err != nil {
		panic("Error parse code")
	}

	w := &moduleWalker{source: code}

	ast.Walk(w, program)

	for i := len(w.delCol) - 1; i >= 0; i-- {
		d := w.delCol[i]
		w.source = w.source[:d[0]] + w.source[d[1]:]
	}

	return Module{Path: path, Code: w.source, Ast: program, imports: w.imports}, w.imports
}

func (mw *moduleWalker) Enter(n ast.Node) ast.Visitor {
	var stateStart, stateEnd file.Idx
	switch ve := n.(type) {
	case *ast.VariableStatement:
		stateStart, stateEnd = n.Idx0()+mw.shift-1, n.Idx1()+mw.shift
		for _, vl := range ve.List {
			switch ve := vl.(type) {
			case *ast.VariableExpression:
				switch cc := ve.Initializer.(type) {
				case *ast.CallExpression:
					switch id := cc.Callee.(type) {
					case *ast.Identifier:
						if id.Name == "require" {
							path := cc.ArgumentList[0].(*ast.StringLiteral).Value
							mw.delCol = append(mw.delCol, []file.Idx{stateStart, stateEnd})
							if loadModules[path] == nil {
								mw.imports = append(mw.imports, path)
							}
						}
					}
				}
			}
		}

	}

	return mw
}

func (mw *moduleWalker) Exit(n ast.Node) {
	//
}
