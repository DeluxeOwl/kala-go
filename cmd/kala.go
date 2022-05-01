package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DeluxeOwl/kala-go/ent"
	"github.com/DeluxeOwl/kala-go/ent/typeconfig"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

type TypeConfig struct {
	Name        string
	Relations   map[string]string
	Permissions map[string]string
}

// TODO: order the list of events https://entgo.io/docs/transactions/#transactional-client
// You also have to add refs, so be careful https://entgo.io/docs/schema-edges#o2o-two-types
// check the dependencies
// if they depend, use the id after
func (h *Handler) CreateTypeConfig(ctx context.Context, tc *TypeConfig) error {
	// if tc.Relations == nil {
	// 	fmt.Println("nil map")
	// }
	return nil
}

func (h *Handler) CreateTypeConfigOld(ctx context.Context) (*ent.TypeConfig, error) {

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

	relSlice := make([]*ent.RelationCreate, len(relations))
	cnt := 0
	for i, r := range relations {
		relSlice[cnt] = h.client.Relation.
			Create().
			SetName(i).
			SetValue(r)

		cnt++
	}

	permSlice := make([]*ent.PermissionCreate, len(permissions))
	cnt = 0
	for i, r := range permissions {
		permSlice[cnt] = h.client.Permission.
			Create().
			SetName(i).
			SetValue(r)

		cnt++
	}

	// Save relations and permissions in db
	relBulk, err := h.client.Relation.CreateBulk(relSlice...).Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating relations: %w", err)
	} else {
		fmt.Println("created relations: ", relBulk)
	}

	permBulk, err := h.client.Permission.CreateBulk(permSlice...).Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating relations: %w", err)
	} else {
		fmt.Println("created permissions: ", permBulk)
	}

	tc, err := h.client.TypeConfig.
		Create().
		SetName("document").
		AddRelations(relBulk...).
		AddPermissions(permBulk...).
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

	tc, _ := h.CreateTypeConfigOld(ctx)
	h.QueryTypeConfig(ctx)
	QueryRelations(ctx, tc)
	_ = h.CreateTypeConfig(ctx, &TypeConfig{Name: "docs"})

	// visualize
	// err := entc.Generate("./ent/schema", &gen.Config{}, entc.Extensions(entviz.Extension{}))
	// if err != nil {
	// 	log.Fatalf("running ent codegen: %v", err)
	// }
}
