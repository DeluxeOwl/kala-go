package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func shouldStop(node ast.Node) bool {

	switch node.(type) {
	case *ast.SelectorExpr:
		return true
	}

	return false
}

func Walk(node ast.Node, f func(ast.Node)) {

	queue := []ast.Node{node}

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		// dont process children if signal is given
		f(current)

		if !shouldStop(current) {
			switch n := current.(type) {

			// Expressions
			case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
				// nothing to do

			case *ast.Ellipsis:
				if n.Elt != nil {
					queue = append(queue, n.Elt)
				}

			case *ast.FuncLit:
				queue = append(queue, n.Type)
				queue = append(queue, n.Body)

			case *ast.ParenExpr:
				queue = append(queue, n.X)

			case *ast.SelectorExpr:
				queue = append(queue, n.X)
				queue = append(queue, n.Sel)

			case *ast.CallExpr:
				queue = append(queue, n.Fun)
				for _, x := range n.Args {
					queue = append(queue, x)
				}

			case *ast.UnaryExpr:
				queue = append(queue, n.X)

			case *ast.BinaryExpr:
				queue = append(queue, n.X)
				queue = append(queue, n.Y)

			// Statements
			case *ast.BadStmt:
				// nothing to do

			case *ast.DeclStmt:
				queue = append(queue, n.Decl)

			case *ast.EmptyStmt:
				// nothing to do

			case *ast.ExprStmt:
				queue = append(queue, n.X)

			default:
				panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
			}
		}

	}
}

func Inspect(node ast.Node, f func(ast.Node)) {
	Walk(node, f)
}

func main() {
	fs := token.NewFileSet()
	tr, _ := parser.ParseExpr("reader & !writer")
	ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("reader | writer | !parent_folder.reader")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("!parent_folder.reader | reader | writer")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("(reader | writer) & !parent_folder.reader")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("reader & writer")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("reader")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("!reader")
	// ast.Print(fs, tr)

	Inspect(tr, func(n ast.Node) {
		var s string

		switch node := n.(type) {
		case *ast.BinaryExpr:
			s = node.Op.String()
			if s != "" {
				fmt.Printf("operation: %s\n", s)
			}
		case *ast.Ident:
			s = node.Name
			if s != "" {
				fmt.Printf("check relation: %s\n", s)
			}
		case *ast.UnaryExpr:
			s = node.Op.String()
			if s != "" {
				fmt.Printf("negation: %s \n", s)
			}

		// continue if selector
		case *ast.SelectorExpr:
			left := node.X.(*ast.Ident).Name
			right := node.Sel.Name
			fmt.Printf("composed relation: %s.%s\n", left, right)
		}
	})
}
