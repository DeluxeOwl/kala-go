package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/DeluxeOwl/kala-go/ent"
	"github.com/DeluxeOwl/kala-go/ent/permission"
	"github.com/DeluxeOwl/kala-go/ent/relation"
	"github.com/DeluxeOwl/kala-go/ent/subject"
	"github.com/DeluxeOwl/kala-go/ent/tuple"
	"github.com/DeluxeOwl/kala-go/ent/typeconfig"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func (h *Handler) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := h.client.Tx(ctx)
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

// TODO: separate these methods in packages
func (h *Handler) Do(ctx context.Context) {
	// WithTx helper.
	if err := h.WithTx(ctx, func(tx *ent.Tx) error {
		h := Handler{
			client: tx.Client(),
		}
		tuple, err := h.CreateTuple(ctx, &TupleReqRelation{
			Subject: &SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "anna",
			},
			Relation: "reader",
			Resource: &SubjectReq{
				TypeConfigName: "document",
				SubjectName:    "secret",
			},
		})
		fmt.Println(tuple, err)
		return err
	}); err != nil {
		fmt.Println(err)
	}
}

type TypeConfigReq struct {
	Name        string
	Relations   map[string]string
	Permissions map[string]string
}

// valid regex for relations, permissions and type names
var regexPropertyName = regexp.MustCompile(`^[a-zA-Z_]{1,64}$`)
var regexTypeName = regexp.MustCompile(`^[a-zA-Z_]{1,64}$`)

// valid regexes for relation, permission values
var regexRelValue = regexp.MustCompile(`^[a-zA-Z_]{1,64}(#[a-zA-Z_]{1,64})?( \| [a-zA-Z_]{1,64}(#[a-zA-Z_]{1,64})?)*$`)
var regexPermValue = regexp.MustCompile(`^((!)?[a-zA-Z_]{1,64}(\.[a-zA-Z_]{1,64})?)((( \| )|( & ))((!)?[a-zA-Z_]{1,64}(\.[a-zA-Z_]{1,64})?))*$`)

// delimiter for the value of a relation
var refValueDelim = " | "

// delimiter for a subrelation
var refSubrelationDelim = "#"

// delimiter for parent relation in permission
var parentRelDelim = "."

// delimiter for permissions
var permDelim = regexp.MustCompile(`(( \| )|( & )(!)?)|!`)

// TODO: make the errors as variables?
func (h *Handler) CreateTypeConfig(ctx context.Context, tcInput *TypeConfigReq) (*ent.TypeConfig, error) {

	if !regexTypeName.MatchString(tcInput.Name) {
		return nil, fmt.Errorf("malformed type name input: '%s'", tcInput.Name)
	}

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

	relSlice := make([]*ent.RelationCreate, len(tcInput.Relations))
	cnt := 0

	log.Printf("for type: %s\n", tcInput.Name)
	for relName, relValue := range tcInput.Relations {

		log.Printf("-> validating '%s: %s'\n", relName, relValue)

		if !regexPropertyName.MatchString(relName) {
			return nil, fmt.Errorf("malformed relation name input: '%s'", relName)
		}

		if !regexRelValue.MatchString(relValue) {
			return nil, fmt.Errorf("malformed relation reference input: '%s'", relValue)
		}

		referencedTypeIDsSlice := make([]int, 0)

		// type with relation, ex: user | group#member
		for _, referencedType := range strings.Split(relValue, refValueDelim) {

			if !strings.Contains(referencedType, refSubrelationDelim) {

				id, err := h.client.TypeConfig.
					Query().
					Where(typeconfig.Name(referencedType)).
					OnlyID(ctx)

				if err != nil {
					return nil, fmt.Errorf("referenced type does not exist: '%s'", referencedType)
				}

				log.Printf("---> found id %d for %s\n", id, referencedType)

				referencedTypeIDsSlice = append(referencedTypeIDsSlice, id)

			} else {
				// checking composed relation
				s := strings.Split(referencedType, refSubrelationDelim)

				refTypeName := s[0]
				refTypeRelation := s[1]

				id, err := h.client.TypeConfig.
					Query().
					Where(typeconfig.Name(refTypeName)).
					OnlyID(ctx)

				if err != nil {
					return nil, fmt.Errorf("referenced type does not exist: '%s'", refTypeName)
				}

				hasRelation, err := h.client.TypeConfig.
					Query().
					Where(typeconfig.IDEQ(id)).
					QueryRelations().
					Where(relation.NameEQ(refTypeRelation)).
					Exist(ctx)

				if !hasRelation || err != nil {
					return nil, fmt.Errorf("referenced relation does not exist: '%s#%s'", refTypeName, refTypeRelation)
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

	permSlice := make([]*ent.PermissionCreate, len(tcInput.Permissions))
	cnt = 0

	inputRelations := maps.Keys(tcInput.Relations)
	fmt.Println("relations:", inputRelations)

	getIdForRelationName := func(name string) (int, error) {
		for _, rel := range relBulk {
			if rel.Name == name {
				return rel.ID, nil
			}
		}

		return -1, fmt.Errorf("adding relation: couldn't get id for relation '%s'", name)
	}

	for permName, permValue := range tcInput.Permissions {

		if !regexPropertyName.MatchString(permName) {
			return nil, fmt.Errorf("malformed permission name input: '%s'", permName)
		}

		if !regexPermValue.MatchString(permValue) {
			return nil, fmt.Errorf("malformed permission reference input: '%s'", permValue)
		}

		referencedTypeIDsSlice := make([]int, 0)

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

				relID, err := getIdForRelationName(referencedRelation)
				if err != nil {
					return nil, err
				}

				referencedTypeIDsSlice = append(referencedTypeIDsSlice, relID)

				fmt.Printf("--> found relation '%s' with id %d \n", referencedRelation, relID)
			} else {
				// check parent relations
				s := strings.Split(referencedRelation, parentRelDelim)

				refParent := s[0]
				refParentRel := s[1]

				if !slices.Contains(inputRelations, refParent) {
					return nil, fmt.Errorf("referenced relation '%s' in permission '%s' not found", referencedRelation, permName)
				}
				// error if composed relation: user | group#member
				referencedType := tcInput.Relations[refParent]

				if strings.Contains(referencedType, refValueDelim) {
					return nil, fmt.Errorf("referenced relation '%s' can't contain a composed value: '%s'", refParent, referencedType)
				}

				hasRelation, err := h.client.TypeConfig.
					Query().
					Where(typeconfig.NameEQ(referencedType)).
					QueryRelations().
					Where(relation.NameEQ(refParentRel)).
					Exist(ctx)

				if !hasRelation || err != nil {
					return nil, fmt.Errorf("referenced type '%s' in relation '%s' doesn't have a '%s' relation",
						referencedType,
						refParent,
						refParentRel)
				}

				relID, err := getIdForRelationName(refParent)
				if err != nil {
					return nil, err
				}

				referencedTypeIDsSlice = append(referencedTypeIDsSlice, relID)

				fmt.Printf("--> found parent relation '%s' with relation '%s' with id %d\n", refParent, refParentRel, relID)
			}

		}

		permSlice[cnt] = h.client.Permission.
			Create().
			SetName(permName).
			SetValue(permValue).
			AddRelationIDs(referencedTypeIDsSlice...)

		cnt++

	}

	// save relations and permissions in db
	permBulk, err := h.client.Permission.CreateBulk(permSlice...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed when creating permissions: %w", err)
	}

	tc, err = tc.Update().AddPermissions(permBulk...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed when adding permissions: %w", err)
	}

	return tc, nil
}

var regexSubjName = regexp.MustCompile(`^[a-zA-Z0-9\._\/-]{1,64}$`)

type SubjectReq struct {
	TypeConfigName string
	SubjectName    string
}

func (h *Handler) CreateSubject(ctx context.Context, s *SubjectReq) (*ent.Subject, error) {

	if !regexSubjName.MatchString(s.SubjectName) {
		return nil, fmt.Errorf("malformed subj name input: '%s'", s.SubjectName)
	}

	exists, err := h.client.TypeConfig.
		Query().
		Where(typeconfig.NameEQ(s.TypeConfigName)).
		Exist(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying type config: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("failed when creating subject, type '%s' does not exist", s.TypeConfigName)
	}

	subj, err := h.client.Subject.Create().SetName(s.SubjectName).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating subject: %w", err)
	}

	_, err = h.client.TypeConfig.Update().
		Where(typeconfig.NameEQ(s.TypeConfigName)).
		AddSubjects(subj).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when updating typeconfig: %w", err)
	}

	return subj, nil
}

type TupleReqRelation struct {
	Subject  *SubjectReq
	Relation string
	Resource *SubjectReq
}

// TODO handle group:* viewer on document
// create a default subject called '*' when creating the type?
func (h *Handler) CreateTuple(ctx context.Context, tr *TupleReqRelation) (*ent.Tuple, error) {

	var subjectName string
	var subjectRelation string = ""

	if !strings.Contains(tr.Subject.SubjectName, "#") {
		subjectName = tr.Subject.SubjectName
	} else {
		subjectName = strings.Split(tr.Subject.SubjectName, "#")[0]
		subjectRelation = strings.Split(tr.Subject.SubjectName, "#")[1]
	}

	// handling cases like group:dev#member and group:dev#*
	if subjectRelation != "" {
		hasRelation, err := h.client.TypeConfig.
			Query().
			Where(typeconfig.NameEQ(tr.Subject.TypeConfigName)).
			QueryRelations().
			Where(relation.NameEQ(subjectRelation)).
			Exist(ctx)
		if !hasRelation || err != nil {
			return nil, fmt.Errorf("error when checking subject relation: %w", err)
		}
	}

	subj, err := h.client.Subject.
		Query().
		Where(
			subject.And(
				subject.NameEQ(subjectName),
				subject.HasTypeWith(
					typeconfig.NameEQ(tr.Subject.TypeConfigName)))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when getting subject: %w", err)
	}

	res, err := h.client.Subject.
		Query().
		Where(
			subject.And(
				subject.NameEQ(tr.Resource.SubjectName),
				subject.HasTypeWith(
					typeconfig.NameEQ(tr.Resource.TypeConfigName)))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when getting resource: %w", err)
	}

	rel, err := h.client.Relation.
		Query().
		Where(
			relation.And(
				relation.HasTypeconfigWith(
					typeconfig.NameEQ(tr.Resource.TypeConfigName)),
				relation.NameEQ(tr.Relation))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when getting relation for resource: %w", err)
	}

	tupleCreate := h.client.Tuple.
		Create().
		SetSubject(subj).
		SetRelation(rel).
		SetResource(res)

	if subjectRelation != "" {
		tupleCreate.SetSubjectRel(subjectRelation)
	}

	tuple, err := tupleCreate.Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when creating tuple: %w", err)
	}

	return tuple, nil
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

	tc, err := h.CreateTypeConfig(ctx,
		&TypeConfigReq{Name: "user"},
	)
	fmt.Println(tc, err)

	tc, err = h.CreateTypeConfig(ctx,
		&TypeConfigReq{
			Name: "group",
			Relations: map[string]string{
				"member": "user",
			},
		})
	fmt.Println(tc, err)

	tc, err = h.CreateTypeConfig(ctx,
		&TypeConfigReq{
			Name: "folder",
			Relations: map[string]string{
				"reader": "user | group#member",
			},
		})
	fmt.Println(tc, err)

	tc, err = h.CreateTypeConfig(ctx,
		&TypeConfigReq{
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

	subjects := []SubjectReq{
		{
			TypeConfigName: "document",
			SubjectName:    "report.csv",
		},
		{
			TypeConfigName: "user",
			SubjectName:    "anna",
		},
		{
			TypeConfigName: "user",
			SubjectName:    "john",
		},
		{
			TypeConfigName: "user",
			SubjectName:    "steve",
		},
		{
			TypeConfigName: "folder",
			SubjectName:    "secret_folder",
		},
		{
			TypeConfigName: "group",
			SubjectName:    "dev",
		},
		{
			TypeConfigName: "group",
			SubjectName:    "test_group",
		},
	}

	for _, v := range subjects {
		subj, err := h.CreateSubject(ctx, &v)
		fmt.Println(subj, err)
	}

	tuples := []TupleReqRelation{
		{
			Subject: &SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "anna",
			},
			Relation: "reader",
			Resource: &SubjectReq{
				TypeConfigName: "document",
				SubjectName:    "report.csv",
			},
		},
		{
			Subject: &SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "anna",
			},
			Relation: "writer",
			Resource: &SubjectReq{
				TypeConfigName: "document",
				SubjectName:    "report.csv",
			},
		},
		{
			Subject: &SubjectReq{
				TypeConfigName: "folder",
				SubjectName:    "secret_folder",
			},
			Relation: "parent_folder",
			Resource: &SubjectReq{
				TypeConfigName: "document",
				SubjectName:    "report.csv",
			},
		},
		{
			Subject: &SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "john",
			},
			Relation: "reader",
			Resource: &SubjectReq{
				TypeConfigName: "folder",
				SubjectName:    "secret_folder",
			},
		},
		{
			Subject: &SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "john",
			},
			Relation: "member",
			Resource: &SubjectReq{
				TypeConfigName: "group",
				SubjectName:    "dev",
			},
		},
		{
			Subject: &SubjectReq{
				TypeConfigName: "group",
				SubjectName:    "dev#member",
			},
			Relation: "reader",
			Resource: &SubjectReq{
				TypeConfigName: "folder",
				SubjectName:    "secret_folder",
			},
		},
		{
			Subject: &SubjectReq{
				TypeConfigName: "group",
				SubjectName:    "test_group#member",
			},
			Relation: "reader",
			Resource: &SubjectReq{
				TypeConfigName: "folder",
				SubjectName:    "secret_folder",
			},
		},
		{
			Subject: &SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "steve",
			},
			Relation: "member",
			Resource: &SubjectReq{
				TypeConfigName: "group",
				SubjectName:    "dev",
			},
		},
	}

	for _, v := range tuples {
		tuple, err := h.CreateTuple(ctx, &v)
		fmt.Println(tuple, err)
	}

	// h.Do(ctx)

	hasRead, err := h.CheckPermission(ctx, &TupleReqPermission{
		Subject: &SubjectReq{
			TypeConfigName: "user",
			SubjectName:    "steve",
		},
		Permission: "read",
		Resource: &SubjectReq{
			TypeConfigName: "document",
			SubjectName:    "report.csv",
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("⟶\tCan steve read?", hasRead)

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

	// reader | writer | parent_folder.reader
	// reader & writer
	// reader | writer
	// reader | !writer
	// !writer
	// reader
	// reader & writer | !tester
	// !tester & reader | writer

	// reader |
	// !writer |
	// reader &
	// writer & reader | !
}

type RelationCheck struct {
	Subj *ent.Subject
	Rel  *ent.Relation
	Res  *ent.Subject
}

// TODO: add some depth, maybe add path taken (ident in a json and print graph?), goroutines
func (h *Handler) CheckRelation(ctx context.Context, rc *RelationCheck, depth int) bool {

	fmt.Println(strings.Repeat("\t", depth+1)+"• Checking relation:", rc)

	subjTypeString := rc.Subj.QueryType().Select(typeconfig.FieldName).StringX(ctx)

	// check if direct
	if !strings.Contains(rc.Rel.Value, refValueDelim) {

		if rc.Rel.Value != subjTypeString {
			fmt.Println("\t\t Type doesnt match with subject")
			return false
		}

		tupleExists, err := h.client.Tuple.
			Query().
			Where(
				tuple.And(
					tuple.SubjectID(rc.Subj.ID),
					tuple.RelationID(rc.Rel.ID),
					tuple.ResourceID(rc.Res.ID),
				),
			).
			Exist(ctx)

		fmt.Printf("\t\tDirect relationship? %t\n", tupleExists)

		if err != nil {
			fmt.Println(err)
		}

		return tupleExists

	} else {
		for _, referencedType := range strings.Split(rc.Rel.Value, refValueDelim) {

			// if direct type
			if referencedType == subjTypeString {
				tupleExists, err := h.client.Tuple.
					Query().
					Where(
						tuple.And(
							tuple.SubjectID(rc.Subj.ID),
							tuple.RelationID(rc.Rel.ID),
							tuple.ResourceID(rc.Res.ID),
						),
					).
					Exist(ctx)

				fmt.Printf("\t\tDirect relationship? %t\n", tupleExists)

				if err != nil {
					fmt.Println(err)
				}

				if tupleExists {
					return true
				}

			} else {
				// checking composed relation
				// TODO: repeated, own function
				s := strings.Split(referencedType, refSubrelationDelim)

				refTypeName := s[0]
				refTypeRelation := s[1]

				// Get all refTypeName and get the relation refTypeName (e.g. group#member) with relation rc.Rel.Name
				fmt.Printf("%sGetting all `%s:<name>#%s` with relation `%s` on `%s`\n",
					strings.Repeat("\t", depth+2),
					refTypeName,
					refTypeRelation,
					rc.Rel.Name,
					rc.Res.Name)
				fmt.Printf("%sand checking if subject `%s` has relation `%s` on each group\n",
					strings.Repeat("\t", depth+2),
					rc.Subj.Name,
					refTypeRelation)

				refTypeRelationObject, err := h.client.TypeConfig.
					Query().
					Where(typeconfig.NameEQ(refTypeName)).
					QueryRelations().
					Where(relation.NameEQ(refTypeRelation)).
					Only(ctx)

				if err != nil {
					fmt.Println(err)
					return false
				}

				subjects, err := h.client.Tuple.
					Query().
					Where(tuple.And(
						tuple.SubjectRelEQ(refTypeRelation),
						tuple.RelationID(rc.Rel.ID),
						tuple.ResourceID(rc.Res.ID),
					)).
					QuerySubject().
					All(ctx)

				if err != nil {
					fmt.Println(err)
					return false
				}

				relExists := false
				for _, s := range subjects {

					relExists = relExists || h.CheckRelation(ctx, &RelationCheck{
						Subj: rc.Subj,
						Rel:  refTypeRelationObject,
						Res:  s,
					}, depth+1)

					if relExists {
						return relExists
					}
				}
			}
		}
	}

	return false
}

func (h *Handler) Subject(ctx context.Context, tfName string, name string) (*ent.Subject, error) {
	subj, err := h.client.Subject.
		Query().
		Where(
			subject.And(
				subject.NameEQ(name),
				subject.HasTypeWith(typeconfig.NameEQ(tfName)))).
		Only(ctx)
	return subj, err
}

type TupleReqPermission struct {
	Subject    *SubjectReq
	Permission string
	Resource   *SubjectReq
}

func (h *Handler) CheckPermission(ctx context.Context, tr *TupleReqPermission) (bool, error) {
	fmt.Printf("====> Does `%s:%s` have `%s` permission on `%s:%s`?\n",
		tr.Subject.TypeConfigName,
		tr.Subject.SubjectName,
		tr.Permission,
		tr.Resource.TypeConfigName,
		tr.Resource.SubjectName)

	subj, err := h.Subject(ctx, tr.Subject.TypeConfigName, tr.Subject.SubjectName)

	if err != nil {
		return false, fmt.Errorf("error when querying subject: %w", err)
	}

	fmt.Println("Checking if subject exists:", subj)

	res, err := h.Subject(ctx, tr.Resource.TypeConfigName, tr.Resource.SubjectName)

	if err != nil {
		return false, fmt.Errorf("error when querying resource: %w", err)
	}

	fmt.Println("Checking if resource exists:", res)

	permQuery := res.
		QueryType().
		QueryPermissions().
		Where(permission.NameEQ(tr.Permission))

	perm, err := permQuery.
		Only(ctx)

	if err != nil {
		return false, fmt.Errorf("error when querying permission: %w", err)
	}

	fmt.Println("Getting relations for perm:", perm)

	// Getting the result from the check relations
	hasPerm := false

	// TODO: refactor this in its own function, same one as in CreateTypeConfig
	// or you should probably parse it
	// probably shouldnt return on all errors
	for _, referencedRelation := range permDelim.Split(perm.Value, -1) {
		if referencedRelation == "" {
			continue
		}

		// check direct relations and indirect relations
		if !strings.Contains(referencedRelation, parentRelDelim) {
			rel, err := permQuery.
				QueryRelations().
				Where(relation.NameEQ(referencedRelation)).
				Only(ctx)

			if err != nil {
				return false, fmt.Errorf("error when querying relation: %w", err)
			}

			hasPerm = hasPerm || h.CheckRelation(ctx, &RelationCheck{
				Subj: subj,
				Rel:  rel,
				Res:  res,
			}, 0)

			if hasPerm {
				return hasPerm, nil
			}

		} else {
			s := strings.Split(referencedRelation, parentRelDelim)

			refParent := s[0]
			refParentRel := s[1]

			// Get the relation, for this type of query
			// it can only be one anyway
			r, err := perm.QueryRelations().
				Where(relation.NameEQ(refParent)).
				QueryRelTypeconfigs().
				QueryRelations().
				Where(relation.NameEQ(refParentRel)).
				Only(ctx)

			if err != nil {
				return false, fmt.Errorf("error when querying relation: %w", err)
			}

			// Get all referenced subjects
			subjects, err := perm.QueryRelations().
				Where(relation.NameEQ(refParent)).
				QueryTuples().
				QuerySubject().
				All(ctx)

			// if it doesn't have subjects, just move on
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, s := range subjects {
				hasPerm = hasPerm || h.CheckRelation(ctx, &RelationCheck{
					Subj: subj,
					Rel:  r,
					Res:  s,
				}, 0)

				if hasPerm {
					return hasPerm, nil
				}
			}
		}

	}

	return hasPerm, nil
}
