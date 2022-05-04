package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/DeluxeOwl/kala-go/ent"
	"github.com/DeluxeOwl/kala-go/ent/relation"
	"github.com/DeluxeOwl/kala-go/ent/typeconfig"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
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

var regexAuthzType = regexp.MustCompile(`^[a-zA-Z_]{1,64}(#[a-zA-Z_]{1,64})?( \| [a-zA-Z_]{1,64}(#[a-zA-Z_]{1,64})?)*$`)
var regexAuthzRel = regexp.MustCompile(`^[a-zA-Z_]{1,64}$`)

// TODO: make the errors as variables?
func (h *Handler) CreateTypeConfig(ctx context.Context, tcInput *TypeConfig) (*ent.TypeConfig, error) {

	tc, err := h.client.TypeConfig.Create().
		SetName(tcInput.Name).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating type config: %w", err)
	}

	if tcInput.Relations == nil && tcInput.Permissions != nil {
		return nil, errors.New("failed creating type config: relations cannot be empty while permissions exist")
	}
	// if no relations, we created just the type with the name
	if tcInput.Relations == nil {
		return tc, nil
	}

	// validate if relation types exist
	var refDelim = " | "
	var refRelDelim = "#"

	relSlice := make([]*ent.RelationCreate, len(tcInput.Relations))
	cnt := 0

	log.Printf("for type: %s\n", tcInput.Name)
	for relName, relValue := range tcInput.Relations {

		log.Printf("-> validating '%s: %s'\n", relName, relValue)

		if !regexAuthzType.MatchString(relValue) {
			return nil, fmt.Errorf("malformed relation reference input: %s", relValue)
		}

		if !regexAuthzRel.MatchString(relName) {
			return nil, fmt.Errorf("malformed relation name input: %s", relName)
		}

		referencedTypeIDsSlice := make([]int, 0)

		// type with relation, ex: user | group#member
		for _, referencedType := range strings.Split(relValue, refDelim) {

			if !strings.Contains(referencedType, refRelDelim) {

				id, err := h.client.TypeConfig.
					Query().
					Where(typeconfig.Name(referencedType)).
					OnlyID(ctx)

				if err != nil {
					return nil, fmt.Errorf("referenced type does not exist: %s", referencedType)
				}

				log.Printf("---> found id %d for %s\n", id, referencedType)

				referencedTypeIDsSlice = append(referencedTypeIDsSlice, id)

			} else {
				// checking composed relation
				s := strings.Split(referencedType, refRelDelim)

				refTypeName := s[0]
				refTypeRelation := s[1]

				id, err := h.client.TypeConfig.
					Query().
					Where(typeconfig.Name(refTypeName)).
					OnlyID(ctx)

				if err != nil {
					return nil, fmt.Errorf("referenced type does not exist: %s", refTypeName)
				}

				hasRelation, err := h.client.TypeConfig.
					Query().
					Where(typeconfig.IDEQ(id)).
					QueryRelations().
					Where(relation.NameEQ(refTypeRelation)).
					Exist(ctx)

				if !hasRelation || err != nil {
					return nil, fmt.Errorf("referenced relation does not exist: %s#%s", refTypeName, refTypeRelation)
				}

				log.Printf("---> found id %d for %s\n", id, refTypeName)
				referencedTypeIDsSlice = append(referencedTypeIDsSlice, id)

			}
		}

		relSlice[cnt] = h.client.Relation.
			Create().
			SetName(relName).
			SetValue(relValue).
			AddRelTypeconfigIDs(referencedTypeIDsSlice...)

		cnt++

	}

	// save relations and permissions in db
	relBulk, err := h.client.Relation.CreateBulk(relSlice...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed when creating relations: %w", err)
	}

	tc, err = tc.Update().AddRelations(relBulk...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed when adding relations: %w", err)
	}

	// validate permissions

	var parentRelDelim = "."

	permDelim := regexp.MustCompile(`(( \| )|( & )(!)?)|!`)
	inputRelations := maps.Keys(tcInput.Relations)
	fmt.Println("relations:", inputRelations)
	for permName, permValue := range tcInput.Permissions {

		// TODO: validate permName and permValue

		fmt.Printf("-> validating '%s: %s'\n", permName, permValue)
		for _, referencedRelation := range permDelim.Split(permValue, -1) {
			if referencedRelation == "" {
				continue
			}
			// check direct relations
			if !strings.Contains(referencedRelation, parentRelDelim) {
				if !slices.Contains(inputRelations, referencedRelation) {
					return nil, fmt.Errorf("referenced relation '%s' in permission '%s' not found", referencedRelation, permName)
				}
				fmt.Printf("--> found relation '%s'\n", referencedRelation)
			} else {
				// TODO: check parent relations
				s := strings.Split(referencedRelation, parentRelDelim)
				refParent := s[0]
				refParentRel := s[1]
				fmt.Printf("--> found parent relation '%s' with relation '%s'\n", refParent, refParentRel)
			}

		}
	}

	return tc, nil
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

	// ac, _ := h.CreateTypeConfigOld(ctx)
	// h.QueryTypeConfig(ctx)
	// QueryRelations(ctx, ac)

	tc, err := h.CreateTypeConfig(ctx,
		&TypeConfig{Name: "user"},
	)
	fmt.Println(tc, err)

	tc, err = h.CreateTypeConfig(ctx,
		&TypeConfig{
			Name: "group",
			Relations: map[string]string{
				"member": "user",
			},
		})
	fmt.Println(tc, err)

	tc, err = h.CreateTypeConfig(ctx,
		&TypeConfig{
			Name: "folder",
			Relations: map[string]string{
				"reader": "user | group#member",
			},
		})
	fmt.Println(tc, err)

	tc, err = h.CreateTypeConfig(ctx,
		&TypeConfig{
			Name: "document",
			Relations: map[string]string{
				"parent_folder": "folder",
				"writer":        "user",
				"reader":        "user",
			},
			Permissions: map[string]string{
				"read":           "reader | writer | parent_folder.reader",
				"read_and_write": "reader & writer",
				"read_only":      "reader | !writer",
			},
		})
	fmt.Println(tc, err)

	// TEST: empty permissions
	// tc, err = h.CreateTypeConfig(ctx,
	// 	&TypeConfig{
	// 		Name: "test",
	// 		Permissions: map[string]string{
	// 			"read":           "reader | writer | parent_folder.reader",
	// 			"read_and_write": "reader & writer",
	// 			"read_only":      "reader & !writer",
	// 		},
	// 	})
	// fmt.Println(tc, err)

	// tc, _ = h.CreateTypeConfig(ctx, &TypeConfig{Name: "docs"})

	// VISUALIZE
	// err := entc.Generate("./ent/schema", &gen.Config{}, entc.Extensions(entviz.Extension{}))
	// if err != nil {
	// 	log.Fatalf("running ent codegen: %v", err)
	// }
}
