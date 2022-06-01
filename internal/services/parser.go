package services

import (
	"context"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"

	"github.com/DeluxeOwl/kala-go/ent/relation"
	"github.com/DeluxeOwl/kala-go/internal/models"
)

var ErrParsing = errors.New("error when parsing expression")

func EvalExpr(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler, logChan chan string) bool {
	if ctx.Err() != nil {
		return false
	}

	switch (*expr).(type) {
	case *ast.ParenExpr:
		return EvalParenExpr(ctx, done, expr, pc, h, logChan)
	case *ast.SelectorExpr:
		return EvalSelectorExpr(ctx, done, expr, pc, h, logChan)
	case *ast.Ident:
		return EvalIdent(ctx, done, expr, pc, h, logChan)
	case *ast.UnaryExpr:
		return EvalUnaryExpr(ctx, done, expr, pc, h, logChan)
	case *ast.BinaryExpr:
		return EvalBinaryExpr(ctx, done, expr, pc, h, logChan)
	}

	return false

}

func EvalParenExpr(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler, logChan chan string) bool {

	if ctx.Err() != nil {
		return false
	}

	switch n := (*expr).(type) {
	case *ast.ParenExpr:
		return EvalExpr(ctx, done, &n.X, pc, h, logChan)
	}
	return false
}

func EvalSelectorExpr(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler, logChan chan string) bool {
	if ctx.Err() != nil {
		return false
	}
	switch n := (*expr).(type) {
	case *ast.SelectorExpr:
		if ident, ok := n.X.(*ast.Ident); ok {
			// Get the relation, for this type of query
			rel, err := pc.Perm.
				QueryRelations().
				Where(relation.NameEQ(ident.Name)).
				QueryRelTypeconfigs().
				QueryRelations().
				Where(relation.NameEQ(n.Sel.Name)).
				Only(ctx)

			if err != nil {
				logChan <- fmt.Sprintf("check relation in eval expr: %s\n", err)
				return false
			}
			// Get all referenced subjects
			subjects, err := pc.Perm.
				QueryRelations().
				Where(relation.NameEQ(ident.Name)).
				QueryTuples().
				QuerySubject().
				All(ctx)

			if err != nil {
				logChan <- fmt.Sprintf("get subjects in eval expr: %s\n", err)
				return false
			}

			hasAnyRelPerm := false

			for _, s := range subjects {
				hasAnyRelPerm = hasAnyRelPerm || h.CheckRelation(ctx, &models.RelationCheck{
					Subj: pc.Subj,
					Rel:  rel,
					Res:  s,
				}, 0, logChan)

				if hasAnyRelPerm {
					return hasAnyRelPerm
				}
			}

			return hasAnyRelPerm
		}
	}
	return false
}

func EvalIdent(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler, logChan chan string) bool {
	if ctx.Err() != nil {
		return false
	}

	switch n := (*expr).(type) {
	case *ast.Ident:

		rel, err := pc.Perm.
			QueryRelations().
			Where(relation.NameEQ(n.Name)).
			Only(ctx)

		if err != nil {
			logChan <- fmt.Sprintf("get relations in eval ident: %s\n", err)
			return false
		}

		return h.CheckRelation(ctx, &models.RelationCheck{
			Subj: pc.Subj,
			Rel:  rel,
			Res:  pc.Res,
		}, 0, logChan)
	}

	return false
}

func EvalUnaryExpr(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler, logChan chan string) bool {

	if ctx.Err() != nil {
		return false
	}

	switch n := (*expr).(type) {
	case *ast.UnaryExpr:
		return !EvalExpr(ctx, done, &n.X, pc, h, logChan)
	}

	return false
}

// EvalBinaryExpr runs both sides of the binary tree concurrently
// for a short circuit evaluation, if one branch in OR returns true, cancel all other goroutines
// if one in AND returns false, cancel all other goroutines as well
func EvalBinaryExpr(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler, logChan chan string) bool {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	switch n := (*expr).(type) {
	case *ast.BinaryExpr:
		go func() {
			done <- EvalExpr(ctx, done, &n.X, pc, h, logChan)
		}()
		go func() {
			done <- EvalExpr(ctx, done, &n.Y, pc, h, logChan)
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

func StartEval(expr *ast.Expr, pc *models.PermissionCheck, h *Handler) (bool, []string) {
	ctx := context.Background()
	done := make(chan bool, 2)

	returned := make(chan bool)

	logs := []string{}
	logChan := make(chan string)

	go func() {
		returned <- EvalExpr(ctx, done, expr, pc, h, logChan)
	}()

	for {
		select {
		case returnValue := <-returned:
			return returnValue, logs
		case logLine := <-logChan:
			logs = append(logs, logLine)
		case <-ctx.Done():
			return false, logs
		}
	}

}

func (h *Handler) ParsePermissionAndEvaluate(permValue string, pc *models.PermissionCheck) (bool, []string, error) {
	// fs := token.NewFileSet()

	tr, err := parser.ParseExpr(permValue)

	// ast.Print(fs, tr)

	if err != nil {
		return false, []string{}, ErrParsing
	}

	hasPerm, logs := StartEval(&tr, pc, h)

	return hasPerm, logs, nil
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
// 		logChan <- fmt.Sprintf("has permission `%s`? %t\n", v, hasPerm)
// 	}

// }
