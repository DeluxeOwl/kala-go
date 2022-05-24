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

func EvalExpr(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler) bool {
	if ctx.Err() != nil {
		return false
	}

	switch (*expr).(type) {
	case *ast.ParenExpr:
		return EvalParenExpr(ctx, done, expr, pc, h)
	case *ast.SelectorExpr:
		return EvalSelectorExpr(ctx, done, expr, pc, h)
	case *ast.Ident:
		return EvalIdent(ctx, done, expr, pc, h)
	case *ast.UnaryExpr:
		return EvalUnaryExpr(ctx, done, expr, pc, h)
	case *ast.BinaryExpr:
		return EvalBinaryExpr(ctx, done, expr, pc, h)
	}

	return false

}

func EvalParenExpr(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler) bool {

	if ctx.Err() != nil {
		return false
	}

	switch n := (*expr).(type) {
	case *ast.ParenExpr:
		return EvalExpr(ctx, done, &n.X, pc, h)
	}
	return false
}

func EvalSelectorExpr(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler) bool {
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
				fmt.Printf("check relation in eval expr: %s\n", err)
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
				fmt.Printf("get subjects in eval expr: %s\n", err)
				return false
			}

			// TODO: goroutines
			hasAnyRelPerm := false

			for _, s := range subjects {
				hasAnyRelPerm = hasAnyRelPerm || h.CheckRelation(ctx, &models.RelationCheck{
					Subj: pc.Subj,
					Rel:  rel,
					Res:  s,
				}, 0)

				if hasAnyRelPerm {
					return hasAnyRelPerm
				}
			}

			return hasAnyRelPerm
		}
	}
	return false
}

func EvalIdent(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler) bool {
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
			fmt.Printf("get relations in eval ident: %s\n", err)
			return false
		}

		return h.CheckRelation(ctx, &models.RelationCheck{
			Subj: pc.Subj,
			Rel:  rel,
			Res:  pc.Res,
		}, 0)
	}

	return false
}

func EvalUnaryExpr(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler) bool {

	if ctx.Err() != nil {
		return false
	}

	switch n := (*expr).(type) {
	case *ast.UnaryExpr:
		return !EvalExpr(ctx, done, &n.X, pc, h)
	}

	return false
}

// EvalBinaryExpr runs both sides of the binary tree concurrently
// for a short circuit evaluation, if one branch in OR returns true, cancel all other goroutines
// if one in AND returns false, cancel all other goroutines as well
func EvalBinaryExpr(ctx context.Context, done chan bool, expr *ast.Expr, pc *models.PermissionCheck, h *Handler) bool {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	switch n := (*expr).(type) {
	case *ast.BinaryExpr:
		go func() {
			done <- EvalExpr(ctx, done, &n.X, pc, h)
		}()
		go func() {
			done <- EvalExpr(ctx, done, &n.Y, pc, h)
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

func StartEval(expr *ast.Expr, pc *models.PermissionCheck, h *Handler) bool {
	ctx := context.Background()
	done := make(chan bool, 2)

	returned := make(chan bool)

	go func() {
		returned <- EvalExpr(ctx, done, expr, pc, h)
	}()

	select {
	case returnValue := <-returned:
		return returnValue
	case <-ctx.Done():
		return false
	}

}

func (h *Handler) ParsePermissionAndEvaluate(permValue string, pc *models.PermissionCheck) (bool, error) {
	// fs := token.NewFileSet()

	tr, err := parser.ParseExpr(permValue)

	// ast.Print(fs, tr)

	if err != nil {
		return false, ErrParsing
	}

	hasPerm := StartEval(&tr, pc, h)

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
// 		fmt.Printf("has permission `%s`? %t\n", v, hasPerm)
// 	}

// }
