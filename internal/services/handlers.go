package services

import "github.com/DeluxeOwl/kala-go/ent"

type Handler struct {
	// the ent client
	Db *ent.Client
}
