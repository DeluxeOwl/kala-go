package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DeluxeOwl/kala-go/ent"
	"github.com/DeluxeOwl/kala-go/ent/typeconfig"
	_ "github.com/mattn/go-sqlite3"
)

func (h *Handler) CreateTypeConfig(ctx context.Context) (*ent.TypeConfig, error) {

	subj, err := h.client.Subject.Create().
		SetName("anna").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating subject: %w", err)
	}
	fmt.Println("subject was created: ", subj)

	relations := map[string]string{
		"parent_folder": "folder",
		"writer":        "user",
		"reader":        "user",
	}

	permissions := map[string]string{
		"read":           "reader | writer | parent_folder.reader",
		"read_and_write": "reader & writer",
		"read_only":      "reader & !writer",
	}

	// TODO: look up bulk create
	// add anna, add the rest of the edges
	relSlice := make([]*ent.Relation, len(relations))
	cnt := 0
	for i, r := range relations {

		rel, err := h.client.Relation.
			Create().
			SetName(i).
			SetValue(r).
			Save(ctx)

		if err != nil {
			return nil, fmt.Errorf("failed creating relation: %w", err)
		}
		relSlice[cnt] = rel
		cnt++
		log.Println("relation was created: ", rel)
	}

	permSlice := make([]*ent.Permission, len(permissions))
	cnt = 0
	for i, r := range permissions {
		rel, err := h.client.Permission.
			Create().
			SetName(i).
			SetValue(r).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating permission: %w", err)
		}
		permSlice[cnt] = rel
		cnt++
		log.Println("permission was created: ", rel)
	}

	tc, err := h.client.TypeConfig.
		Create().
		SetName("document").
		AddRelations(relSlice...).
		AddPermissions(permSlice...).
		AddSubjects(subj).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating type config: %w", err)
	}
	log.Println("type config was created: ", tc)
	return tc, nil
}
func (h *Handler) QueryTypeConfig(ctx context.Context) (*ent.TypeConfig, error) {
	tc, err := h.client.TypeConfig.
		Query().
		Where(typeconfig.Name("document")).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying type config: %w", err)
	}
	log.Println("type config returned: ", tc)
	return tc, nil
}

func QueryRelations(ctx context.Context, tc *ent.TypeConfig) error {
	relations, err := tc.QueryRelations().
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying tc relations: %w", err)
	}
	log.Println("returned relations:\n", relations)

	return nil
}

type Handler struct {
	// the ent client
	client *ent.Client
}

func main() {
	// TODO: handler struct
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()

	h := Handler{
		client: client,
	}

	tc, _ := h.CreateTypeConfig(ctx)
	h.QueryTypeConfig(ctx)
	QueryRelations(ctx, tc)

}
