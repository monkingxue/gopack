package optimize

import (
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/file"

	cp "github.com/monkingxue/gopack/src/composite"
)

type (
	OptimizeStatement struct {
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
		TopStatement OptimizeStatement
	}

	AnalyzeWalker struct {
		source string
		shift  file.Idx
	}
)

func (*OptimizeStatement) _statementNode()     {}
func (self *OptimizeStatement) Idx0() file.Idx { return self.Node.Node.Idx0() }
func (self *OptimizeStatement) Idx1() file.Idx { return self.Node.Node.Idx1() }

var gal = Analyze{Scope: *new(Scope)}

func addToScope(name string, hoist bool) {
	gal.Scope.Add(name, hoist)

	if gal.Scope.Parent == nil {
		gal.TopStatement.Defines[name] = true
	}
}

func analyze(program *ast.Program, code string, module *cp.Module) {
	for i, statement := range program.Body {
		opStatement := OptimizeStatement{
			Node:   OpNode{Node: statement},
			Module: module,
			Margin: [2]file.Idx{0, 0},
			Source: code,
		}
	}

	w := &AnalyzeWalker{source: code}

	ast.Walk(w, opList)
}

func (aw *AnalyzeWalker) Enter(n ast.Node) ast.Visitor {

	switch n := n.(type) {
	case *ast.FunctionLiteral:
		addToScope(n.Name.Name, false)
	}
}

func (aw *AnalyzeWalker) Exit(n ast.Node) {

	switch n := n.(type) {
	case *ast.FunctionLiteral:
		addToScope(n.Name.Name, false)
	}
}
