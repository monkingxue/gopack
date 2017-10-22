package optimize

import (
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/file"

	cp "github.com/monkingxue/gopack/src/composite"
)

type (
	OpBlock struct {
		Node      OpNode
		Defines   map[string]bool
		Modifies  map[string]bool
		DependsOn map[string]bool
		Included  bool
		Module    *cp.Module
		Source    string
		Margin    [2]file.Idx
	}

	OpNode struct {
		Node     ast.Node
		NewScope *Scope
	}

	Analyze struct {
		Scope        Scope
		TopStatement OpBlock
	}

	AnalyzeWalker struct {
		source  string
		shift   file.Idx
	}
)

var gal = Analyze{Scope: *new(Scope)}

func addToScope(name string, hoist bool) {
	gal.Scope.Add(name, hoist)

	if gal.Scope.Parent == nil {
		gal.TopStatement.Defines[name] = true
	}
}

func analyze(ast *ast.Program, code string, module *cp.Module) {
	var opList []OpBlock
	for _, statement := range ast.Body {
		opItem := OpBlock{
			Node:   OpNode{Node: statement},
			Module: module,
			Margin: [2]file.Idx{0, 0},
			Source: code,
		}

		opList = append(opList, opItem)
	}


}

func (aw *AnalyzeWalker)  Enter(n ast.Node) ast.Visitor {

	switch n := n.(type) {
	case *ast.FunctionLiteral :
		addToScope(n.Name.Name, false)
	}
}
