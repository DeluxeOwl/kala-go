package eval

import (
	"context"
	"errors"
	"go/ast"
	"go/parser"
	"time"
)

var ErrParsing = errors.New("error when parsing expression")

func CheckRelation(ctx context.Context, rel string) bool {

	if ctx.Err() != nil {
		return false
	}

	// fmt.Printf("checking rel: %s\n", rel)
	time.Sleep(time.Second)
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
			return CheckRelation(ctx, "!"+ident.Name+"."+n.Sel.Name)
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
		return CheckRelation(ctx, n.Name)
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

// EvalBinaryExpr runs both sides of the binary tree concurrently
// for a short circuit evaluation, if one branch in OR returns true, cancel all other goroutines
// if one in AND returns false, cancel all other goroutines as well
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
	case returnValue := <-returned:
		return returnValue
	case <-ctx.Done():
		return false
	}

}

func ParsePermissionAndEvaluate(permValue string) (bool, error) {
	// fs := token.NewFileSet()

	tr, err := parser.ParseExpr(permValue)

	// ast.Print(fs, tr)

	if err != nil {
		return false, ErrParsing
	}

	hasPerm := StartEval(&tr)

	return hasPerm, nil
}

// func main() {

// 	// tr, _ := parser.ParseExpr("reader & !writer")
// 	// ast.Print(fs, tr)

// 	// tr, _ := parser.ParseExpr("!parent_folder.reader | reader | writer")
// 	// ast.Print(fs, tr)

// 	// tr, _ := parser.ParseExpr("(reader | writer | tester) & !parent_folder.reader")
// 	// ast.Print(fs, tr)

// 	// tr, _ := parser.ParseExpr("reader & writer")
// 	// ast.Print(fs, tr)

// 	// tr, _ := parser.ParseExpr("reader")
// 	// ast.Print(fs, tr)

// 	// tr, _ := parser.ParseExpr("!reader")
// 	// ast.Print(fs, tr)

// 	permValues := []string{
// 		"reader & !writer",
// 		"!parent_folder.reader | reader | writer",
// 		"(reader | writer | tester) & !parent_folder.reader",
// 		"reader & writer",
// 		"reader",
// 		"!reader",
// 	}

// 	for _, v := range permValues {
// 		hasPerm, err := ParsePermissionAndEvaluate(v)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Printf("has permission '%s'? %t\n", v, hasPerm)
// 	}

// }
