package main

import (
	"testing"

	"github.com/DeluxeOwl/kala-go/ent"
	_ "github.com/mattn/go-sqlite3"
)

func TestClientCreationInMemory(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Fatalf("failed opening connection to sqlite in memory: %v", err)
	}
	defer client.Close()
}
