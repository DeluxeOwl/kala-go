package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Walk(node ast.Node, f func(ast.Node) bool) {

	queue := []ast.Node{node}

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		// dont process children if signal is given
		stop := f(current)

		if !stop {
			switch n := current.(type) {
			// Comments and fields
			case *ast.Comment:
				// nothing to do

			case *ast.CommentGroup:
				for _, c := range n.List {
					queue = append(queue, c)
				}

			case *ast.Field:
				if n.Doc != nil {
					queue = append(queue, n.Doc)
				}
				for _, x := range n.Names {
					queue = append(queue, x)
				}
				if n.Type != nil {
					queue = append(queue, n.Type)
				}
				if n.Tag != nil {
					queue = append(queue, n.Tag)
				}
				if n.Comment != nil {
					queue = append(queue, n.Comment)
				}

			case *ast.FieldList:
				for _, f := range n.List {
					queue = append(queue, f)
				}

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

			case *ast.CompositeLit:
				if n.Type != nil {
					queue = append(queue, n.Type)
				}
				for _, x := range n.Elts {
					queue = append(queue, x)
				}

			case *ast.ParenExpr:
				queue = append(queue, n.X)

			case *ast.SelectorExpr:
				queue = append(queue, n.X)
				queue = append(queue, n.Sel)

			case *ast.IndexExpr:
				queue = append(queue, n.X)
				queue = append(queue, n.Index)

			case *ast.IndexListExpr:
				queue = append(queue, n.X)
				for _, index := range n.Indices {
					queue = append(queue, index)
				}

			case *ast.SliceExpr:
				queue = append(queue, n.X)
				if n.Low != nil {
					queue = append(queue, n.Low)
				}
				if n.High != nil {
					queue = append(queue, n.High)
				}
				if n.Max != nil {
					queue = append(queue, n.Max)
				}

			case *ast.TypeAssertExpr:
				queue = append(queue, n.X)
				if n.Type != nil {
					queue = append(queue, n.Type)
				}

			case *ast.CallExpr:
				queue = append(queue, n.Fun)
				for _, x := range n.Args {
					queue = append(queue, x)
				}

			case *ast.StarExpr:
				queue = append(queue, n.X)

			case *ast.UnaryExpr:
				queue = append(queue, n.X)

			case *ast.BinaryExpr:
				queue = append(queue, n.X)
				queue = append(queue, n.Y)

			case *ast.KeyValueExpr:
				queue = append(queue, n.Key)
				queue = append(queue, n.Value)

			// Types
			case *ast.ArrayType:
				if n.Len != nil {
					queue = append(queue, n.Len)
				}
				queue = append(queue, n.Elt)

			case *ast.StructType:
				queue = append(queue, n.Fields)

			case *ast.FuncType:
				if n.TypeParams != nil {
					queue = append(queue, n.TypeParams)
				}
				if n.Params != nil {
					queue = append(queue, n.Params)
				}
				if n.Results != nil {
					queue = append(queue, n.Results)
				}

			case *ast.InterfaceType:
				queue = append(queue, n.Methods)

			case *ast.MapType:
				queue = append(queue, n.Key)
				queue = append(queue, n.Value)

			case *ast.ChanType:
				queue = append(queue, n.Value)

			// Statements
			case *ast.BadStmt:
				// nothing to do

			case *ast.DeclStmt:
				queue = append(queue, n.Decl)

			case *ast.EmptyStmt:
				// nothing to do

			case *ast.LabeledStmt:
				queue = append(queue, n.Label)
				queue = append(queue, n.Stmt)

			case *ast.ExprStmt:
				queue = append(queue, n.X)

			case *ast.SendStmt:
				queue = append(queue, n.Chan)
				queue = append(queue, n.Value)

			case *ast.IncDecStmt:
				queue = append(queue, n.X)

			case *ast.AssignStmt:
				for _, x := range n.Lhs {
					queue = append(queue, x)
				}
				for _, x := range n.Rhs {
					queue = append(queue, x)
				}

			case *ast.GoStmt:
				queue = append(queue, n.Call)

			case *ast.DeferStmt:
				queue = append(queue, n.Call)

			case *ast.ReturnStmt:
				for _, x := range n.Results {
					queue = append(queue, x)
				}

			case *ast.BranchStmt:
				if n.Label != nil {
					queue = append(queue, n.Label)
				}

			case *ast.BlockStmt:
				for _, x := range n.List {
					queue = append(queue, x)
				}

			case *ast.IfStmt:
				if n.Init != nil {
					queue = append(queue, n.Init)
				}
				queue = append(queue, n.Cond)
				queue = append(queue, n.Body)
				if n.Else != nil {
					queue = append(queue, n.Else)
				}

			case *ast.CaseClause:
				for _, x := range n.List {
					queue = append(queue, x)
				}
				for _, x := range n.Body {
					queue = append(queue, x)
				}

			case *ast.SwitchStmt:
				if n.Init != nil {
					queue = append(queue, n.Init)
				}
				if n.Tag != nil {
					queue = append(queue, n.Tag)
				}
				queue = append(queue, n.Body)

			case *ast.TypeSwitchStmt:
				if n.Init != nil {
					queue = append(queue, n.Init)
				}
				queue = append(queue, n.Assign)
				queue = append(queue, n.Body)

			case *ast.CommClause:
				if n.Comm != nil {
					queue = append(queue, n.Comm)
				}
				for _, x := range n.Body {
					queue = append(queue, x)
				}

			case *ast.SelectStmt:
				queue = append(queue, n.Body)

			case *ast.ForStmt:
				if n.Init != nil {
					queue = append(queue, n.Init)
				}
				if n.Cond != nil {
					queue = append(queue, n.Cond)
				}
				if n.Post != nil {
					queue = append(queue, n.Post)
				}
				queue = append(queue, n.Body)

			case *ast.RangeStmt:
				if n.Key != nil {
					queue = append(queue, n.Key)
				}
				if n.Value != nil {
					queue = append(queue, n.Value)
				}
				queue = append(queue, n.X)
				queue = append(queue, n.Body)

			// Declarations
			case *ast.ImportSpec:
				if n.Doc != nil {
					queue = append(queue, n.Doc)
				}
				if n.Name != nil {
					queue = append(queue, n.Name)
				}
				queue = append(queue, n.Path)
				if n.Comment != nil {
					queue = append(queue, n.Comment)
				}

			case *ast.ValueSpec:
				if n.Doc != nil {
					queue = append(queue, n.Doc)
				}
				for _, x := range n.Names {
					queue = append(queue, x)
				}
				if n.Type != nil {
					queue = append(queue, n.Type)
				}
				for _, x := range n.Values {
					queue = append(queue, x)
				}
				if n.Comment != nil {
					queue = append(queue, n.Comment)
				}

			case *ast.TypeSpec:
				if n.Doc != nil {
					queue = append(queue, n.Doc)
				}
				queue = append(queue, n.Name)
				if n.TypeParams != nil {
					queue = append(queue, n.TypeParams)
				}
				queue = append(queue, n.Type)
				if n.Comment != nil {
					queue = append(queue, n.Comment)
				}

			case *ast.BadDecl:
				// nothing to do

			case *ast.GenDecl:
				if n.Doc != nil {
					queue = append(queue, n.Doc)
				}
				for _, s := range n.Specs {
					queue = append(queue, s)
				}

			case *ast.FuncDecl:
				if n.Doc != nil {
					queue = append(queue, n.Doc)
				}
				if n.Recv != nil {
					queue = append(queue, n.Recv)
				}
				queue = append(queue, n.Name)
				queue = append(queue, n.Type)
				if n.Body != nil {
					queue = append(queue, n.Body)
				}

			// Files and packages
			case *ast.File:
				if n.Doc != nil {
					queue = append(queue, n.Doc)
				}
				queue = append(queue, n.Name)
				for _, x := range n.Decls {
					queue = append(queue, x)
				}
				// don't walk n.Comments - they have been
				// visited already through the individual
				// nodes

			case *ast.Package:
				for _, f := range n.Files {
					queue = append(queue, f)
				}

			default:
				panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
			}
		}

	}
}

func Inspect(node ast.Node, f func(ast.Node) bool) {
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

	Inspect(tr, func(n ast.Node) bool {
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
				fmt.Printf("negation: %s\n", s)
			}
			// switch nodeInner := node.X.(type) {
			// case *ast.Ident:
			// 	if s != "" {
			// 		fmt.Printf("check relation: %s\n", nodeInner.Name)
			// 	}
			// }

		// continue if selector
		case *ast.SelectorExpr:
			left := node.X.(*ast.Ident).Name
			right := node.Sel.Name
			fmt.Printf("composed relation: %s.%s\n", left, right)
			return true
		}
		return false
	})
}
