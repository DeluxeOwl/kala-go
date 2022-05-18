package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func CheckRelation(ctx context.Context, done chan bool, rel string) bool {

	if ctx.Err() != nil {
		return false
	}

	fmt.Printf("checking rel: %s\n", rel)
	return true
}

func EvalExpr(ctx context.Context, done chan bool, expr *ast.Expr) bool {

	if ctx.Err() != nil {
		return false
	}

	switch (*expr).(type) {
	case *ast.ParenExpr:
		return EvalParenExpr(ctx, done, expr)
	case *ast.SelectorExpr:
		return EvalSelectorExpr(ctx, done, expr)
	case *ast.Ident:
		return EvalIdent(ctx, done, expr)
	case *ast.UnaryExpr:
		return EvalUnaryExpr(ctx, done, expr)
	case *ast.BinaryExpr:
		return EvalBinaryExpr(ctx, done, expr)
	}

	return false

}

func EvalParenExpr(ctx context.Context, done chan bool, expr *ast.Expr) bool {

	if ctx.Err() != nil {
		return false
	}

	switch n := (*expr).(type) {
	case *ast.ParenExpr:
		return EvalExpr(ctx, done, &n.X)
	}
	return false
}
func EvalSelectorExpr(ctx context.Context, done chan bool, expr *ast.Expr) bool {
	if ctx.Err() != nil {
		return false
	}
	switch n := (*expr).(type) {
	case *ast.SelectorExpr:
		if ident, ok := n.X.(*ast.Ident); ok {
			return CheckRelation(ctx, done, "!"+ident.Name+"."+n.Sel.Name)
		}
	}
	return false
}

func EvalIdent(ctx context.Context, done chan bool, expr *ast.Expr) bool {
	if ctx.Err() != nil {
		return false
	}

	switch n := (*expr).(type) {
	case *ast.Ident:
		return CheckRelation(ctx, done, n.Name)
	}

	return false
}

func EvalUnaryExpr(ctx context.Context, done chan bool, expr *ast.Expr) bool {

	if ctx.Err() != nil {
		return false
	}

	switch n := (*expr).(type) {
	case *ast.UnaryExpr:
		return !EvalExpr(ctx, done, &n.X)
	}

	return false
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
						fmt.Println("cancel")
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
						fmt.Println("cancel")
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
	case returnValue := <-returned:
		return returnValue
	case <-ctx.Done():
		return false
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

	// tr, _ := parser.ParseExpr("(reader | writer | tester) & !parent_folder.reader")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("reader & writer")
	// ast.Print(fs, tr)

	// tr, _ := parser.ParseExpr("reader")
	// ast.Print(fs, tr)

	tr, _ := parser.ParseExpr("!reader")
	ast.Print(fs, tr)

	hasPerm := StartEval(&tr)
	fmt.Println("has permission?", hasPerm)
	// time.Sleep(5 * time.Second)

}
