package services

import (
	"log"

	"github.com/DeluxeOwl/kala-go/ent"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	// the ent client
	Db   *ent.Client
	Http *echo.Echo
}

func NewHandler() *Handler {

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://kala.andreisurugiu.com"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	return &Handler{
		Db:   client,
		Http: e,
	}
}
