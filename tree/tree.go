package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func CheckRelation(rel string) bool {
	fmt.Println("checking:", rel)
	return true
}

func EvalExpr(ctx context.Context, done chan bool, expr *ast.Expr) bool {

	returned := make(chan bool)

	go func() {
		switch (*expr).(type) {
		case *ast.ParenExpr:
			returned <- EvalParenExpr(ctx, done, expr)
		case *ast.SelectorExpr:
			returned <- EvalSelectorExpr(ctx, done, expr)
		case *ast.Ident:
			returned <- EvalIdent(ctx, done, expr)
		case *ast.UnaryExpr:
			returned <- EvalUnaryExpr(ctx, done, expr)
		case *ast.BinaryExpr:
			returned <- EvalBinaryExpr(ctx, done, expr)
		}
	}()

	select {
	case <-ctx.Done():
		return false
	case returnValue := <-returned:
		return returnValue
	}

}

func EvalParenExpr(ctx context.Context, done chan bool, expr *ast.Expr) bool {
	returned := make(chan bool)
	go func() {
		switch n := (*expr).(type) {
		case *ast.ParenExpr:
			returned <- EvalExpr(ctx, done, &n.X)
		}
	}()
	select {
	case <-ctx.Done():
		return false
	case returnValue := <-returned:
		return returnValue

	}
}
func EvalSelectorExpr(ctx context.Context, done chan bool, expr *ast.Expr) bool {
	returned := make(chan bool)
	go func() {
		switch n := (*expr).(type) {
		case *ast.SelectorExpr:
			if ident, ok := n.X.(*ast.Ident); ok {
				returned <- CheckRelation("!" + ident.Name + "." + n.Sel.Name)
			}
		}
	}()

	select {
	case <-ctx.Done():
		return false
	case returnValue := <-returned:
		return returnValue

	}
}

func EvalIdent(ctx context.Context, done chan bool, expr *ast.Expr) bool {
	returned := make(chan bool)

	go func() {
		switch n := (*expr).(type) {
		case *ast.Ident:
			returned <- CheckRelation(n.Name)
		}
	}()

	select {
	case <-ctx.Done():
		return false
	case returnValue := <-returned:
		return returnValue

	}

}

func EvalUnaryExpr(ctx context.Context, done chan bool, expr *ast.Expr) bool {
	returned := make(chan bool)

	go func() {
		switch n := (*expr).(type) {
		case *ast.UnaryExpr:
			returned <- !EvalExpr(ctx, done, &n.X)
		}
	}()
	select {
	case <-ctx.Done():
		return false
	case returnValue := <-returned:
		return returnValue
	}
}

// TODO: use values instead?
func EvalBinaryExpr(ctx context.Context, done chan bool, expr *ast.Expr) bool {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	switch n := (*expr).(type) {
	case *ast.BinaryExpr:
		go func() {
			done <- EvalExpr(ctx, done, &n.X)
		}()
		go func() {
			done <- EvalExpr(ctx, done, &n.Y)
		}()

		if n.Op.String() == "|" {
			for i := 0; i < 2; i++ {
				select {
				case isDone := <-done:
					if isDone {
						cancel()
						return true
					}
				case <-ctx.Done():
				}
			}

		} else if n.Op.String() == "&" {
			for i := 0; i < 2; i++ {

				select {
				case isDone := <-done:
					if !isDone {
						cancel()
						return false
					}
				case <-ctx.Done():
				}

			}
			return true
		}

	}
	return false
}

func StartEval(expr *ast.Expr) bool {
	ctx := context.Background()
	done := make(chan bool, 2)

	returned := make(chan bool)

	go func() {
		returned <- EvalExpr(ctx, done, expr)
	}()

	select {
	case <-ctx.Done():
		return false
	case returnValue := <-returned:
		return returnValue
	}

}

func main() {
	fs := token.NewFileSet()
	// tr, _ := parser.ParseExpr("reader & !writer")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("reader | writer | !parent_folder.reader")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("!parent_folder.reader | reader | writer")
	// ast.Print(fs, tr)

	tr, _ := parser.ParseExpr("(reader | writer | tester) & !parent_folder.reader")
	ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("reader & writer")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("reader")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("!reader")
	// ast.Print(fs, tr)

	hasPerm := StartEval(&tr)
	fmt.Println(hasPerm)

}
