package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func CheckRelation(rel string) bool {
	fmt.Println("checking:", rel)
	return true
}

func EvalExpr(expr *ast.Expr) bool {

	switch (*expr).(type) {
	case *ast.ParenExpr:
		return EvalParenExpr(expr)
	case *ast.SelectorExpr:
		return EvalSelectorExpr(expr)
	case *ast.Ident:
		return EvalIdent(expr)
	case *ast.UnaryExpr:
		return EvalUnaryExpr(expr)
	case *ast.BinaryExpr:
		return EvalBinaryExpr(expr)
	}
	return false
}

func EvalParenExpr(expr *ast.Expr) bool {
	switch n := (*expr).(type) {
	case *ast.ParenExpr:
		return EvalExpr(&n.X)
	}
	return false
}
func EvalSelectorExpr(expr *ast.Expr) bool {
	switch n := (*expr).(type) {
	case *ast.SelectorExpr:
		if ident, ok := n.X.(*ast.Ident); ok {
			return CheckRelation(ident.Name + "." + n.Sel.Name)
		}
	}
	return false
}

func EvalIdent(expr *ast.Expr) bool {
	switch n := (*expr).(type) {
	case *ast.Ident:
		return CheckRelation(n.Name)
	}
	return false
}

func EvalUnaryExpr(expr *ast.Expr) bool {
	switch n := (*expr).(type) {
	case *ast.UnaryExpr:
		return !EvalExpr(&n.X)
	}
	return false
}

// TODO: use values instead?
func EvalBinaryExpr(expr *ast.Expr) bool {
	switch n := (*expr).(type) {
	case *ast.BinaryExpr:
		if n.Op.String() == "|" {
			return EvalExpr(&n.X) || EvalExpr(&n.Y)
		} else if n.Op.String() == "&" {
			return EvalExpr(&n.X) && EvalExpr(&n.Y)
		}

	}
	return false
}

func main() {
	fs := token.NewFileSet()
	// tr, _ := parser.ParseExpr("reader & !writer")
	// ast.Print(fs, tr)

	tr, _ := parser.ParseExpr("reader | writer | !parent_folder.reader")
	ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("!parent_folder.reader | reader | writer")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("(reader | writer | tester) & !parent_folder.reader")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("reader & writer")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("reader")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("!reader")
	// ast.Print(fs, tr)

	hasPerm := EvalExpr(&tr)
	fmt.Println(hasPerm)

}
