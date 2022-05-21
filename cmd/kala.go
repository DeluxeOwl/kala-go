package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/DeluxeOwl/kala-go/internal/models"
	"github.com/DeluxeOwl/kala-go/internal/services"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	h := services.NewHandler()
	defer h.Db.Close()
	// Run the auto migration tool.
	if err := h.Db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	h.Http.POST("/v0/typeconfig/batch", func(c echo.Context) error {

		ctx := c.Request().Context()
		h.DeleteEverything(ctx)

		tcReqs := new([]models.TypeConfigReq)

		if err := c.Bind(tcReqs); err != nil {
			return err
		}

		for _, tcReq := range *tcReqs {
			tc, err := h.CreateTypeConfig(ctx, &tcReq)

			if err != nil {
				h.DeleteEverything(ctx)
				return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
					"message": err.Error(),
				})
			}

			fmt.Println(tc)
		}

		return c.JSON(http.StatusCreated, tcReqs)
	})

	h.Http.POST("/v0/subject/batch", func(c echo.Context) error {

		ctx := c.Request().Context()
		h.DeleteSubjects(ctx)

		subjReqs := new([]models.SubjectReq)

		if err := c.Bind(subjReqs); err != nil {
			return err
		}

		for _, subjReq := range *subjReqs {
			subj, err := h.CreateSubject(ctx, &subjReq)

			if err != nil {
				h.DeleteSubjects(ctx)

				return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
					"message": err.Error(),
				})
			}

			fmt.Println(subj)
		}

		return c.JSON(http.StatusCreated, subjReqs)
	})

	h.Http.POST("/v0/tuple/batch", func(c echo.Context) error {

		ctx := c.Request().Context()
		h.DeleteTuples(ctx)

		tupleReqs := new([]models.TupleReqRelation)

		if err := c.Bind(tupleReqs); err != nil {
			return err
		}

		for _, tupleReq := range *tupleReqs {
			tuple, err := h.CreateTuple(ctx, &tupleReq)

			if err != nil {
				h.DeleteTuples(ctx)

				return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
					"message": err.Error(),
				})
			}

			fmt.Println(tuple)
		}

		return c.JSON(http.StatusCreated, tupleReqs)
	})

	h.Http.POST("/v0/permission-check/batch", func(c echo.Context) error {
		permReqs := new([]models.TupleReqPermission)

		if err := c.Bind(permReqs); err != nil {
			return err
		}

		response := []map[string]any{}

		for _, permReq := range *permReqs {

			permCtx := context.Background()
			hasPerm, err := h.CheckPermission(permCtx, &permReq)

			if err != nil {
				response = append(response, map[string]any{
					"permission": false,
					"message":    err.Error(),
				})
			} else {
				var msgFormat string

				if hasPerm {
					msgFormat = "`%s:%s` has permission `%s` on `%s:%s`"
				} else {
					msgFormat = "`%s:%s` doesn't have permission `%s` on `%s:%s`"
				}

				response = append(response, map[string]any{
					"permission": hasPerm,
					"message": fmt.Sprintf(msgFormat,
						permReq.Subject.TypeConfigName,
						permReq.Subject.SubjectName,
						permReq.Permission,
						permReq.Resource.TypeConfigName,
						permReq.Resource.SubjectName),
				})
			}
		}

		return c.JSON(http.StatusCreated, response)
	})

	h.Http.POST("/v0/permission-check", func(c echo.Context) error {
		permReq := new(models.TupleReqPermission)

		if err := c.Bind(permReq); err != nil {
			return err
		}

		var response map[string]any

		permCtx := c.Request().Context()
		hasPerm, err := h.CheckPermission(permCtx, permReq)

		if err != nil {
			response = map[string]any{
				"permission": false,
				"message":    err.Error(),
			}
		} else {
			var msgFormat string

			if hasPerm {
				msgFormat = "`%s:%s` has permission `%s` on `%s:%s`"
			} else {
				msgFormat = "`%s:%s` doesn't have permission `%s` on `%s:%s`"
			}

			response = map[string]any{
				"permission": hasPerm,
				"message": fmt.Sprintf(msgFormat,
					permReq.Subject.TypeConfigName,
					permReq.Subject.SubjectName,
					permReq.Permission,
					permReq.Resource.TypeConfigName,
					permReq.Resource.SubjectName),
			}
		}

		return c.JSON(http.StatusCreated, response)
	})

	h.Http.Logger.Fatal(h.Http.Start(":1323"))

	// h.DeleteEverything(ctx)

}
