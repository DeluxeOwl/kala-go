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

	h.Http.POST("/typeconfig/batch", func(c echo.Context) error {

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

	h.Http.Logger.Fatal(h.Http.Start(":1323"))

	// h.DeleteEverything(ctx)

}
