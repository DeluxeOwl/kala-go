package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DeluxeOwl/kala-go/ent"
	"github.com/DeluxeOwl/kala-go/ent/schema"
	"github.com/DeluxeOwl/kala-go/ent/typeconfig"
	_ "github.com/mattn/go-sqlite3"
)

func CreateTypeConfig(ctx context.Context, client *ent.Client) (*ent.TypeConfig, error) {
	tc, err := client.TypeConfig.
		Create().
		SetName("document").
		SetRelations(&schema.Relations{
			"parent_folder": "folder",
			"writer":        "user",
			"reader":        "user",
		}).
		SetPermissions(&schema.Permissions{
			"read":           "reader | writer | parent_folder.reader",
			"read_and_write": "reader & writer",
			"read_only":      "reader & !writer",
		}).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating type config: %w", err)
	}
	log.Println("type config was created: ", tc)
	return tc, nil
}
func QueryTypeConfig(ctx context.Context, client *ent.Client) (*ent.TypeConfig, error) {
	tc, err := client.TypeConfig.
		Query().
		// TODO: figure out json
		Where(typeconfig.Name("document")).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying type config: %w", err)
	}
	log.Println("type config returned: ", tc)
	return tc, nil
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

	CreateTypeConfig(ctx, client)
	QueryTypeConfig(ctx, client)
}
