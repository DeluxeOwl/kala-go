package services

import (
	"log"

	"github.com/DeluxeOwl/kala-go/ent"
	"github.com/labstack/echo/v4"
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

	return &Handler{
		Db:   client,
		Http: echo.New(),
	}
}
