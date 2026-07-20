package allfxevents

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/types/typeutil"
)

var Analyzer = &analysis.Analyzer{
	Name: "allfxevents",
	Doc:  "check for unhandled fxevent.Events",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var _filter = []ast.Node{
	&ast.File{},
	&ast.FuncDecl{},
	&ast.CaseClause{},
	&ast.TypeAssertExpr{},
}

func run(pass *analysis.Pass) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

type visitor struct {
	Fset    *token.FileSet
	Info    *types.Info
	Fxevent fxevent
	Report  func(analysis.Diagnostic)

	loggerType types.Type
	funcEvents *typeSet
}

func (v *visitor) Visit(n ast.Node, push bool) (recurse bool) {
	_ = "STUB: not implemented"
	return false
}

func (v *visitor) funcDeclEnter(n *ast.FuncDecl) bool { _ = "STUB: not implemented"; return false }

func (v *visitor) funcDeclExit(n *ast.FuncDecl) bool { _ = "STUB: not implemented"; return false }

func findPackage(pkg *types.Package, importPath string) (_ *types.Package, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type fxevent struct {
	Logger          types.Type
	LoggerInterface *types.Interface

	Event  types.Type
	Events typeSet
}

func inspectFxevent(pkg *types.Package) fxevent { _ = "STUB: not implemented"; return *new(fxevent) }

type typeSet struct{ m typeutil.Map }

func (ts *typeSet) Len() int { _ = "STUB: not implemented"; return 0 }

func (ts *typeSet) Put(t types.Type) { _ = "STUB: not implemented"; return }

func (ts *typeSet) Remove(t types.Type) (found bool) { _ = "STUB: not implemented"; return false }

func (ts *typeSet) Iterate(f func(types.Type)) { _ = "STUB: not implemented"; return }

func (ts *typeSet) Clone() *typeSet { _ = "STUB: not implemented"; return nil }

func emptyQualifier(*types.Package) string { _ = "STUB: not implemented"; return "" }
